package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arr := os.Args[1]
	file, _ := os.ReadFile("standard.txt")
	filesplit := strings.Split(string(file), "\n")
	arrsplit := strings.Split(arr, "\\n")
	for _, word := range arrsplit {
		for j := 1; j < 9; j++ {
			for i := 0; i < len(word); i++ {
				fmt.Print(filesplit[(int(word[i])-32)*9+j])
			}
			fmt.Println()
		}
	}
}
