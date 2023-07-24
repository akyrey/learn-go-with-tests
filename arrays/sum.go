package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	result := make([]int, 0)

	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}

	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	result := make([]int, 0)

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(numbers[1:]))
		}
	}

	return result
}
