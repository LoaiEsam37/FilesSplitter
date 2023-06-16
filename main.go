package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inputFile := flag.String("i", "", "input file")
	lines := flag.Int("l", 0, "number of lines")
	flag.Parse()
	if *inputFile == "" || *lines == 0 {
		fmt.Println("Error: both -i and -l flags are required")
		os.Exit(1)
	}
	file, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println("Error: could not open file")
		os.Exit(1)
	}
	defer file.Close()
	var parentSlice [][]string
	var linesSlice []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linesSlice = append(linesSlice, scanner.Text())
		if len(linesSlice) == *lines {
			parentSlice = append(parentSlice, linesSlice)
			linesSlice = []string{}
		}
	}
	if len(linesSlice) > 0 {
		parentSlice = append(parentSlice, linesSlice)
	}
	err = os.Mkdir("temp", 0755)
	if err != nil {
		fmt.Println("Error: could not create temp directory")
		os.Exit(1)
	}
	for i, sublines := range parentSlice {
		fileName := fmt.Sprintf("File%02d.txt", i+1)
		filePath := filepath.Join("temp", fileName)
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error: could not create file")
			os.Exit(1)
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		_, err = writer.WriteString(strings.Join(sublines, "\n"))
		if err != nil {
			fmt.Println("Error: could not write to file")
			os.Exit(1)
		}
		writer.Flush()
	}
}
