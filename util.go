package zk

import (
	"bytes"
	"errors"
	"fmt"
	"net"

	"github.com/facebookincubator/zk/internal/proto"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

// WriteRecords takes in one or more RecordWriter instances, serializes them to a byte array
// and writes them to the provided net.Conn.
func WriteRecords(conn net.Conn, generated ...jute.RecordWriter) error {
	sendBuf := &bytes.Buffer{}
	enc := jute.NewBinaryEncoder(sendBuf)

	for _, generatedStruct := range generated {
		if generatedStruct == nil {
			return errors.New("cannot encode nil struct")
		}
		if err := generatedStruct.Write(enc); err != nil {
			return fmt.Errorf("could not encode struct: %w", err)
		}
	}
	// copy encoded request bytes
	requestBytes := append([]byte(nil), sendBuf.Bytes()...)

	// use encoder to prepend request length to the request bytes
	sendBuf.Reset()

	if err := enc.WriteBuffer(requestBytes); err != nil {
		return fmt.Errorf("could not write buffer: %w", err)
	}
	if err := enc.WriteEnd(); err != nil {
		return fmt.Errorf("could not write buffer: %w", err)
	}

	if _, err := conn.Write(sendBuf.Bytes()); err != nil {
		return fmt.Errorf("error writing to net.conn: %w", err)
	}

	return nil
}

// ReadRecord reads the request header and body depending on the opcode.
// It returns the serialized request header and body, or an error if it occurs.
func ReadRecord(dec *jute.BinaryDecoder) (*proto.RequestHeader, jute.RecordReader, error) {
	if _, err := dec.ReadInt(); err != nil {
		return nil, nil, fmt.Errorf("error reading request length: %w", err)
	}

	header := &proto.RequestHeader{}

	if err := dec.ReadRecord(header); err != nil {
		return nil, nil, fmt.Errorf("error reading RequestHeader: %w", err)
	}

	var req jute.RecordReader
	switch header.Type {
	case opGetData:
		req = &proto.GetDataRequest{}
	case opGetChildren:
		req = &proto.GetChildrenRequest{}
	default:
		return nil, nil, fmt.Errorf("unrecognized header type: %d", header.Type)
	}

	if err := dec.ReadRecord(req); err != nil {
		return nil, nil, fmt.Errorf("error reading request: %w", err)
	}

	return header, req, nil
}
