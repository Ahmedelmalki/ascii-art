package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	arg := os.Args[1]
	r := []rune(arg)
	for i := 0; i < len(r); i++ {
		temp(r[i], file)

	}
	defer file.Close()
}

func temp(char rune, file *os.File) /* []string */ {

	scanner := bufio.NewScanner(file)
	position := 0
	file.Seek(0, 0)
	for scanner.Scan() {
		position++

		startLine := int(char-32)*9 + 2
		endLine := startLine + 7
		if position < startLine {

			continue
		}
		if position > endLine {
			break
		}
	
		text := scanner.Text()

		fmt.Println(text)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	// return
}
