package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
	"strconv"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		cmd, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		cmd = strings.TrimRight(cmd, "\n")

		handleCommand(cmd)
	}
}

func handleCommand(cmd string) {
	cmdList := strings.Split(cmd, " ")

	commands := map[string]func(){
		"exit": func() {
			code, err := strconv.Atoi(cmdList[1])
			if errors.Is(err, strconv.ErrSyntax) {
				os.Exit(0)
			}
			os.Exit(code)
		},
		"echo": func() {
			fmt.Println(strings.Join(cmdList[1:], " "))
		},
		"type": func() {
			switch cmdList[1] {
			case "exit", "echo", "type":
				fmt.Printf("%s is a shell builtin\n", cmdList[1])
			default:
				fmt.Printf("%s not found\n", cmdList[1])
			}
		},
	}
	fn, ok := commands[cmdList[0]]
	if !ok {
		fmt.Printf("%s: command not found\n", cmd)
		return
	}
	fn()
}
