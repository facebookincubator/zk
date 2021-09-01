// Autogenerated jute compiler
// @generated from 'zookeeper.jute'

package proto // github.com/facebookincubator/zk/internal/proto

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type GetAllChildrenNumberResponse struct {
	TotalNumber int32 // totalNumber
}

func (r *GetAllChildrenNumberResponse) GetTotalNumber() int32 {
	if r != nil {
		return r.TotalNumber
	}
	return 0
}

func (r *GetAllChildrenNumberResponse) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.TotalNumber, err = dec.ReadInt()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *GetAllChildrenNumberResponse) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteInt(r.TotalNumber); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *GetAllChildrenNumberResponse) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetAllChildrenNumberResponse(%+v)", *r)
}
