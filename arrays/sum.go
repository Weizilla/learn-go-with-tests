package arrays

func Sum(values []int) int {
	var s int
	for _, number := range values {
		s += number
	}
	return s
}

func SumAll(nums ...[]int) []int {
	var sums []int

	for _, ns := range nums {
		sums = append(sums, Sum(ns))
	}
	return sums
}

func SumAllTails(nums ...[]int) []int {
	var sums []int

	for _, ns := range nums {
		if len(ns) > 0 {
			nss := ns[1:]
			sums = append(sums, Sum(nss))
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}
