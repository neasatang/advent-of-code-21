package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	prev := 0
	count := 0

	for scanner.Scan() {
		intVar, _ := strconv.Atoi(scanner.Text())
		if prev < intVar && prev != 0 {
			count++
		}
		prev = intVar
	}

	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
