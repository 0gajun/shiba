package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/briandowns/spinner"
)

const DAY_OF_WEEK = 7

const (
	SHIBA_TYPE_MOST ShibaType = iota + 1
	SHIBA_TYPE_MORE
	SHIBA_TYPE_MEDIUM
	SHIBA_TYPE_LESS
	SHIBA_TYPE_NONE
	SHIBA_TYPE_UNDEFINED
)

const (
	SHIBA_GITHUB_COLOR_MOST   = "#196127"
	SHIBA_GITHUB_COLOR_MORE   = "#239a3b"
	SHIBA_GITHUB_COLOR_MEDIUM = "#7bc96f"
	SHIBA_GITHUB_COLOR_LESS   = "#c6e48b"
	SHIBA_GITHUB_COLOR_NONE   = "#ebedf0"
)

const (
	SHIBA_TERMINAL_COLOR_CODE_MOST   = 22
	SHIBA_TERMINAL_COLOR_CODE_MORE   = 28
	SHIBA_TERMINAL_COLOR_CODE_MEDIUM = 34
	SHIBA_TERMINAL_COLOR_CODE_LESS   = 40
	SHIBA_TERMINAL_COLOR_CODE_NONE   = 250
)

type ShibaType uint8
type Shiba [][]ShibaType

type SvgShiba struct {
	Columns []SvgShibaColumn `xml:"g>g"`
}

type SvgShibaColumn struct {
	Rects []SvgShibaRect `xml:"rect"`
}

type SvgShibaRect struct {
	Count int64  `xml:"data-count,attr"`
	Date  string `xml:"data-date,attr"`
	Color string `xml:"fill,attr"`
}

func Show(userName string, timeZone string) int {
	sp := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	sp.Suffix = " Fetching shiba..."
	sp.Prefix = " "
	sp.Start()

	shibaObj, err := loadShiba(userName, timeZone)

	sp.Stop()

	if err != nil {
		fmt.Println(err)
		return 1
	}

	fmt.Println()
	fmt.Printf("TimeZone: %s\n", timeZone)
	fmt.Printf("User:     %s\n\n", userName)
	printShiba(userName, shibaObj)

	return 0
}

func loadShiba(userName string, timeZone string) (Shiba, error) {
	shibaSvgStr, err := getShibaSvgStr(userName, timeZone)

	if err != nil {
		return nil, fmt.Errorf("Cannot get contribution data\nUnknown user: %s", userName)
	}

	svgShiba := new(SvgShiba)
	if err := xml.Unmarshal([]byte(shibaSvgStr), svgShiba); err != nil {
		return nil, fmt.Errorf("XML Unmarshal error: ", err)
	}

	return svgToShiba(svgShiba), nil
}

func svgToShiba(svgShiba *SvgShiba) Shiba {
	colSize := len(svgShiba.Columns)
	shiba := newEmptyShiba(colSize)

	for colIndex, col := range svgShiba.Columns {
		for rowIndex, row := range col.Rects {
			shiba[rowIndex][colIndex] = detectShibaType(row.Color)
		}
	}

	return shiba
}

func newEmptyShiba(colSize int) Shiba {
	shiba := make(Shiba, DAY_OF_WEEK)
	for i := range shiba {
		shiba[i] = make([]ShibaType, colSize)
		for j := range shiba[i] {
			shiba[i][j] = SHIBA_TYPE_UNDEFINED
		}
	}
	return shiba
}

func detectShibaType(color string) ShibaType {
	switch color {
	case SHIBA_GITHUB_COLOR_MOST:
		return SHIBA_TYPE_MOST
	case SHIBA_GITHUB_COLOR_MORE:
		return SHIBA_TYPE_MORE
	case SHIBA_GITHUB_COLOR_MEDIUM:
		return SHIBA_TYPE_MEDIUM
	case SHIBA_GITHUB_COLOR_LESS:
		return SHIBA_TYPE_LESS
	case SHIBA_GITHUB_COLOR_NONE:
		return SHIBA_TYPE_NONE
	}

	fmt.Fprintf(os.Stderr, "Unknown Shiba color : %s\n", color)
	fmt.Fprintf(os.Stderr, "Is today special day? (like Haloween)\n")
	fmt.Fprintf(os.Stderr, "If today is special day, the GitHub sometimes change shiba colors\n")
	fmt.Fprintf(os.Stderr, "So, please create issue and paste shiba color codes!\n")
	fmt.Fprintf(os.Stderr, "-> https://github.com/0gajun/shiba/issues\n")
	os.Exit(-1)
	return SHIBA_TYPE_UNDEFINED
}

func getShibaSvgStr(usr string, timeZone string) (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://github.com/users/"+usr+"/contributions", nil)
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Cookie", "tz="+timeZone)
	response, err := client.Do(req)

	if err != nil || response.StatusCode != 200 {
		return "", fmt.Errorf("Cannot get contribution data")
	}

	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return string(body), nil
}

func printShiba(userName string, shiba Shiba) {
	for _, row := range shiba {
		for _, shibaType := range row {
			printShibaRect(shibaType)
		}
		fmt.Println()
	}
	resetColor()
}

func printShibaRect(shibaType ShibaType) {
	switch shibaType {
	case SHIBA_TYPE_MOST:
		printColoredRect(SHIBA_TERMINAL_COLOR_CODE_MOST)
	case SHIBA_TYPE_MORE:
		printColoredRect(SHIBA_TERMINAL_COLOR_CODE_MORE)
	case SHIBA_TYPE_MEDIUM:
		printColoredRect(SHIBA_TERMINAL_COLOR_CODE_MEDIUM)
	case SHIBA_TYPE_LESS:
		printColoredRect(SHIBA_TERMINAL_COLOR_CODE_LESS)
	case SHIBA_TYPE_NONE:
		printColoredRect(SHIBA_TERMINAL_COLOR_CODE_NONE)
	case SHIBA_TYPE_UNDEFINED:
		// IGNORED
	}
}

func printColoredRect(colorCode int16) {
	const RECT = "\u2588"
	fmt.Printf("\033[38;5;%dm%s ", colorCode, RECT)
}

func resetColor() {
	fmt.Printf("\033[0m")
}
