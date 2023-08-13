package loops

import "fmt"

func Loops() {
	var s string = "hello"
	for i, j := range s {
		fmt.Printf("Index for %c is %d\n", j, i)
	}

}
