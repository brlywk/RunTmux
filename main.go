package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// get tmux dir from arguments
	var tmuxDir string
	var sessionName string

	flag.StringVar(&tmuxDir, "d", "/opt/homebrew/bin/tmux", "Specifies the directory where tmux is installed.")
	flag.StringVar(&sessionName, "s", "default", "Name of session to connect to by default.")
	flag.Parse()

	_, err := exec.LookPath(tmuxDir)
	if err != nil {
		fmt.Printf("Unable to find tmux executable in: %v\n", tmuxDir)
		os.Exit(1)
	}

	// we always want tmux to start / attach to a 'default' session
	tmux := exec.Command(tmuxDir, "attach", "-t", sessionName)
	tmux.Stdin = os.Stdin
	tmux.Stdout = os.Stdout
	tmux.Stderr = os.Stderr

	// if no default session is found, create one
	if err := tmux.Run(); err != nil {
		fmt.Printf("Error:\t%v", err)

		tmuxNewDefault := exec.Command(tmuxDir, "new-session", "-s", sessionName)
		tmuxNewDefault.Stdin = os.Stdin
		tmuxNewDefault.Stdout = os.Stdout
		tmuxNewDefault.Stderr = os.Stderr

		if err := tmuxNewDefault.Run(); err != nil {
			fmt.Printf("Unable to attach to or create 'default' session.\nError:\t%v", err)
			os.Exit(1)
		}
	}
}
