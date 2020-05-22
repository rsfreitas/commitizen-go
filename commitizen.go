package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var args Arguments
	parseArgs(&args)

	// open debug switch
	if args.debug {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
	} else {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	if args.install {
		if path, err := Install(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Install commitizen to %s\n", path)
		}
	} else {
		// exit if not git directory
		ExitIfNotGitDirectory()
		answers := newAnswers(args)

		// ask the question
		if err := AskForCommitMessage(answers); err != nil {
			log.Fatal(err)
		}

		// assemble the answers to commit message
		var buf bytes.Buffer
		answers.AssembleIntoMessage(&buf)

		proceed, err := ConfirmCommit(&buf, answers.HasBody())
		if err != nil {
			log.Fatal(err)
		}
		if proceed {
			// do git commit
			result, err := CommitMessage(buf.Bytes(), args.all)
			if err != nil {
				log.Printf("run git commit failed, \n")
				log.Printf("commit message is: \n\n\t%s\n\n", buf.String())
			}
			fmt.Print(result)
		}
	}
}
