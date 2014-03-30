package main

import (
  "./protocol"
  "fmt"
  "net"
  "io"
)

const gatewayPort = 56700

func main() {
  reader, writer := io.Pipe()

  go func() {
    socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
      IP:   net.IPv4(255, 255, 255, 255),
      Port: gatewayPort,
    })
    if err != nil {
      panic("Failed to broadcast")
    }
    defer socket.Close()

    io.CopyN(writer, socket, 128)
  }()

  msgs := protocol.NewMessageDecoder(reader)

  for msg := range(msgs) {
    fmt.Printf("%#v\n", msg)
  }
}
