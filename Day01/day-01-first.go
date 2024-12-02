package main

import (
	helpermethods "advent-of-code-2024"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	leftIds, rightIds := readFile("Day01/input-01.txt")

	sort.Ints(leftIds)
	sort.Ints(rightIds)

	zippedListOfIds := zipLists(leftIds, rightIds)

	sumOfAllDiffs := sumOfAllDiffsPerPair(zippedListOfIds)

	fmt.Println(sumOfAllDiffs)
}

func sumOfAllDiffsPerPair(zippedListOfIds [][2]int) int {
	var sumOfAllDiffs int

	for n := range zippedListOfIds {
		var diff = zippedListOfIds[n][1] - zippedListOfIds[n][0]
		sumOfAllDiffs += helpermethods.Abs(diff)
	}
	
	return sumOfAllDiffs
}

func readFile(filePath string) ([]int, []int) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	var leftList []int
	var rightList []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		leftNumber, rightNumber := parseLine(line)
		leftList = append(leftList, leftNumber)
		rightList = append(rightList, rightNumber)
	}

	return leftList, rightList
}

func parseLine(line string) (int, int) {
	parts := strings.Fields(line)
	num1, _ := strconv.Atoi(parts[0])
	num2, _ := strconv.Atoi(parts[1])
	return num1, num2
}

func zipLists(leftIds []int, rightIds []int) [][2]int {
	zippedListOfIds := make([][2]int, len(leftIds))
	for i := 0; i < len(leftIds); i++ {
		zippedListOfIds[i] = [2]int{leftIds[i], rightIds[i]}
	}
	return zippedListOfIds
}
