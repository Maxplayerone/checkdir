package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func IsExe(file string) bool {
	name_and_ext := strings.Split(file, ".")
	if name_and_ext[1] == "exe" {
		return true
	} else {
		return false
	}
}

func WriteFileInfo() {
	files, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	//checking the longest file in directory
	longest := 0
	for _, file := range files {
		if IsExe(file.Name()) {
			continue
		}

		if len(file.Name()) > longest {
			longest = len(file.Name())
		}
	}

	fmt.Println("========================")
	for _, file := range files {
		if IsExe(file.Name()) {
			continue
		}

		var builder strings.Builder
		local_len := len(file.Name())
		len_diff := longest - local_len

		builder.WriteString(file.Name())
		for len_diff > 0 {
			builder.WriteRune(' ')
			len_diff -= 1
		}
		builder.WriteString("          ")
		builder.WriteString("hi")
		fmt.Println(builder.String())
	}
	fmt.Println("========================")
}

func main() {
	args := os.Args[1:]
	if args[0] == "fileinf" {
		WriteFileInfo()
	}
}
