package main

import (
	"errors"
	"io"
	"net"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// ConnHubStatus reflects the status of the ConnHubs state.
type ConnHubStatus int

const (
	// Dead
	Offline ConnHubStatus = iota

	// StartupRequested state means the startup sequence has been initiated
	// by user.
	StartupRequested

	// Initializing and preparing environment
	PreparingSession

	// Stable and ready for client connections
	AcceptingConnections

	// Busy can also be though of as NotAcceptingConnections
	Busy

	// ShutdownRequested by user
	ShutdownRequested

	// ShuttingDown closes client connections and prepares for gracefull
	// shutdown.
	ShuttingDown
)

// ConnHub is a server that hosts the jam session, and clients join by
// connecting and registering their streams.
type ConnHub struct {
	HostAddr string
	state    ConnHubStatus
	cc       chan net.Listener
}

func NewConnHub() *ConnHub {
	return &ConnHub{state: Offline, cc: make(chan net.Listener)}
}

func (ch *ConnHub) Status() ConnHubStatus {
	return ch.state
}

func (ch *ConnHub) Network() string {
	return ch.HostAddr[strings.Index(ch.HostAddr, ":"):]
}

func (ch *ConnHub) String() string {
	return ch.HostAddr
}

func (ch *ConnHub) Addr() net.Addr {
	return ch
}

func (ch *ConnHub) ListenAndServe(network, addr string) error {
	ch.state = PreparingSession
	ip, err := net.ResolveIPAddr(network, addr)
	if err != nil {
		return err
	}
	ch.HostAddr = ip.String()
	c, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	ch.cc <- c
	ch.state = AcceptingConnections
	return nil

}

func (ch *ConnHub) Accept() (net.Conn, error) {
	if ch.state != AcceptingConnections {

		return nil, errors.New("error inconsistent state: ch.Status() != AcceptingConnections")
	}

	if c, ok := <-ch.cc; ok {
		return c.Accept()
	}
	return nil, errors.New("unknown error")
}

// ConnStream is an abstract representation of a client connections signal
// streams.
type ConnStream struct {
	ToConn   io.ReadWriteCloser
	FromConn io.ReadWriteCloser
}

// JamSession is an active session, with zero or more participants.
type JamSession struct {
	// number of participants in the session
	ClientCount int

	// Each conn gets assigned a session uuid which is used as a key for
	// their signal io.
	ClientConns map[uuid.UUID]*ConnStream
}
