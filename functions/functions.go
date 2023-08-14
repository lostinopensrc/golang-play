package functions

func Swap(a, b int) (int, int) {

	var x, y int
	x, y = b, a

	return x, y
}

func Swap1(a, b *int) (int, int) {

	var x, y int
	x, y = *b, *a

	return x, y
}
