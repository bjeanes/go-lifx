package main

import (
	"fmt"
	"github.com/bjeanes/go-lifx/protocol"
	"github.com/fatih/color"
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

	for {
		b := <-bytes
		select {
		case msg := <-msgs:
			out := color.New(color.FgGreen)
			out.Printf("DATA: %v\n", b)
			out.Add(color.Bold).Printf("MSG:  %v", &msg)
		case err := <-errs:
			out := color.New(color.FgRed)
			out.Printf("DATA: %v\n", b)
			out.Add(color.Bold).Printf("ERR:  %s", err.Error())
		}
		fmt.Printf("\n\n")
	}
}
