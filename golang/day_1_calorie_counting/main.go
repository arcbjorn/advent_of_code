package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	filename := os.Args[2]
	packs := getPacks(filename)
	maxPack := findMaxPack(packs)

	// 66487
	fmt.Println(maxPack)
}
