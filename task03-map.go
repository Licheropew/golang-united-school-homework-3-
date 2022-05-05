package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	var indexArr []int
	for k := range input {
		indexArr = append(indexArr, k)
	}
	sort.Ints(indexArr)
	for i := range indexArr {
		result = append(result, input[i])
	}
	return
}
