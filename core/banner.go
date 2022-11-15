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

	putAsciiArt("__                                     __\n")
	putAsciiArt("_   @@     @@@@@@@@@@@@@@@@@@@     @@   _")
	printLogo(`    ___________      __ __           __               `)
	fmt.Println()
	putAsciiArt("  @@@@    @@@@@@@@@@@@@@@@@@@@@    @@@@  ")
	printLogo(`    \_   _____/__  _|__|  |    ____ |__| ____ ___  ___`)
	fmt.Println()
	putAsciiArt("  @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@  ")
	printLogo(`     |    __)_\  \/ /  |  |   / __ \|  |/    \\  \/  /`)
	fmt.Println()
	putAsciiArt("    @@@@@@@@@@###@@@@@@@###@@@@@@@@@@    ")
	printLogo(`     |        \\   /|  |  |__/ /_/  >  |   |  \>    < `)
	fmt.Println()
	putAsciiArt("      @@@@@@@#####@@@@@#####@@@@@@@      ")
	printLogo(`    /_______  / \_/ |__|____/\___  /|__|___|  /__/\_ \`)
	fmt.Println()
	putAsciiArt("       @@@@@@@###@@@@@@@###@@@@@@@       ")
	printLogo(`            \/              /_____/         \/      \/`)
	fmt.Println()
	putAsciiArt("      @@@@@@@@@@@@@@@@@@@@@@@@@@@@@      \n")
	putAsciiArt("     @@@@@WW@@@WW@@WWW@@WW@@@WW@@@@@     ")
	printUpdateName()
	fmt.Println()
	putAsciiArt("    @@@@@@WW@@@WW@@WWW@@WW@@@WW@@@@@@    \n")
	//printOneliner2()
	//fmt.Println()
	putAsciiArt("_   @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@   _")
	printOneliner1()
	fmt.Println()
	putAsciiArt("__                                     __\n")
	fmt.Println()
}
