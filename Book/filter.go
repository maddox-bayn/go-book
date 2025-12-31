package Book

func filter(n int, fn func(int) bool) []int {
	var result []int
	for i := 0; i < n; i++ {
		if fn(i) {
			result = append(result, i)
		}
	}
	return result
}
