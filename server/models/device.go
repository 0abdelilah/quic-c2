package models

import (
	"sync"

	"github.com/quic-go/quic-go"
)

type Device struct {
	Name   string
	Conn   *quic.Conn
	Stream *quic.Stream
}

type Command struct {
	Action string   `json:"action"`
	Args   []string `json:"args"`
}

var Mu sync.Mutex
