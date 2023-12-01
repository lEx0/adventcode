package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	//goland:noinspection GoUnhandledErrorResult
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		sum += GetContainNumberFromString(scanner.Text())
	}

	println(sum)
}

func GetContainNumberFromString(input string) int {
	numbers := ""

	for i := 0; len(input) >= 1; i++ {
		switch {
		case input[:1] == "1", len(input) >= 3 && input[:3] == "one":
			numbers += "1"
		case input[:1] == "2", len(input) >= 3 && input[:3] == "two":
			numbers += "2"
		case input[:1] == "3", len(input) >= 5 && input[:5] == "three":
			numbers += "3"
		case input[:1] == "4", len(input) >= 4 && input[:4] == "four":
			numbers += "4"
		case input[:1] == "5", len(input) >= 4 && input[:4] == "five":
			numbers += "5"
		case input[:1] == "6", len(input) >= 3 && input[:3] == "six":
			numbers += "6"
		case input[:1] == "7", len(input) >= 5 && input[:5] == "seven":
			numbers += "7"
		case input[:1] == "8", len(input) >= 5 && input[:5] == "eight":
			numbers += "8"
		case input[:1] == "9", len(input) >= 4 && input[:4] == "nine":
			numbers += "9"
		}

		input = input[1:]
	}

	if digit, err := strconv.Atoi(numbers[0:1] + numbers[len(numbers)-1:]); err != nil {
		panic(err)
	} else {
		return digit
	}
}
