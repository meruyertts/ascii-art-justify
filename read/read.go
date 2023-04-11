package read

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ReadExactLine(fileName string, lineNumber int) (string, error) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer inputFile.Close()
	br := bufio.NewReader(inputFile)
	for i := 1; i < lineNumber; i++ {
		_, _ = br.ReadString('\n')
	}
	str, err := br.ReadString('\n')
	if err != nil {
		return "", err
	}

	return str, nil
}

func GetTermWidth() int {
	out, er1 := exec.Command("tput", "cols").Output()
	out1 := strings.TrimSuffix(string(out), "\n")
	num, er2 := strconv.Atoi(string(out1))
	if er1 != nil {
		log.Fatal(er1)
		os.Exit(1)
	}
	if er2 != nil {
		log.Fatal(er2)
		os.Exit(1)
	}
	return num
}
