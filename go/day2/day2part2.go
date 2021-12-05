package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	horizontal_position := 0
	depth_position := 0
	aim := 0

	for scanner.Scan() {
		directions := strings.Split(scanner.Text(), " ")
		distance, _ := strconv.Atoi(directions[1])
		if directions[0] == "forward" {
			horizontal_position = horizontal_position + distance
			depth_position += (distance * aim)

		} else if directions[0] == "up" {
			aim = aim - distance
		} else if directions[0] == "down" {
			aim = aim + distance
		}
	}

	fmt.Println(horizontal_position * depth_position)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
