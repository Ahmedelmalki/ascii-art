package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("tiaaaaaaaaakh")
	} else {
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
		isprinted := false
		for index, word := range arrsplit {
			if word != "" {
				for j := 1; j < 9; j++ {
					for i := 0; i < len(word); i++ {
						fmt.Print(filesplit[(int(word[i])-32)*9+j])
						isprinted = true
					}
					fmt.Println("y")
				}
			} else if word == "" {
				if index == len(arrsplit)-1 && !isprinted {
					continue
				} else {
					fmt.Println("x")
					continue
				}
			}
		}
	}

}
