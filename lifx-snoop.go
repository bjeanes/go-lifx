package main

import (
	"fmt"
	"github.com/bjeanes/go-lifx/protocol"
	"github.com/fatih/color"
)

// Print byte array as a grid of hex numbers
func printBytes(bytes []byte) {
	out := color.New(color.FgWhite)
	out.Printf("DATA: length=%d (%04x)", len(bytes), len(bytes))
	for i, b := range bytes {
		if i%16 == 0 {
			out.Println()
			out.Print("      ")
		}

		out.Printf("%02x ", b)
	}
	out.Println()
}

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
		printBytes(b)
		color.Set(color.Bold)
		select {
		case msg := <-msgs:
			color.Set(color.FgGreen)
			fmt.Printf("MSG:  %v", &msg)
		case err := <-errs:
			color.Set(color.FgRed)
			fmt.Printf("ERR:  %s", err.Error())
		}
		color.Unset()
		fmt.Printf("\n\n")
	}
}
