// Autogenerated jute compiler
// @generated from 'zookeeper.jute'

package proto // github.com/facebookincubator/zk/internal/proto

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type GetChildren2Request struct {
	Path  string // path
	Watch bool   // watch
}

func (r *GetChildren2Request) GetPath() string {
	if r != nil {
		return r.Path
	}
	return ""
}

func (r *GetChildren2Request) GetWatch() bool {
	if r != nil {
		return r.Watch
	}
	return false
}

func (r *GetChildren2Request) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.Path, err = dec.ReadString()
	if err != nil {
		return err
	}
	r.Watch, err = dec.ReadBoolean()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *GetChildren2Request) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteString(r.Path); err != nil {
		return err
	}
	if err := enc.WriteBoolean(r.Watch); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *GetChildren2Request) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetChildren2Request(%+v)", *r)
}
