package protocol

import (
	"io"
	"net"
)

var (
	broadcastIP = net.IPv4(255, 255, 255, 255)
	listenIP    = net.IPv4(0, 0, 0, 0)
)

const (
	broadcastPort   = 56700
	peerPort        = 56750 // not yet needed
	defaultReadSize = 128
)

type Connection struct {
	Datagrams chan datagram
	sockets   struct {
		read, write *net.UDPConn
	}
}

type datagram struct {
	From net.Addr
	Data []byte
}

func (conn Connection) Close() (err error) {
	err = conn.sockets.read.Close()
	if err != nil {
		return
	}

	err = conn.sockets.write.Close()
	return
}

func (conn Connection) writeTo(writer io.WriteCloser) {
	go func() {
		for {
			_, err := io.CopyN(writer, conn.sockets.read, 128)
			if err != nil {
				writer.Close()
				conn.sockets.read.Close()
				break
			}
		}
	}()
}

func (conn *Connection) setupSockets() (err error) {
	write, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   broadcastIP,
		Port: broadcastPort,
	})
	if err != nil {
		return
	}

	read, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   listenIP,
		Port: broadcastPort,
	})
	if err != nil {
		return
	}

	conn.sockets.write = write
	conn.sockets.read = read

	return
}

func Connect() (conn Connection, err error) {
	conn.Datagrams = make(chan datagram)

	if err := conn.setupSockets(); err == nil {
		go func() {
			b := make([]byte, defaultReadSize)

			for {
				n, addr, err := conn.sockets.read.ReadFrom(b)

				if err != nil {
					close(conn.Datagrams)
					return
				}

				conn.Datagrams <- datagram{addr, b[0:n]}
			}
		}()
	}

	return
}
