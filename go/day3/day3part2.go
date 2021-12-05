package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func get_oxygen_generator_rating(result []string, index int) string {
	var zeroes_numbers []string
	var ones_numbers []string
	zero := 0
	one := 0

	ones_numbers, zeroes_numbers = nil, nil

	for len(result) > 1 {
		for i := 0; i < len(result); i++ {
			if result[i][index] == '1' {
				one++
				ones_numbers = append(ones_numbers, result[i])
			} else {
				zero++
				zeroes_numbers = append(zeroes_numbers, result[i])
			}
		}

		if zero > one {
			result = zeroes_numbers
		} else if one > zero {
			result = ones_numbers
		} else if zero == one {
			result = ones_numbers
		}

		index++
		zero, one = 0, 0
		ones_numbers, zeroes_numbers = nil, nil
	}
	return result[0]
}

func get_co2_scrubber_rating(result []string, index int) string {
	var zeroes_numbers []string
	var ones_numbers []string
	zero := 0
	one := 0

	ones_numbers, zeroes_numbers = nil, nil
	for len(result) > 1 {
		for i := 0; i < len(result); i++ {
			fmt.Println("result", result[i])
			if result[i][index] == '1' {
				one++
				ones_numbers = append(ones_numbers, result[i])
			} else {
				zero++
				zeroes_numbers = append(zeroes_numbers, result[i])
			}
		}

		if zero > one {
			result = ones_numbers
		} else if one > zero {
			result = zeroes_numbers
		} else if zero == one {
			result = zeroes_numbers
		}

		index++
		zero, one = 0, 0
		ones_numbers, zeroes_numbers = nil, nil
	}
	return result[0]
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var zeroes_numbers []string
	var ones_numbers []string
	var most_common_result []string
	var least_common_result []string
	var zero, one, oxygen_generator_rating, co2_scrubber_rating = 0, 0, "", ""

	for scanner.Scan() {
		if scanner.Text()[0] == '1' {
			one++
			ones_numbers = append(ones_numbers, scanner.Text())

		} else if scanner.Text()[0] == '0' {
			zero++
			zeroes_numbers = append(zeroes_numbers, scanner.Text())
		}
	}

	if zero > one {
		most_common_result = zeroes_numbers
		least_common_result = ones_numbers
	} else if one > zero {
		most_common_result = ones_numbers
		least_common_result = zeroes_numbers
	} else if zero == one {
		most_common_result = ones_numbers
		least_common_result = zeroes_numbers
	}

	oxygen_generator_rating = get_oxygen_generator_rating(most_common_result, 1)
	co2_scrubber_rating = get_co2_scrubber_rating(least_common_result, 1)

	most_common_bits_in_decimal, _ := strconv.ParseInt(oxygen_generator_rating, 2, 64)
	least_common_bits_in_decimal, _ := strconv.ParseInt(co2_scrubber_rating, 2, 64)

	fmt.Println(most_common_bits_in_decimal * least_common_bits_in_decimal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
