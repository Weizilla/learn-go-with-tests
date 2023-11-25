package arrays

func Sum(numbers []int) int {
	s := 0
	for _, i := range numbers {
		s += i
	}
	return s
}

func SumAll(allNumbers ...[]int) []int {
	var sums []int
	for _, numbers := range allNumbers {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(allNumbers ...[]int) []int {
	var sums []int
	for _, numbers := range allNumbers {
		if len(numbers) > 1 {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}
