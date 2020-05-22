package main

import (
	"flag"
	"fmt"
	"os"
)

type Arguments struct {
	all     bool
	install bool
	debug   bool
	bullet  bool
}

var revision string

var usage string = `Usage: commitizen 
       or 
       git cz (after install to git-core)

Options:
    -a, -all        Tell the command to automatically stage files that have been modified and deleted, but new files you have not told Git about are not affected
    -i, -install    Install this tool to git-core as git-cz
    -d, -debug      Debug mode
    -v, -version    Print version information and quit
    -b, -bullet     Consider body sentences as bullet points
`

func parseArgs(args *Arguments) {
	var (
		help    bool
		version bool
	)
	flag.BoolVar(&args.all, "a", false, "Tell the command to automatically stage files that have been modified and deleted, but new files you have not told Git about are not affected.")
	flag.BoolVar(&args.all, "all", false, "Tell the command to automatically stage files that have been modified and deleted, but new files you have not told Git about are not affected.")
	flag.BoolVar(&args.install, "i", false, "Install this tool to git-core as git-cz.")
	flag.BoolVar(&args.install, "install", false, "Install this tool to git-core as git-cz.")
	flag.BoolVar(&args.debug, "d", false, "Debug mode.")
	flag.BoolVar(&args.debug, "debug", false, "Debug mode.")
	flag.BoolVar(&help, "h", false, "Show the help.")
	flag.BoolVar(&help, "help", false, "Show the help.")
	flag.BoolVar(&version, "v", false, "")
	flag.BoolVar(&version, "version", false, "")
	flag.BoolVar(&args.bullet, "b", false, "Consider body sentences as bullet points.")
	flag.BoolVar(&args.bullet, "bullet", false, "Consider body sentences as bullet points.")

	flag.Usage = func() {
		fmt.Println(usage)
	}

	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	} else if version {
		fmt.Printf("Commitizen-go version 0.1.0, build revision %s\n", revision)
		os.Exit(0)
	}
}
