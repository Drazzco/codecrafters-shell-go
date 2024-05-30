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

		fmt.Printf("%s: command not found\n", strings.TrimRight(in, "\n"))
	}
}
