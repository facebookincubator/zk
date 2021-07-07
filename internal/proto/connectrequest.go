// Autogenerated jute compiler
// @generated from 'zookeeper.jute'

package proto // github.com/facebookincubator/zk/internal/proto

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type ConnectRequest struct {
	ProtocolVersion int32  // protocolVersion
	LastZxidSeen    int64  // lastZxidSeen
	TimeOut         int32  // timeOut
	SessionId       int64  // sessionId
	Passwd          []byte // passwd
}

func (r *ConnectRequest) GetProtocolVersion() int32 {
	if r != nil {
		return r.ProtocolVersion
	}
	return 0
}

func (r *ConnectRequest) GetLastZxidSeen() int64 {
	if r != nil {
		return r.LastZxidSeen
	}
	return 0
}

func (r *ConnectRequest) GetTimeOut() int32 {
	if r != nil {
		return r.TimeOut
	}
	return 0
}

func (r *ConnectRequest) GetSessionId() int64 {
	if r != nil {
		return r.SessionId
	}
	return 0
}

func (r *ConnectRequest) GetPasswd() []byte {
	if r != nil && r.Passwd != nil {
		return r.Passwd
	}
	return nil
}

func (r *ConnectRequest) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.ProtocolVersion, err = dec.ReadInt()
	if err != nil {
		return err
	}
	r.LastZxidSeen, err = dec.ReadLong()
	if err != nil {
		return err
	}
	r.TimeOut, err = dec.ReadInt()
	if err != nil {
		return err
	}
	r.SessionId, err = dec.ReadLong()
	if err != nil {
		return err
	}
	r.Passwd, err = dec.ReadBuffer()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *ConnectRequest) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteInt(r.ProtocolVersion); err != nil {
		return err
	}
	if err := enc.WriteLong(r.LastZxidSeen); err != nil {
		return err
	}
	if err := enc.WriteInt(r.TimeOut); err != nil {
		return err
	}
	if err := enc.WriteLong(r.SessionId); err != nil {
		return err
	}
	if err := enc.WriteBuffer(r.Passwd); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *ConnectRequest) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ConnectRequest(%+v)", *r)
}