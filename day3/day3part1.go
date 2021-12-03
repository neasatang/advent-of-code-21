package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	ones := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	zeroes := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	most_common_bits := ""
	least_common_bits := ""

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for i := 0; i < len(ones); i++ {
			if scanner.Text()[i] == '1' {
				zeroes[i] = zeroes[i] + 1
			} else if scanner.Text()[i] == '0' {
				ones[i] = ones[i] + 1
			}
		}
	}

	for i := 0; i < len(ones); i++ {
		if ones[i] > zeroes[i] {
			most_common_bits += "0"
		} else {
			most_common_bits += "1"
		}
	}

	for i := 0; i < len(ones); i++ {
		convert_to_number, _ := strconv.Atoi(string(most_common_bits[i]))
		invert := convert_to_number ^ 1
		least_common_bits += strconv.Itoa(invert)
	}

	most_common_bits_in_decimal, _ := strconv.ParseInt(most_common_bits, 2, 64)
	least_common_bits_in_decimal, _ := strconv.ParseInt(least_common_bits, 2, 64)

	fmt.Println(most_common_bits_in_decimal * least_common_bits_in_decimal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
