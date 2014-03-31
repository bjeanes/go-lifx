package main

import (
  "github.com/bjeanes/go-lifx/protocol"
  "fmt"
)

func main() {
  conn, err := protocol.Connect()

  if err != nil { panic(err.Error()) }

  msgs := protocol.NewMessageDecoder(conn.Datagrams)

  for msg := range(msgs) {
    fmt.Printf("%v\n", &msg)
  }
}
