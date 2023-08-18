package arrays

import "fmt"

func Array1() {
	myarray := [3]string{"hello", "hi", "bye"}

	for i := 0; i < len(myarray); i++ {
		fmt.Println(myarray[i])
	}
}

func Array2() {
	myarray := [...]int{1, 2, 3}

	for x := 0; x < len(myarray); x++ {
		fmt.Println(myarray[x])
	}
}
