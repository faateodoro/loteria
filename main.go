package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var numbers []int
	var i = 0
	for i < 15 {
		var random = rand.Intn(25) + 1
		if implContains(numbers, random) {
			continue
		}
		numbers = append(numbers, random)
		i++
	}
	sort.Ints(numbers)
	fmt.Println(numbers)
}

func implContains(sortedNumbers []int, random int) bool{
	for _, value := range sortedNumbers {
		if value == random {
			return true
		}
	}
	return false
}