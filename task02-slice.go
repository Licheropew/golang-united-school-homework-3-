package homework

func reverse(input []int64) (result []int64) {
	result = make([]int64, len(input))
	copy(result, input)
	i := 0
	j := len(result) - 1
	for i < j {
		result[i], result[j] = result[j], result[i]
		i += 1
		j -= 1
	}
	return
}
