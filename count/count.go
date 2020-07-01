package count

func Count(s string, r rune) (res int) {
	for _, letter := range s {
		if letter == r {
			res++
		}
	}
	return
}
