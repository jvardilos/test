package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/apex/log"
	"github.com/marty-farce/test/util"
)

func openFile() string {
	readFile, err := os.Open("sample-json.json")
	if err != nil {
		fmt.Println(err)
		log.Warn("error opening try absolute path")
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string
	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
	readFile.Close()

	var jsonOneLine string
	for _, val := range fileTextLines {
		jsonOneLine += strings.TrimSpace(val)
	}

	return jsonOneLine
}

func main() {
	x := openFile()
	util.Run(x)
}
