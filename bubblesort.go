package main

import "fmt"

func main() {
	var numbers = []int{20, 3, 34, 5, 6}
	fmt.Println("Unsorted:", numbers)

	bubbleSort(numbers)
	fmt.Println("Sorted:", numbers)
}

func bubbleSort(numbers []int) {
	var N = len(numbers)
	var i int
	for i = 0; i < N; i++ {
		if !sweep(numbers, i) {
			return
		}
	}
}

func sweep(numbers []int, prevPasses int) bool {
	var N = len(numbers)
	var firstIndex = 0
	var secondIndex = 1
	var didSwap = false

	for secondIndex < (N - prevPasses) {
		var firstNumber = numbers[firstIndex]
		var secondNumber = numbers[secondIndex]

		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			didSwap = true
		}

		firstIndex++
		secondIndex++
	}
	return didSwap
}
