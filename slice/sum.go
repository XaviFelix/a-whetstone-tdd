package main

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// func SumAll(arraysOfNumbers ...[]int) []int {

// 	numberOfArrays := len(arraysOfNumbers)
// 	sums := make([]int, numberOfArrays)

// 	for i, currentArray := range arraysOfNumbers {
// 		sums[i] = Sum(currentArray)
// 	}

// 	return sums
// }

func SumAll(arraysOfNumbers ...[]int) []int {
	var sums []int
	for _, current := range arraysOfNumbers {
		sums = append(sums, Sum(current))
	}
	return sums
}

func SumAllTails(arraysOfnumbers ...[]int) []int {
	var sum []int
	for _, current := range arraysOfnumbers {
		if len(current) == 0 {
			sum = append(sum, 0)
		} else {
			tail := current[1:]
			sum = append(sum, Sum(tail))
		}
	}
	return sum
}
