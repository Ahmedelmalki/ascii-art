package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arr := os.Args[1]
	for _, char := range arr {
		if char > 126 || char < 32 {
			fmt.Println("invalid charactere")
			return
		}
	}
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	filesplit := strings.Split(string(file), "\n")
	arrsplit := strings.Split(arr, "\\n")
	if len(arrsplit) >= 2 && arrsplit[0] == "" && arrsplit[1] == "" {
		if len(arrsplit) == 2 {
			fmt.Println()
		}
		arrsplit = arrsplit[1:]
	}
	for _, word := range arrsplit {
		if word == "" && len(arrsplit) == 1 {
			continue
		} else if word == "" {
			fmt.Println()
			continue
		}
		for j := 1; j < 9; j++ {
			for i := 0; i < len(word); i++ {
				fmt.Print(filesplit[(int(word[i])-32)*9+j])
			}
			fmt.Println()
		}
	}
}
