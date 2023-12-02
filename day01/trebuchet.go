package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"wilhus/advent"
)

func getIntNum(line string) int {
	nums := regexp.MustCompile("[0-9]").FindAllString(line, -1)
	num := nums[0] + nums[len(nums)-1]
	intNum, _ := strconv.Atoi(num)
	return intNum
}

func part1(input string) {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		sum += getIntNum(line)
	}
	fmt.Println("Part 1:", sum)
}

func part2(input string) {
	digits := map[string]string{
		"one": "1", "two": "2", "three": "3",
		"four": "4", "five": "5", "six": "6",
		"seven": "7", "eight": "8", "nine": "9",
	}
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		for key, val := range digits {
			keyLen := len(key)
			line = strings.ReplaceAll(line, key, key[:keyLen/2]+val+key[keyLen/2:])
		}
		sum += getIntNum(line)
	}
	fmt.Println("Part 2:", sum)
}

func main() {
	input := advent.ReadFile("input.txt")
	part1(input)
	part2(input)
}
