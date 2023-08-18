package slices

import "fmt"

func Slice_Me() {
	var s []int
	s = append(s, 4, 5, 6)
	fmt.Println("The slice populated is:", s)
	for key, value := range s {
		fmt.Printf("Index = %d has value = %d populated\n", key, value)
	}
}
