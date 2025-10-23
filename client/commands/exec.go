package commands

import (
	"fmt"
	"os/exec"
)

func Exec(args []string) string {
	cmd := exec.Command("cmd", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	return string(output)
}
