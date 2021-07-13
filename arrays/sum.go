package arrays

func Sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		var sum int
		if len(numbers) == 0 {
			sum = 0
		} else {
			sum = Sum(numbers[1:])
		}
		sums = append(sums, sum)
	}

	return sums
}
