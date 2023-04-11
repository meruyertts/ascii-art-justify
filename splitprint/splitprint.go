package splitprint

import (
	"ascii-art-justify/read"
	"fmt"
	"log"
	"regexp"
	"strings"
)

var (
	termWidth int = read.GetTermWidth()
	lenText   []int
)

func SplitWord(myStr, myFile, align string) {
	myStr = Check(myStr, PrintWordJustify(myStr, myFile, align))
	re := regexp.MustCompile(`\\n`)
	newStr := re.Split(myStr, -1)

	for i := 0; i < len(newStr); i++ {
		if len(newStr[i]) > 0 {
			PrintWord(newStr[i], myFile, align)
		}
		if newStr[i] == "" {
			fmt.Println("")
		}
	}
}

func PrintWord(s, fileName, align string) {
	myLine := ""
	myarray := [8]string{}
	for _, char := range s {
		for line := 2; line <= 9; line++ {
			myrune := int(char)
			for i := ' '; i <= '~'; i++ {
				j := (int(i) - ' ') * 9
				if myrune == int(i) {
					firstline, err := read.ReadExactLine(fileName, line+j)
					if err != nil {
						log.Print(err)
						return
					}
					myLine += firstline
				}
			}
		}
		temp := strings.Split(myLine, "\n")
		for index, s := range temp[:len(temp)-1] {
			myarray[index] += s
		}
		myLine = ""
	}
	PrintAlign(myarray, align, s, fileName)
}

func PrintAlign(s [8]string, align, input, filename string) {
	if align == "left" {
		for _, i := range s {
			fmt.Println(i)
		}
	} else if align == "right" {
		for _, i := range s {
			fmt.Print(printSpaces(termWidth - len(i)))
			fmt.Print(i)
		}
	} else if align == "center" {
		for _, i := range s {
			fmt.Print(printSpaces((termWidth - len(i)) / 2))
			fmt.Print(i)
			fmt.Print(printSpaces((termWidth - len(i)) / 2))
			fmt.Println()
		}
	} else if align == "justify" {
		Justify(input, filename)
	} else {
		log.Println("choose for align left, right, center or justify")
	}
}

func printSpaces(num int) string {
	a := ""
	for i := 1; i <= num; i++ {
		a += " "
	}
	return a
}

func PrintWordJustify(s, fileName, align string) [8]string {
	myLine := ""
	myarray := [8]string{}
	for _, char := range s {
		for line := 2; line <= 9; line++ {
			myrune := int(char)
			for i := ' '; i <= '~'; i++ {
				j := (int(i) - ' ') * 9
				if myrune == int(i) {
					firstline, err := read.ReadExactLine(fileName, line+j)
					if err != nil {
						log.Print(err)
						return myarray
					}
					myLine += firstline
				}
			}
		}
		temp := strings.Split(myLine, "\n")
		for index, s := range temp[:len(temp)-1] {
			if index == 0 {
				lenText = append(lenText, len(s))
			}
			myarray[index] += s
		}
		myLine = ""
	}
	return myarray
}

func Check(str string, w [8]string) string {
	textlen := 0
	for index := range str {
		if string(str[index]) == "\\" && index != len(str)-1 {
			if string(str[index+1]) == "n" {
				continue
			}
		}
		if string(str[index]) == "n" && index != 0 {
			if string(str[index-1]) == "\\" {
				continue
			}
		}

		textlen += lenText[index]
		if textlen >= termWidth {
			str = str[:index] + "\\n" + str[index:]
			textlen = 0
		}
	}
	return str
}

func Justify(str, filename string) {
	var WordArr [][8]string
	output := strings.Split(str, " ")
	for _, word := range output {
		if word != "" {
			WordArr = append(WordArr, PrintWordJustify(word, filename, "justify"))
		}
	}
	if len(WordArr) == 1 {
		PrintWord(str, filename, "left")
		return
	}
	PrintJustify(WordArr, termWidth, filename)
}

func PrintJustify(s [][8]string, termWidth int, filename string) {
	textlen := 0
	for _, char := range s {
		textlen += len(char[0])
	}

	numSpaces := (termWidth - textlen) / (len(s) - 1)
	for i := 0; i < 8; i++ {
		for k, v := range s {
			fmt.Print(v[i])
			if k != len(s)-1 {
				fmt.Print(printSpaces(numSpaces))
			}
		}
		fmt.Println()
	}
}
