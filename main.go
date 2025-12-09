package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	// read a line of input from the user
	reader := bufio.NewReader(os.Stdin)

	for {
		// display the prompt
		fmt.Print("ccsh> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			os.Exit(0)
		}
		// run command
		cmd := exec.Command(input)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}

}
