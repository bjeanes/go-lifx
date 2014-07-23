package main

import (
  "encoding/hex"
  "fmt"
  "github.com/bjeanes/go-lifx/protocol"
  "github.com/fatih/color"
  "regexp"
)

func main() {
  conn, err := protocol.Connect()

  if err != nil {
    panic(err.Error())
  }

  datagrams := make(chan protocol.Datagram)
  bytes := make(chan []byte)

  go func() {
    for d := range conn.Datagrams {
      datagrams <- d
      bytes <- d.Data
    }

    close(datagrams)
    close(bytes)
  }()

  msgs, errs := protocol.NewMessageDecoder(datagrams)

  re := regexp.MustCompilePOSIX("^00000")

  for {
    b := <-bytes

    out := color.New(color.FgWhite)
    out.Printf("DATA: length=%d\n", len(b))
    out.Print(re.ReplaceAllString(hex.Dump(b), "      "))
    color.Set(color.Bold)

    select {
    case msg := <-msgs:
      out.Add(color.FgGreen)
      out.Printf("MSG:  %+v\n      %T %+v", msg.Header, msg.Payload, msg.Payload)
    case err := <-errs:
      out.Add(color.FgRed)
      out.Printf("ERR:  %s", err.Error())
    }
    fmt.Printf("\n\n")
  }
}
