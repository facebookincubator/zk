package zk

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"time"

	"github.com/facebookincubator/zk/flw"
	"github.com/facebookincubator/zk/internal/proto"
	"github.com/go-zookeeper/jute/lib/go/jute"
)

const defaultProtocolVersion = 0

var ErrSessionExpired = errors.New("zk: session has been expired by the server")
var emptyPassword = make([]byte, 16)

type Connection struct {
	provider       HostProvider
	conn           net.Conn
	lastZxid       int64
	sessionTimeout time.Duration
	sessionID      int64
	passwd         []byte
	server         string
}

func Connect(servers []string) (*Connection, error) {
	conn := &Connection{
		provider:       &DNSHostProvider{},
		sessionTimeout: time.Second,
		passwd:         emptyPassword,
	}

	err := conn.provider.Init(flw.FormatServers(servers))
	if err != nil {
		return nil, err
	}

	err = conn.dial()
	if err != nil {
		return nil, err
	}

	err = conn.authenticate()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Connection) dial() error {
	c.server, _ = c.provider.Next()
	conn, err := net.Dial("tcp", c.server)
	if err != nil {
		return err
	}

	c.conn = conn

	return nil
}

func (c *Connection) authenticate() error {
	sendSlice := make([]byte, 256)
	sendBuf := &bytes.Buffer{}

	// create and encode request for zk server
	request := proto.ConnectRequest{
		ProtocolVersion: defaultProtocolVersion,
		LastZxidSeen:    c.lastZxid,
		TimeOut:         int32(c.sessionTimeout.Milliseconds()),
		SessionId:       c.sessionID,
		Passwd:          c.passwd,
	}
	enc := jute.NewBinaryEncoder(sendBuf)
	if err := request.Write(enc); err != nil {
		return err
	}

	// write bytes of request to buf
	requestLen := sendBuf.Len()
	copy(sendSlice[4:], sendBuf.Bytes())
	binary.BigEndian.PutUint32(sendSlice[:4], uint32(requestLen))

	// send request bytes from sendSlice via net.conn
	c.conn.Write(sendSlice[:requestLen+4])

	// receive bytes from same socket, reading the message length first
	recvBuf := make([]byte, 256)
	_, err := io.ReadFull(c.conn, recvBuf[:4])
	if err != nil {
		return err
	}

	recvReader := bytes.NewReader(recvBuf)
	dec := jute.NewBinaryDecoder(recvReader)

	responseLength, err := dec.ReadInt()
	if err != nil {
		return err
	}
	if cap(recvBuf) < int(responseLength) {
		recvBuf = make([]byte, responseLength)
	}

	_, err = io.ReadFull(c.conn, recvBuf[:responseLength])
	if err != nil {
		return err
	}

	response := proto.ConnectResponse{}
	recvReader = bytes.NewReader(recvBuf)
	dec = jute.NewBinaryDecoder(recvReader)
	if err = response.Read(dec); err != nil {
		return err
	}

	if response.SessionId == 0 {
		c.sessionID = 0
		c.passwd = emptyPassword
		c.lastZxid = 0
		return ErrSessionExpired
	}

	c.sessionID = response.SessionId
	c.sessionTimeout = time.Duration(response.TimeOut) * time.Millisecond
	c.passwd = response.Passwd

	return nil
}
