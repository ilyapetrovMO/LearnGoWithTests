package main

func Sum(numbersToSum []int) (sum int) {
	for _, el := range numbersToSum {
		sum += el
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, 0, len(numbersToSum))
	for _, el := range numbersToSum {
		sums = append(sums, Sum(el))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {

	sums := make([]int, 0, len(numbersToSum))
	for _, el := range numbersToSum {
		if len(el) == 0 {
			sums = append(sums, 0)
		} else {
			tail := el[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
