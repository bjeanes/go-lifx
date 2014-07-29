package main

import (
	"encoding/hex"
	"fmt"
	"github.com/bjeanes/go-lifx/protocol"
	"github.com/fatih/color"
	"regexp"
	"strings"
)

func main() {
	conn, err := protocol.Connect()

	if err != nil {
		panic(err.Error())
	}

	snoopDatagrams := make(chan protocol.Datagram)
	decoderDatagrams := make(chan protocol.Datagram)

	go func() {
		for d := range conn.Datagrams {
			snoopDatagrams <- d
			decoderDatagrams <- d
		}

		close(snoopDatagrams)
		close(decoderDatagrams)
	}()

	msgs, errs := protocol.NewMessageDecoder(decoderDatagrams)

	re := regexp.MustCompilePOSIX("^00000")

	for {
		d := <-snoopDatagrams

		out := color.New(color.FgWhite)
		out.Printf("FROM: %s\n", d.From)
		out.Printf("SIZE: %d\n", len(d.Data))
		data := strings.TrimLeft(hex.Dump(d.Data), "0")
		out.Printf("DATA: 000" + re.ReplaceAllString(data, "      "))
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
