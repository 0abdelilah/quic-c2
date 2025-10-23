package commands

import (
	"log"
	"os"
	"strings"
)

func ListDir(dir string) string {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	names := []string{}

	for _, f := range files {
		names = append(names, f.Name())
	}

	return strings.Join(names, "\n")
}
