package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func main() {
	flag.Parse()
	args := flag.Args()

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start command. ARGS: %q. ERROR: %+v\n", args, err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatalf("Failed to wait for command. ARGS: %q. ERROR: %+v\n", args, err)
	}
}
