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

func IsDotFirst(file string) bool {
	if file[0] == '.' {
		return true
	} else {
		return false
	}
}

func DigitCount(num int) int {
	digit_count := 1
	for num > 0 {
		num /= 10
		digit_count += 1
	}
	return digit_count
}

func WriteFileInfo() {
	files, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	//checking the longest file in directory
	longest := 0
	longest_size := 0
	for _, file := range files {
		if IsExe(file.Name()) || IsDotFirst((file.Name())) {
			continue
		}

		if len(file.Name()) > longest {
			longest = len(file.Name())
		}

		info, _ := file.Info()
		if DigitCount(int(info.Size())) > longest_size {
			longest_size = DigitCount(int(info.Size()))
		}
	}

	fmt.Println("========================")
	for _, file := range files {
		if IsExe(file.Name()) || IsDotFirst((file.Name())) {
			continue
		}

		//name
		var builder strings.Builder
		local_len := len(file.Name())
		len_diff := longest - local_len

		builder.WriteString(file.Name())
		for len_diff > 0 {
			builder.WriteRune(' ')
			len_diff -= 1
		}
		builder.WriteString("          ")

		//size
		info, _ := file.Info()
		builder.WriteString(IntToString(int((info.Size()))))

		local_size_len := DigitCount(int(info.Size()))
		size_len_diff := longest_size - local_size_len
		for size_len_diff > 0 {
			builder.WriteRune(' ')
			size_len_diff -= 1
		}

		builder.WriteString(" kb")
		fmt.Println(builder.String())
	}
	fmt.Println("========================")
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "fileinf" {
		WriteFileInfo()
	}
}
