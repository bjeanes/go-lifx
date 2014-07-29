package main

import (
	"encoding/hex"
	"fmt"
	"github.com/bjeanes/go-lifx/protocol"
	"github.com/fatih/color"
	"os"
	"regexp"
	"strings"
)

func main() {
	filter := ""

	if len(os.Args) >= 2 {
		filter = strings.ToLower(os.Args[1])
	}

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

	for {
		d := <-snoopDatagrams

		out := color.New(color.FgWhite)

		select {
		case msg := <-msgs:
			t := strings.ToLower(fmt.Sprintf("%T", msg.Payload))
			if !strings.Contains(t, filter) {
				break
			}
			printDatagram(out, &d)
			out.Add(color.FgGreen)
			out.Printf("MSG:  %+v\n      %T %+v", msg.Header, msg.Payload, msg.Payload)
			fmt.Printf("\n\n")
		case err := <-errs:
			printDatagram(out, &d)
			out.Add(color.FgRed)
			out.Printf("ERR:  %s", err.Error())
			fmt.Printf("\n\n")
		}
	}
}

func printDatagram(out *color.Color, d *protocol.Datagram) {
	re := regexp.MustCompilePOSIX("^00000")

	out.Printf("FROM: %s\n", d.From)
	out.Printf("SIZE: %d\n", len(d.Data))
	data := strings.TrimLeft(hex.Dump(d.Data), "0")
	out.Printf("DATA: 000" + re.ReplaceAllString(data, "      "))
	color.Set(color.Bold)
}
