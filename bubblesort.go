package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5}
	fmt.Println("After 0 sweep(s):", numbers)
	//sweep
	sweep(numbers)
	fmt.Println("After 1 sweep(s):", numbers)

}

func sweep(numbers []int) {
	var N = len(numbers)
	fmt.Println(N)
	var first = 0
	var second = 1
	for second < N {

		var firstNumber = numbers[first]
		var secondNumber = numbers[second]
		if firstNumber > secondNumber {
			numbers[first] = secondNumber
			numbers[second] = first
		}
		first++
		second++
	}
}
