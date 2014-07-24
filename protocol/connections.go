package protocol

import (
	"net"
	"runtime"
)

const (
	broadcastPort   = 56700
	peerPort        = 56750 // not yet used
	defaultReadSize = 128
)

type Connection struct {
	Datagrams chan Datagram
	connected bool
	lastErr   error
	sockets   struct {
		read, write *net.UDPConn
	}
}

type Datagram struct {
	From net.Addr
	Data []byte
}

func (conn Connection) LastError() error {
	return conn.lastErr
}

func (conn Connection) IsError() bool {
	return conn.LastError() != nil
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
	// NOTE(bo): On the IP address used for send and receive connections.
	//
	// Go only sets SO_REUSEADDR and SO_REUSEPORT when using a Multicast
	// IP address[1][2][3]. Without dropping down to C, there isn't a way great around
	// this.
	//
	// Despite not being the 255.255.255.255 used in other LIFX libraries, In
	// practice, it seems to work for receiving messages. Sending messages is
	// still somewhat unverified.
	//
	// I may be confusing multicast and broadcast a bit here, but I don't know how else to
	// enable binding to the same interface and port multiple times...
	//
	// [1]: http://golang.org/src/pkg/net/sock_posix.go?h=setDefaultMulticastSockopts#L161
	// [2]: http://golang.org/src/pkg/net/sockopt_bsd.go (also ./sockopt_linux.go)
	// [3]: http://en.wikipedia.org/wiki/Multicast_address#Local_subnetwork
	ip := net.IPv4(224, 0, 0, 1)

	// We may be able to write directly to the below multicast socket anyway, in which
	// case this may not be needed. Not tested yet.
	write, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   ip,
		Port: broadcastPort,
	})
	if err != nil {
		return
	}

	read, err := net.ListenMulticastUDP("udp4", nil, &net.UDPAddr{
		IP:   ip,
		Port: broadcastPort,
	})
	if err != nil {
		return
	}

	conn.sockets.write = write
	conn.sockets.read = read

	return
}

// Starts UDP connections to send and receive datagrams. Returns a Connection struct
// which contains a channel that should be used to receive new UDP packets.
//
// The connection channel will be closed on a socket error. The error can be retrieved
// with the LastError() method on the connection.
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
				conn.lastErr = err

				if conn.IsError() {
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
