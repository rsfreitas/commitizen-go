package main

import (
	"bytes"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

var (
	Black   = color("\033[1;30m%s\033[0m")
	Red     = color("\033[1;31m%s\033[0m")
	Green   = color("\033[1;32m%s\033[0m")
	Yellow  = color("\033[1;33m%s\033[0m")
	Purple  = color("\033[1;34m%s\033[0m")
	Magenta = color("\033[1;35m%s\033[0m")
	Teal    = color("\033[1;36m%s\033[0m")
	White   = color("\033[1;37m%s\033[0m")
)

func color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString, fmt.Sprint(args...))
	}

	return sprint
}

func ConfirmCommit(buf *bytes.Buffer) (bool, error) {
	// preview commit
	fmt.Printf("\n%s\n", Teal("Commit message preview"))
	fmt.Printf("%s\n\n", Teal("======================"))
	fmt.Printf("%s\n\n", buf.String())

	proceed := false
	prompt := &survey.Confirm{
		Message: "Confirm commit?",
	}

	if err := survey.AskOne(prompt, &proceed); err != nil {
		return false, err
	}

	return proceed, nil
}
