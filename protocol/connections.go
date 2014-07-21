package protocol

import (
	"net"
	"runtime"
)

var (
	broadcastIP = net.IPv4(224, 0, 0, 1)
	listenIP    = net.IPv4(0, 0, 0, 0)
)

const (
	broadcastPort   = 56700
	peerPort        = 56750 // not yet needed
	defaultReadSize = 128
)

type Connection struct {
	Datagrams chan Datagram
	connected bool
	sockets   struct {
		read, write *net.UDPConn
	}
}

type Datagram struct {
	From net.Addr
	Data []byte
}

func (conn Connection) IsConnected() bool {
	return conn.connected
}

func (conn Connection) Close() (err error) {
	if !conn.IsConnected() {
		return
	}

	err = conn.sockets.read.Close()
	if err != nil {
		return
	}
	err = conn.sockets.write.Close()
	if err != nil {
		return
	}

	close(conn.Datagrams)

	conn.connected = false
	return
}

func (conn *Connection) setupSockets() (err error) {
	write, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   broadcastIP,
		Port: broadcastPort,
	})
	if err != nil {
		return
	}

	read, err := net.ListenMulticastUDP("udp4", nil, &net.UDPAddr{
		IP:   broadcastIP,
		Port: broadcastPort,
	})
	if err != nil {
		return
	}

	conn.sockets.write = write
	conn.sockets.read = read

	return
}

func Connect() (*Connection, error) {
	conn := &Connection{
		Datagrams: make(chan Datagram),
	}

	err := conn.setupSockets()
	if err == nil {
		go func() {
			b := make([]byte, defaultReadSize)

			for {
				n, addr, err := conn.sockets.read.ReadFrom(b)

				if err != nil {
					close(conn.Datagrams)
					break
				}

				conn.Datagrams <- Datagram{addr, b[0:n]}
			}
		}()
	}

	runtime.SetFinalizer(conn, func(c *Connection) {
		c.Close()
	})

	return conn, err
}
