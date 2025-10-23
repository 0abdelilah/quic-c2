package tui

import (
	"log"
	"os"

	"golang.org/x/term"
)

func GetTermWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal("Error:", err)
		return 0
	}
	return width
}
