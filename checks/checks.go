package checks

import (
	"bufio"
	"crypto/sha256"
	"errors"
	"io"
	"log"
	"os"
	"regexp"
	"unicode"
)

func FileNameCheck(fName string) string {
	myFile := ""
	switch fName {
	case "standard":
		myFile = "standard.txt"
	case "shadow":
		myFile = "shadow.txt"
	case "thinkertoy":
		myFile = "thinkertoy.txt"
	default:
		myFile = ""
	}

	return myFile
}

func LineCounter(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	if lineCount < 855 {
		return errors.New("file does not contain all characters")
	}
	return nil
}

func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func TxtFileCheck(fileName string) bool {
	hashStandard := []byte{225, 148, 241, 3, 52, 66, 97, 122, 184, 167, 142, 28, 166, 58, 32, 97, 245, 204, 7, 163, 240, 90, 194, 38, 237, 50, 235, 157, 253, 34, 166, 191}
	hashShadow := []byte{184, 17, 37, 168, 183, 46, 207, 226, 35, 69, 169, 190, 218, 184, 99, 86, 141, 179, 152, 16, 96, 21, 242, 206, 76, 172, 130, 232, 162, 21, 7, 76}
	hashThinkertoy := []byte{236, 241, 252, 123, 255, 114, 166, 211, 68, 247, 17, 86, 18, 3, 196, 224, 126, 132, 206, 58, 147, 120, 23, 16, 71, 60, 235, 235, 128, 88, 253, 28}
	file, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return false
	}
	defer file.Close()
	buf := make([]byte, 30*1024)
	sha256 := sha256.New()
	for {
		n, err := file.Read(buf)
		if n > 0 {
			_, err := sha256.Write(buf[:n])
			if err != nil {
				log.Fatal(err)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Read %d bytes: %v", n, err)
			break
		}
	}
	sum := sha256.Sum(nil)
	switch fileName {
	case "shadow.txt":
		if string(sum) == string(hashShadow) {
			return true
		}
	case "standard.txt":
		if string(sum) == string(hashStandard) {
			return true
		}
	case "thinkertoy.txt":
		if string(sum) == string(hashThinkertoy) {
			return true
		}
	}
	return false
}

func MyAlign(myAlignStr string) (string, error) {
	var err error
	re := regexp.MustCompile(`--align=`)
	match := re.Split(myAlignStr, -1)
	if len(match) == 1 {
		return "", errors.New("choose left, right, center or justify")
	}
	return match[1], err
}
