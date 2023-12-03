package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"unicode"
)

func read_file(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}

func find_digit(line string) (int, error) {
	var (
		result int = 0
	)

	for _, char := range line {
		if unicode.IsNumber(char) {
			result = int(char - '0')
			break
		}
	}

	for i := range line {
		var char rune = ([]rune(line))[len(line)-i-1]
		if unicode.IsNumber(char) {
			result = result*10 + int(char-'0')
			return result, nil
		}
	}

	return -1, errors.New("No digit found")
}

func main() {
	var filePath string
	if len(os.Args) == 2 {
		filePath = os.Args[1]
	} else if len(os.Args) == 1 {
		filePath = "inputs.txt"
	} else {
		panic("Unknown amount of input parameters.")
	}

	input, err := read_file(filePath)
	if err != nil {
		panic(err)
	}

	var result int = 0
	for i, line := range input {
		val, err := find_digit(line)

		if err != nil {
			fmt.Println("Error on line ", i, ":")
			panic(err)
		}

		result += val
	}

	fmt.Println("Result: ", result)
}
