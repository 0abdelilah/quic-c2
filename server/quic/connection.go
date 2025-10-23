package server

import (
	"context"
	"fmt"
	"log"
	"quic-c2/models"
	"time"

	"github.com/quic-go/quic-go"
)

const addr = "localhost:4242"

// Listen for incoming connection and update deviceList once a device is connected
// bug: device is not removed on client disconnect
func StartServer() {
	quicConfig := &quic.Config{
		MaxIdleTimeout:  120 * time.Hour,
		KeepAlivePeriod: 30 * time.Second,
	}

	listener, err := quic.ListenAddr(addr, generateTLSConfig(), quicConfig)
	if err != nil {
		log.Fatalf("failed to start listener: %v", err)
	}

	for {
		conn, err := listener.Accept(context.Background())
		if err != nil {
			log.Println("accept error:", err)
			continue
		}

		stream, err := conn.OpenStreamSync(context.Background())
		if err != nil {
			log.Println("stream open failed:", err)
			conn.CloseWithError(0, "stream open failed")
			continue
		}

		device := models.Device{
			Name:   fmt.Sprintf("Device %d", len(models.Manager.Devices)),
			Conn:   conn,
			Stream: stream,
		}

		go models.Manager.Add(device)

		// remove device when it exits, not working for unknown reason
		go models.Manager.Monitor(device)
	}
}
