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
	fmt.Println("\tshiba [username] [--tz <timezone>]")
	fmt.Println()
	fmt.Println("ARGUMENT:")
	fmt.Println("\tusername\tGitHub's username(You can also specify this using environment variable, 'SHIBA_GITHUB_USER_NAME')")
	fmt.Println()
	fmt.Println("OPTION:")
	fmt.Println("\t--tz\tLocal time zone with IANA TZ Database format, like 'Asia/Tokyo'.")
	fmt.Println("\t\tYou can also specify this using environment variable, 'SHIBA_TIME_ZONE')")
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
	var (
		isHelp   = false
		timeZone = ""
	)

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.BoolVar(&isHelp, "help", false, "show help")
	f.BoolVar(&isHelp, "h", false, "show help")
	f.StringVar(&timeZone, "tz", "", "local time zone")

	f.Parse(os.Args[1:])
	for 0 < f.NArg() {
		f.Parse(f.Args()[1:])
	}

	if isHelp {
		usage()
		return
	}

	if timeZone == "" {
		if envTz := os.Getenv("SHIBA_LOCAL_TIME_ZONE"); envTz != "" {
			timeZone = envTz
		} else {
			timeZone = "Asia/Tokyo"
		}
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

	os.Exit(Show(userName, timeZone))
}
