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
	leftIds, rightIds := readFile("day01/input-01.txt")

	firstPuzzle(leftIds, rightIds)
	secondPuzzle(leftIds, rightIds)
}

func secondPuzzle(leftIds []int, rightIds []int) {
	copyLeftIds := getCopyOfList(leftIds)
	copyRightIds := getCopyOfList(rightIds)

	dict := getDictWithOccurrences(copyRightIds)

	var similarityScore int
	for _, element := range copyLeftIds {
		numberOfOccurrences := dict[element]
		similarityScore += numberOfOccurrences * element
	}

	fmt.Println(similarityScore)
}

func getDictWithOccurrences(copyRightIds []int) map[int]int {
	dict := make(map[int]int)
	for _, rightId := range copyRightIds {
		dict[rightId]++
	}

	return dict
}

func firstPuzzle(leftIds []int, rightIds []int) {
	copyLeftIds := getCopyOfList(leftIds)
	copyRightIds := getCopyOfList(rightIds)

	sort.Ints(copyRightIds)
	sort.Ints(copyLeftIds)

	zippedListOfIds := zipLists(copyLeftIds, copyRightIds)

	sumOfAllDiffs := sumOfAllDiffsPerPair(zippedListOfIds)

	fmt.Println(sumOfAllDiffs)
}

func getCopyOfList(list []int) []int {
	var copyList = make([]int, len(list))
	copy(copyList, list)
	return copyList
}

func sumOfAllDiffsPerPair(zippedListOfIds [][2]int) int {
	var sumOfAllDiffs int

	for _, pair := range zippedListOfIds {
		var diff = pair[1] - pair[0]
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
