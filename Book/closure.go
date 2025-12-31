package Book

func Counter(n int) func() int {
	return func() int {
		n++
		return n
	}
}
