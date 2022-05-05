package homework

func average(input [15]float32) (result float32) {
	var arrSum float32
	for _, n := range input {
		arrSum += n
	}
	result = arrSum / float32(len(input))
	return
}
