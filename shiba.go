package main

import "os"
import "encoding/xml"
import "io/ioutil"
import "net/http"
import "fmt"

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
	SHIBA_GITHUB_COLOR_MOST   = "#1e6823"
	SHIBA_GITHUB_COLOR_MORE   = "#44a340"
	SHIBA_GITHUB_COLOR_MEDIUM = "#8cc665"
	SHIBA_GITHUB_COLOR_LESS   = "#d6e685"
	SHIBA_GITHUB_COLOR_NONE   = "#eeeeee"
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

func main() {
	userName := "0gajun"
	svgShiba := new(SvgShiba)
	shibaSvgStr := getShibaSvgStr(userName)
	if err := xml.Unmarshal([]byte(shibaSvgStr), svgShiba); err != nil {
		fmt.Println("XML Unmarshal error: ", err)
		return
	}
	shiba := svgToShiba(svgShiba)
	printShiba(userName, shiba)
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

	fmt.Errorf("Unknown color!!!! : %s\n", color)
	os.Exit(-1)
	return SHIBA_TYPE_UNDEFINED
}

func getShibaSvgStr(usr string) string {
	response, _ := http.Get("https://github.com/users/" + usr + "/contributions")
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return string(body)
}

func printShiba(userName string, shiba Shiba) {
	fmt.Printf("User: %s\n\n", userName)
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
