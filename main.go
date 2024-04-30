package main

import (
	"bufio"
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

func GetSizeAndMetric(size int) (float32, string) {
	if size < 1024 {
		return float32(size), "kb"
	} else if size < 1024*1024 {
		return float32(size) / 1024.0, "mb"
	} else if size < 1024*1024*1024 {
		return float32(size) / (1024.0 * 1024.0), "gb"
	}
	return float32(size), "kb"
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
	fmt.Println("=========================================")
	fmt.Println("FileName             loc          Size")
	fmt.Println("=========================================")
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
		builder.WriteString("        ")

		//lines of code
		readFile, err := os.Open(file.Name())

		if err != nil {
			fmt.Println(err)
		}
		fileScanner := bufio.NewScanner(readFile)

		fileScanner.Split(bufio.ScanLines)

		number_of_lines := 0
		for fileScanner.Scan() {
			if len(fileScanner.Text()) != 0 {
				number_of_lines += 1
			}
		}
		builder.WriteString(IntToString(number_of_lines))
		builder.WriteString("       ")
		readFile.Close()

		//size
		info, _ := file.Info()
		file_size := int((info.Size()))

		new_size, metric := GetSizeAndMetric(file_size)
		new_size_formated := fmt.Sprintf("%.1f", new_size)
		builder.WriteString(new_size_formated)

		local_size_len := DigitCount(int(info.Size()))
		size_len_diff := longest_size - local_size_len
		for size_len_diff > 0 {
			builder.WriteRune(' ')
			size_len_diff -= 1
		}

		builder.WriteString(metric)

		fmt.Println(builder.String())
	}
	fmt.Println("=========================================")
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "fileinf" {
		WriteFileInfo()
	}
}
