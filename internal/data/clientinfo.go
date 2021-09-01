// Autogenerated jute compiler
// @generated from 'zookeeper.jute'

package data // github.com/facebookincubator/zk/internal/data

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type ClientInfo struct {
	AuthScheme string // authScheme
	User       string // user
}

func (r *ClientInfo) GetAuthScheme() string {
	if r != nil {
		return r.AuthScheme
	}
	return ""
}

func (r *ClientInfo) GetUser() string {
	if r != nil {
		return r.User
	}
	return ""
}

func (r *ClientInfo) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.AuthScheme, err = dec.ReadString()
	if err != nil {
		return err
	}
	r.User, err = dec.ReadString()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *ClientInfo) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteString(r.AuthScheme); err != nil {
		return err
	}
	if err := enc.WriteString(r.User); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *ClientInfo) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ClientInfo(%+v)", *r)
}
