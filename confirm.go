package main

import (
	"bytes"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func ConfirmCommit(buf *bytes.Buffer) (bool, error) {
	// preview commit
	fmt.Printf("\nCommit message preview:\n\n%s\n\n", buf.String())

	proceed := false
	prompt := &survey.Confirm{
		Message: "Confirm commit?",
	}

	if err := survey.AskOne(prompt, &proceed); err != nil {
		return false, err
	}

	return proceed, nil
}
