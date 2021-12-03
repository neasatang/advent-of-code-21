package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	horizontal_position := 0
	depth_position := 0

	for scanner.Scan() {
		directions := strings.Split(scanner.Text(), " ")
		distance, _ := strconv.Atoi(directions[1])
		if directions[0] == "forward" {
			horizontal_position = horizontal_position + distance
		} else if directions[0] == "up" {
			depth_position = depth_position - distance
		} else if directions[0] == "down" {
			depth_position = depth_position + distance
		}
	}

	fmt.Println(horizontal_position * depth_position)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
