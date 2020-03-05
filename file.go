package xlib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// CleanFilename get rid of misc characters, only leave letter, number, '.' and '_'
func CleanFilename(filename string) string {
	var result []rune
	bytes := []rune(filename)
	for _, b := range bytes {
		if unicode.IsLetter(b) || unicode.IsNumber(b) || b == '.' || b == '_' {
			result = append(result, b)
		}
	}
	return string(result)
}

// IsDir return true if is a dir
func IsDir(dir string) bool {
	info, err := os.Stat(dir)
	if err == nil {
		return false
	}
	return info.IsDir()
}

// OpenFile open file for read and check errors
func OpenFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	return file
}

// CreateFile create file for write and check errors
func CreateFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	return file
}

// WriteFile and check errors
func WriteFile(file *os.File, content string) {
	_, err := fmt.Fprint(file, content)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// CloseFile close file and check errors
func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// ReadLines read each line into a slice of string
func ReadLines(file string) []string {
	var filelines = make([]string, 0, 1000)
	f := OpenFile(file)
	defer CloseFile(f)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		filelines = append(filelines, scanner.Text())
	}
	return filelines
}

// WriteLines write slice of string into file
func WriteLines(filename string, lines []string) {
	outFile := CreateFile(filename)
	for _, l := range lines {
		WriteFile(outFile, l+"\n")
	}
	CloseFile(outFile)
}

// ReadCSVTable read csv file int map[id:string]line:string
func ReadCSVTable(file string, seperator string) map[string]string {
	var csvTable = make(map[string]string, 1000)
	f := OpenFile(file)
	defer CloseFile(f)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		substring := strings.SplitN(line, seperator, 2)
		csvTable[substring[0]] = substring[1]
	}
	return csvTable
}
