package zk_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	. "github.com/facebookincubator/zk"
	"github.com/facebookincubator/zk/testutils"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

const defaultMaxRetries = 5

func TestClientRetryLogic(t *testing.T) {
	failCalls := defaultMaxRetries

	// Create a new handler which will make the test server return an error for a set number of tries.
	// We expect the client to recover from these errors and retry the RPC calls until success on the last try.
	server, err := testutils.NewServer(func(req jute.RecordReader) jute.RecordWriter {
		if failCalls > 0 {
			failCalls--
			return nil // nil response causes error
		}

		return testutils.DefaultHandler(req)
	})
	if err != nil {
		t.Fatalf("error creating test server: %v", err)
	}
	defer server.Close()

	client := &Client{
		MaxRetries: defaultMaxRetries,
		Network:    server.Addr().Network(),
		Ensemble:   server.Addr().String(),
	}

	children, err := client.GetChildren(context.Background(), "/")
	if err != nil {
		t.Fatalf("unexpected error calling GetData: %v", err)
	}

	if expected := []string{"test"}; !reflect.DeepEqual(expected, children) {
		t.Fatalf("getChildren error: expected %v, got %v", expected, children)
	}
}

func TestClientRetryLogicFails(t *testing.T) {
	server, err := testutils.NewDefaultServer()
	if err != nil {
		t.Fatalf("error creating test server: %v", err)
	}

	client := &Client{
		MaxRetries: defaultMaxRetries,
		Network:    server.Addr().Network(),
		Ensemble:   server.Addr().String(),
	}

	// close server before client makes RPC call
	if err = server.Close(); err != nil {
		t.Fatalf("unexpected error closing server: %v", err)
	}

	_, err = client.GetChildren(context.Background(), "/")
	if err == nil || !errors.Is(err, ErrMaxRetries) {
		t.Fatalf("expected error: \"%v\", got error: \"%v\"", ErrMaxRetries, err)
	}
}

func TestClientContextCanceled(t *testing.T) {
	calls := 0
	server, err := testutils.NewServer(func(req jute.RecordReader) jute.RecordWriter {
		calls++

		return testutils.DefaultHandler(req)
	})
	if err != nil {
		t.Fatalf("error creating test server: %v", err)
	}

	client := &Client{
		MaxRetries: defaultMaxRetries,
		Ensemble:   server.Addr().String(),
		Network:    server.Addr().Network(),
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// expect the client not to retry when ctx is canceled
	if _, err = client.GetData(ctx, "/"); !errors.Is(err, ctx.Err()) {
		t.Fatalf("unexpected error calling GetData: %v", err)
	}
	// fail if client attempted retries on canceled ctx
	if calls > 1 {
		t.Fatalf("ctx.Err() is non-retryable, expected only 1 call, got %d", calls)
	}
}
