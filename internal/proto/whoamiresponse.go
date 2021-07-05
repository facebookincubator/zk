// Autogenerated jute compiler
// @generated from 'zookeeper.jute'

package proto // github.com/facebookincubator/zk/internal/proto

import (
	"fmt"

	"github.com/facebookincubator/zk/internal/data"
	"github.com/go-zookeeper/jute/lib/go/jute"
)

type WhoAmIResponse struct {
	ClientInfo []*data.ClientInfo // clientInfo
}

func (r *WhoAmIResponse) GetClientInfo() []*data.ClientInfo {
	if r != nil && r.ClientInfo != nil {
		return r.ClientInfo
	}
	return nil
}

func (r *WhoAmIResponse) Read(dec jute.Decoder) (err error) {
	var size int
	if err = dec.ReadStart(); err != nil {
		return err
	}
	size, err = dec.ReadVectorStart()
	if err != nil {
		return err
	}
	if size < 0 {
		r.ClientInfo = nil
	} else {
		r.ClientInfo = make([]*data.ClientInfo, size)
		for i := 0; i < size; i++ {
			if err = dec.ReadRecord(r.ClientInfo[i]); err != nil {
				return err
			}
		}
	}
	if err = dec.ReadVectorEnd(); err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *WhoAmIResponse) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteVectorStart(len(r.ClientInfo), r.ClientInfo == nil); err != nil {
		return err
	}
	for _, v := range r.ClientInfo {
		if err := enc.WriteRecord(v); err != nil {
			return err
		}
	}
	if err := enc.WriteVectorEnd(); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *WhoAmIResponse) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WhoAmIResponse(%+v)", *r)
}
