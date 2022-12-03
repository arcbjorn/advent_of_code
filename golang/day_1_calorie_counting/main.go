package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getPacks(filename string) []int {
	var packs []int

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return packs
	}

	defer file.Close()

	var line string
	var amount int

	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		if line == "" {
			packs = append(packs, sum)
			sum = 0

		} else {
			amount, err = strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			sum += amount
		}
	}

	return packs
}

func findMaxPack(packs []int) int {
	var maxAmount int

	for _, amount := range packs {
		if amount > maxAmount {
			maxAmount = amount
		}
	}

	return maxAmount
}

func getMaxPacks(packs []int) (maxPack int, last3Sum int) {
	sort.Ints(packs[:])
	last3 := packs[len(packs)-3:]

	// 66487
	maxPack = packs[len(packs)-1]

	for _, amount := range last3 {
		// 197301
		last3Sum += amount
	}

	return maxPack, last3Sum
}

func main() {
	filename := os.Args[2]
	packs := getPacks(filename)
	maxPack, last3Sum := getMaxPacks(packs)

	fmt.Println(maxPack, last3Sum)
}
