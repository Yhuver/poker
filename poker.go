package main

import (
	"fmt"
	"sort"
)

func isStraight(cards []int) bool {
	sort.Ints(cards)

	sum := 1
	if cards[0] == 2 && cards[len(cards)-1] == 14 {
		sum++
	}
	for i := 1; i < len(cards); i++ {
		if cards[i]-1 == cards[i-1] {
			sum++
		} else {
			sum = 1
		}
		if sum > 4 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("escalera:", isStraight([]int{14, 2, 3, 4, 5}))
	fmt.Println("escalera:", isStraight([]int{9, 10, 11, 12, 13}))
	fmt.Println("escalera:", isStraight([]int{2, 7, 8, 5, 10, 9, 11}))
	fmt.Println("escalera:", isStraight([]int{7, 8, 12, 13, 14}))
}
