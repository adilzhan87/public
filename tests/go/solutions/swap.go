package solutions

func Swap(a, b *int) {
	*a, *b = *b, *a
}
