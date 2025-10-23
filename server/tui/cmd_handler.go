package tui

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"quic-c2/models"
	"strings"
)

// SendCMD sends a JSON-encoded command to the selected device over its stream.
//
//	it marshals the command, and writes it to the device stream.
//
// then it runs goroutine PrintDeviceOutput to monitor for incoming output from that stream
func SendCMD(selectedDevice, cmd string) error {
	device, ok := models.Manager.Get(selectedDevice)
	if !ok {
		return fmt.Errorf("device %q not found", selectedDevice)
	}

	parts := strings.Fields(cmd)
	command := models.Command{
		Action: parts[0],
		Args:   parts[1:],
	}

	data, err := json.Marshal(command)
	if err != nil {
		return fmt.Errorf("marshal command: %w", err)
	}
	data = append(data, '\n')

	if _, err := device.Stream.Write(data); err != nil {
		return fmt.Errorf("write stream: %w", err)
	}

	go PrintDeviceOutput(device, cmd)

	return nil
}

// PrintDeviceOutput continuously reads and displays a device's output
// on the CommandView in real-time.
func PrintDeviceOutput(device models.Device, cmd string) {
	scanner := bufio.NewScanner(device.Stream)

	fmt.Fprintf(models.CommandView, "[yellow]%s > [white]%s output: \n", device.Name, cmd)
	for scanner.Scan() {
		fmt.Fprintf(models.CommandView, "%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "read error from %s: %v\n", device.Name, err)
	}
}
