package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func mean(numbers []int) float64 {
	var sum int

	for _, elem := range numbers {
		sum += elem
	}
	return float64(sum / len(numbers))
}

func median(numbers []int) float64 {
	sort.Ints(numbers)
	if len(numbers) % 2 != 0 {
		middle := float64(len(numbers) / 2) + 0.5
		return float64(numbers[int(middle)])
	} else {
		middle := len(numbers) / 2
		result := float64(numbers[middle - 1] + numbers[middle]) / 2
		return result
	}
}

func mode(numbers []int) int {
	var elements map[int]int
	var max, result int = 0, 0

	elements = make(map[int]int)

	for _, elem := range numbers {
		if _, ok := elements[elem]; ok {
			elements[elem]++
		} else {
			elements[elem] = 0
		}
	}
	for key, value := range elements {
		if value > max {
			max = value
			result = key
		}
	}
	return result
}

func main() {
	var numbers []int

	for {
		fmt.Print("Type a number or 'exit': ")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if strings.TrimSpace(text) == "exit" {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to read from stdin: %s\n", err)
			continue
		}
		tmp, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to convert to an int: %s\n", err)
			continue
		}
		if tmp > -100000 && tmp < 100000 {
			numbers = append(numbers, tmp)
		} else {
			fmt.Fprintf(os.Stderr, "Number is out of range (-100000:100000)!\n")
		}
	}
	if len(numbers) != 0 {
		fmt.Println("```")
		fmt.Printf("Mean: %.2f\n", mean(numbers))
		fmt.Printf("Median: %.2f\n", median(numbers))
		fmt.Printf("Mode: %d\n", mode(numbers))
		fmt.Println("```")
	}
}
