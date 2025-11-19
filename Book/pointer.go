package Book

func add10(p *int) int {
	*p += 10
	return *p
}
