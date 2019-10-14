package slices

// Sum sums a list of integers
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}

// SumAll Takes a varying list of slices of integers and sum them
func SumAll(listOfNumbers ...[]int) []int {
	var sums []int
	for _, numbers := range listOfNumbers {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails sums the varying list of slices except their first element
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}

	return sums
}
