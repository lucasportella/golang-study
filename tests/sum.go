package tests

func Sum(i ...int) int {
	result := 0
	for _, v := range i {
		result += v
	}
	return result
}
