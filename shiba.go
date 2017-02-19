package main

import "fmt"

func main() {
	rect := "\u2588"
	printMost(rect)
	printMore(rect)
	printMedium(rect)
	printLess(rect)
	printNone(rect)
}

func printMost(txt string) {
	printColored(txt, 22)
}

func printMore(txt string) {
	printColored(txt, 28)
}

func printMedium(txt string) {
	printColored(txt, 34)
}

func printLess(txt string) {
	printColored(txt, 40)
}

func printNone(txt string) {
	printColored(txt, 250)
}

func printColored(txt string, colorCode int16) {
	fmt.Printf("\033[38;5;%dm%s\033[0m\n", colorCode, txt)
}
