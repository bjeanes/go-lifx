package main

import (
  "github.com/bjeanes/go-lifx/protocol"
  "fmt"
)

func main() {
  conn, err := protocol.Connect()

  if err != nil { panic(err.Error()) }

  msgs, errs := protocol.NewMessageDecoder(conn.Datagrams)

  for {
    select {
    case msg := <-msgs:
      fmt.Printf("%v\n", &msg)
    case err := <-errs:
      switch e := err.(type) {
      case protocol.BadDatagram:
        fmt.Printf("Error (%s) decoding datagram: %+v", e.Error(), e.Datagram)
      default:
        // ignore
      }
    }
  }
}
