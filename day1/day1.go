package day1

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part1() {

	file, err := os.ReadFile("./inputs/day1/part1_input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileString := strings.Split(string(file), "\n")

	var c1 []int
	var c2 []int

	for _, v := range fileString {
		v1, err := strconv.Atoi(v[:5])
		if err != nil {
			return
		}

		v2, err := strconv.Atoi(v[len(v)-5:])
		if err != nil {
			return
		}

		c1 = append(c1, v1)
		c2 = append(c2, v2)
	}
	sort.Ints(c1)
	sort.Ints(c2)

	var total float64

	for i, _ := range c1 {
		result := c2[i] - c1[i]
		total += math.Abs(float64(result))
	}

	fmt.Println(int(total))
}

func Part2() {

	file, err := os.ReadFile("./inputs/day1/part1_input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileString := strings.Split(string(file), "\n")

	var c1 []int
	var c2 []int

	for _, v := range fileString {
		v1, err := strconv.Atoi(v[:5])
		if err != nil {
			return
		}

		v2, err := strconv.Atoi(v[len(v)-5:])
		if err != nil {
			return
		}

		c1 = append(c1, v1)
		c2 = append(c2, v2)
	}
	sort.Ints(c1)
	sort.Ints(c2)

	var similarity int

	for _, v1 := range c1 {
		var occurance int
		for _, v2 := range c2 {
			if v1 == v2 {
				occurance++
			}
		}
		similarity += v1 * occurance
	}

	fmt.Println(similarity)
}
