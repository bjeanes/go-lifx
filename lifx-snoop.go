package main

import (
	"fmt"
	"github.com/bjeanes/go-lifx/protocol"
)

func main() {
	conn, err := protocol.Connect()

	if err != nil {
		panic(err.Error())
	}

	msgs, errs := protocol.NewMessageDecoder(conn.Datagrams)

	for {
		select {
		case msg := <-msgs:
			fmt.Printf("%v\n", &msg)
		case err := <-errs:
			fmt.Println(err.Error())
		}
	}
}
