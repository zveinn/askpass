package main

import (
	"bytes"
	"os"
	"os/exec"
)

var prompt string

func main() {
	if len(os.Args) < 2 || os.Args[0] == "" {
		os.Exit(1)
	}
	not()
	if finger() {
		cmd := exec.Command("sudo", "bash")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		os.Stdout.Sync()
		os.Stdin.Sync()
		os.Stderr.Sync()
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		return

	}

	os.Exit(1)
}

func finger() bool {
	cmd := exec.Command("fprintd-verify", os.Args[1])
	out, err := cmd.CombinedOutput()
	if err == nil {
		if bytes.Contains(out, []byte("verify-match (done)")) {
			return true
		}
	}
	return false
}
