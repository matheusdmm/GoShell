package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	USER            string = "matheu" //	Fake username
	HOSTNAME        string = "mfckr"  //	Fake hostname
	OPERATOR_SYMBOL        = ">~ "    //	Editable operator symbol
)

var ErrNoPath = errors.New("error: path required")

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(USER, "@", HOSTNAME, OPERATOR_SYMBOL)
		//	Read input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		//	Remove the newline char
		input = strings.TrimSuffix(input, "\n")
		//skip empty input
		if input == "" {
			continue
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	args := strings.Split(input, " ")

	//	Args, mfuckr
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return ErrNoPath
		}

	case "ls":
		if len(args) < 2 {
			return ErrNoPath
		}

	case "git":
		if len(args) < 3 {
			return ErrNoPath
		}

	case "exit":
		os.Exit(0)
	}

	//	Ready to exec command
	cmd := exec.Command(args[0], args[1:]...)

	//	Correct output
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
