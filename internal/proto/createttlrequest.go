// Autogenerated jute compiler
// @generated from 'zookeeper.jute'

package proto // github.com/facebookincubator/zk/internal/proto

import (
	"fmt"

	"github.com/facebookincubator/zk/internal/data"
	"github.com/go-zookeeper/jute/lib/go/jute"
)

type CreateTTLRequest struct {
	Path  string     // path
	Data  []byte     // data
	Acl   []data.ACL // acl
	Flags int32      // flags
	Ttl   int64      // ttl
}

func (r *CreateTTLRequest) GetPath() string {
	if r != nil {
		return r.Path
	}
	return ""
}

func (r *CreateTTLRequest) GetData() []byte {
	if r != nil && r.Data != nil {
		return r.Data
	}
	return nil
}

func (r *CreateTTLRequest) GetAcl() []data.ACL {
	if r != nil && r.Acl != nil {
		return r.Acl
	}
	return nil
}

func (r *CreateTTLRequest) GetFlags() int32 {
	if r != nil {
		return r.Flags
	}
	return 0
}

func (r *CreateTTLRequest) GetTtl() int64 {
	if r != nil {
		return r.Ttl
	}
	return 0
}

func (r *CreateTTLRequest) Read(dec jute.Decoder) (err error) {
	var size int
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.Path, err = dec.ReadString()
	if err != nil {
		return err
	}
	r.Data, err = dec.ReadBuffer()
	if err != nil {
		return err
	}
	size, err = dec.ReadVectorStart()
	if err != nil {
		return err
	}
	if size < 0 {
		r.Acl = nil
	} else {
		r.Acl = make([]data.ACL, size)
		for i := 0; i < size; i++ {
			if err = dec.ReadRecord(&r.Acl[i]); err != nil {
				return err
			}
		}
	}
	if err = dec.ReadVectorEnd(); err != nil {
		return err
	}
	r.Flags, err = dec.ReadInt()
	if err != nil {
		return err
	}
	r.Ttl, err = dec.ReadLong()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *CreateTTLRequest) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteString(r.Path); err != nil {
		return err
	}
	if err := enc.WriteBuffer(r.Data); err != nil {
		return err
	}
	if err := enc.WriteVectorStart(len(r.Acl), r.Acl == nil); err != nil {
		return err
	}
	for _, v := range r.Acl {
		if err := enc.WriteRecord(&v); err != nil {
			return err
		}
	}
	if err := enc.WriteVectorEnd(); err != nil {
		return err
	}
	if err := enc.WriteInt(r.Flags); err != nil {
		return err
	}
	if err := enc.WriteLong(r.Ttl); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *CreateTTLRequest) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CreateTTLRequest(%+v)", *r)
}
