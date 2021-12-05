package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input2.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	Scanner := bufio.NewScanner(file)

	var numbers []int
	for Scanner.Scan() {
		result, _ := strconv.Atoi(Scanner.Text())
		numbers = append(numbers, result)
	}

	count := 0

	for i := 0; i < len(numbers)-3; i++ {
		if numbers[i+3] > numbers[i] {
			count++
		}
	}
	fmt.Println(count)
}
