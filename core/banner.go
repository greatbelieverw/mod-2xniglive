package core

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const (
	VERSION = "0.0.1"
)

func putAsciiArt(s string) {

}

func printLogo(s string) {

}

func printUpdateName() {
	nameClr := color.New(color.FgHiRed)
	txt := nameClr.Sprintf("                 - --  Go Phishing  -- -")
	fmt.Fprintf(color.Output, "%s", txt)
}

func printOneliner1() {
	handleClr := color.New(color.FgHiBlue)
	versionClr := color.New(color.FgGreen)
	textClr := color.New(color.FgHiBlack)
	spc := strings.Repeat(" ", 10-len(VERSION))
	txt := textClr.Sprintf("      by Zhang Wei Li (") + handleClr.Sprintf("@zhangwei-g1t") + textClr.Sprintf(")") + spc + textClr.Sprintf("version ") + versionClr.Sprintf("%s", VERSION)
	fmt.Fprintf(color.Output, "%s", txt)
}

func printOneliner2() {
	textClr := color.New(color.FgHiBlack)
	red := color.New(color.FgRed)
	white := color.New(color.FgWhite)
	txt := textClr.Sprintf("                   no ") + red.Sprintf("nginx") + white.Sprintf(" - ") + textClr.Sprintf("pure ") + red.Sprintf(" mod-2xniglive
")
	fmt.Fprintf(color.Output, "%s", txt)
}

func Banner() {
	fmt.Println()
	fmt.Println()
	printUpdateName()
	fmt.Println()
	printOneliner2()
	fmt.Println()
	printOneliner1()
	fmt.Println()
	fmt.Println()
}
