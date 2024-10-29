package array

func Sum(numbers []int) int {
	sum := 0
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }
	// return sum
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	// sum := []int{0, 0}
	// for _, number := range numbers1 {
	// 	sum[0] += number
	// }
	// for _, number := range numbers2 {
	// 	sum[1] += number
	// }
	// return sum

	// totalSum := []int{}
	// for _, num := range numbers {
	// 	sum := 0
	// 	for _, n := range num {
	// 		sum += n
	// 	}
	// 	totalSum = append(totalSum, sum)
	// }

	// return totalSum

	lenSlices := len(numbers)
	sums := make([]int, lenSlices)

	for i, slice := range numbers {
		sums[i] = Sum(slice)
	}
	return sums
}
