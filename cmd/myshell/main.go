package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"errors"
	"strconv"
	"path"
	"path/filepath"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		cmd, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		cmd = strings.TrimRight(cmd, "\n")

		handleCommand(cmd)
	}
}

func findExecutable(name string) (string, bool) {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, path := range paths {
		fullpath := filepath.Join(path, name)
		if _, err := os.Stat(fullpath); err == nil {
			return fullpath, true
		}
	}
	if _, err := os.Stat(name); err == nil {
		return name, true
	}
	return "", false
}

func runCommand(cmd string, args []string) {
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: command not found\n", cmd)
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
				if path, ok := findExecutable(cmdList[1]); ok {
					fmt.Printf("%s is %s\n", cmdList[1], path)
				} else {
					fmt.Printf("%s not found\n", cmdList[1])
				}
			}
		},
		"pwd": func() {
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("error reading working directory")
				return
			}
			fmt.Println(dir)
		},
		"cd": func() {
			var dir string
			if len(cmdList[1]) == 0 {
				dir = getHomeDir()
			} else {
				if cmdList[1] == "~" {
					dir = getHomeDir()
				} else {
					dir = cmdList[1]
				}
			}
			p := path.Clean(dir)
			if !path.IsAbs(p) {
				dir, _ := os.Getwd()
				p = path.Join(dir, p)
			}
			err := os.Chdir(p)
			if err != nil {
				fmt.Printf("%s: No such file or directory\n", p)
			}
		},
	}
	fn, ok := commands[cmdList[0]]
	if !ok {
		if _, ok := findExecutable(cmdList[0]); ok {
			runCommand(cmdList[0], cmdList[1:])
		} else {
			fmt.Printf("%s: command not found\n", cmdList[0])
		}
		return
	}
	fn()
}

func getHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	return home
}