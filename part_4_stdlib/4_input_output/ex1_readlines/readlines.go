package main

import (
	"os"
	"strings"
	"bufio"
)


func readLines(name string) ([]string, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return []string{}, err
	}
	var buffer strings.Builder
	result := []string{}
	for _, c := range data {
		if c == '\n' {
			result = append(result, buffer.String())
			buffer.Reset()
		} else {
			buffer.WriteByte(c)
		}
	}
	return result, nil
}

func readLines2(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return []string{}, err	
	}
	defer file.Close()
	scanner := bufio.NewScanner(file) 
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result, nil
}

func main() {
	readLines("test.txt")
	readLines2("test2.txt")
}