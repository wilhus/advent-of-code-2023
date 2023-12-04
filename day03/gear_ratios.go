package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"wilhus/advent"
)

func isSymbol(s string) bool {
	return regexp.MustCompile(`[^\d.]`).MatchString(s)
}

func isNumber(s string) bool {
	return regexp.MustCompile(`\d`).MatchString(s)
}

func fullNumber(matrix [][]string, num string, row, col int) (int, bool) {
	for leftCol := col - 1; leftCol >= 0; leftCol-- {
		leftItem := matrix[row][leftCol]
		if isNumber(leftItem) {
			num = leftItem + num
		} else {
			break
		}
	}
	shift := false
	for rightCol := col + 1; rightCol < len(matrix[row]); rightCol++ {
		rightItem := matrix[row][rightCol]
		if isNumber(rightItem) {
			num = num + rightItem
			shift = true
		} else {
			break
		}
	}
	intNum, _ := strconv.Atoi(num)
	return intNum, shift
}

func findAdjacentItems(matrix [][]string, row, col int) []int {
	dirs := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	adjacentItems := []int{}
	for dir := 0; dir < len(dirs); dir++ {
		rowOffset := dirs[dir][0]
		colOffset := dirs[dir][1]
		newRow, newCol := row+rowOffset, col+colOffset
		if newRow >= 0 && newRow < len(matrix) && newCol >= 0 && newCol < len(matrix[row]) {
			item := matrix[newRow][newCol]
			if isNumber(item) {
				num, shift := fullNumber(matrix, item, newRow, newCol)
				adjacentItems = append(adjacentItems, num)
				if colOffset <= 0 && shift {
					dir += 2
				}
			}
		}
	}
	return adjacentItems
}

func main() {
	input := advent.ReadFile("input.txt")
	matrix := [][]string{}
	for _, line := range strings.Split(input, "\n") {
		row := []string{}
		row = append(row, strings.Split(line, "")...)
		matrix = append(matrix, row)
	}
	sum1 := 0
	sum2 := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			item := matrix[row][col]
			if !isSymbol(item) {
				continue
			}
			adjascentItems := findAdjacentItems(matrix, row, col)
			ratio := 1
			for _, adjascentItem := range adjascentItems {
				sum1 += adjascentItem
				if item == "*" && len(adjascentItems) == 2 {
					ratio *= adjascentItem
				}
			}
			if ratio > 1 {
				sum2 += ratio
			}
		}
	}
	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}
