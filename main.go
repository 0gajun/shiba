package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Println("NAME:")
	fmt.Println("\tshiba - The viewer of GitHub's contribution graph")
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println("\tshiba [username]")
	fmt.Println()
	fmt.Println("ARGUMENT:")
	fmt.Println("\tusername\tGitHub's username(You can also specify this using environment variable, 'SHIBA_GITHUB_USER_NAME')")
	fmt.Println()
	fmt.Println("AUTHOR:")
	fmt.Println("\t0gajun <oga.ivc.s27@gmail.com>")
	fmt.Println()
	fmt.Println("GLOBAL OPTIONS:")
	fmt.Println("\t--help, -h\tshow help")
	fmt.Println()
	fmt.Println("")
}

func main() {
	var userName = ""
	var isHelp = false

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.BoolVar(&isHelp, "help", false, "show help")
	f.BoolVar(&isHelp, "h", false, "show help")

	f.Parse(os.Args[1:])
	for 0 < f.NArg() {
		f.Parse(f.Args()[1:])
	}

	if isHelp {
		usage()
		return
	}

	if len(os.Args) == 2 {
		userName = os.Args[1]
	} else if shibaGithubUserName := os.Getenv("SHIBA_GITHUB_USER_NAME"); shibaGithubUserName != "" {
		userName = shibaGithubUserName
	}

	if userName == "" {
		usage()
		return
	}

	Show(userName)
}
