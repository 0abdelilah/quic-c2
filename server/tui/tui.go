package tui

import (
	"os"
	"quic-c2/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// DrawPage creates the main UI layout (device list + command view).
func DrawPage(input *tview.InputField) (*tview.Application, *tview.TextView, *tview.List) {
	app := tview.NewApplication()

	// Left panel: list of connected clients.
	deviceList := tview.NewList().
		ShowSecondaryText(false)
	deviceList.SetBorder(true).
		SetTitle("Clients")

	// Selecting a device focuses the input for typing commands.
	deviceList.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		models.Mu.Lock()
		models.SelectedDevice = mainText
		models.Mu.Unlock()
		app.SetFocus(input)
	})

	// Right panel: command output area.
	commandView := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)

	commandView.SetBorder(true).SetTitle("Command Page")
	commandView.SetChangedFunc(func() { app.Draw() }) // auto-refresh

	// Keyboard shortcuts for navigation and quitting.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyLeft:
			// Switch focus back to device list.
			app.SetFocus(deviceList)
			return nil
		case tcell.KeyCtrlC:
			// Graceful exit.
			os.Exit(0)
		}
		return event
	})

	return app, commandView, deviceList
}
