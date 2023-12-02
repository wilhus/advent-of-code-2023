package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"wilhus/advent"
)

func main() {
	input := advent.ReadFile("input.txt")
	bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sum1 := 0
	sum2 := 0
	for gameIdx, game := range strings.Split(input, "\n") {
		poss := true
		pow := 1
		for col, lim := range bag {
			maxNum := 0
			matches := regexp.MustCompile(`(\d+)\s`+col).FindAllStringSubmatch(game, -1)
			for _, match := range matches {
				intNum, _ := strconv.Atoi(match[1])
				maxNum = max(maxNum, intNum)
				if intNum > lim {
					poss = false
				}
			}
			pow *= maxNum
		}
		if poss {
			sum1 += gameIdx + 1
		}
		sum2 += pow
	}
	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}
