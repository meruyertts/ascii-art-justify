package main

import (
	"ascii-art-justify/checks"
	"ascii-art-justify/splitprint"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := os.Args[1:]
	if len(input) != 3 {
		log.Println("Usage: go run . [STRING] [BANNER] [OPTION]\nEX: go run . something standard --align=right")
		return
	}
	myStr := input[0]
	fileName := checks.FileNameCheck(input[1])
	if !checks.TxtFileCheck(fileName) {
		return
	}
	align := "left"
	a, err := checks.MyAlign(strings.ToLower(input[2]))
	if err != nil {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\nEX: go run . something standard --align=right")
		return
	}
	align = a

	if !checks.IsASCII(myStr) {
		log.Println("non-ASCII character was entered")
		return
	}
	if len(myStr) == 0 {
		return
	}
	if fileName == "" {
		log.Println("only following templates names (standard, shadow, thinkertoy) are available")
		return
	}
	if checks.LineCounter(fileName) != nil {
		return
	}
	splitprint.SplitWord(myStr, fileName, align)
}
