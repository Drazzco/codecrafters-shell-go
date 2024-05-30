package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		in, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		cmds := strings.Split(in, " ")

		if cmds[0] == "exit" {
			os.Exit(0)
		} else {
			fmt.Printf("%s: command not found\n", strings.TrimSpace(cmds[0]))
		}
	}
}
