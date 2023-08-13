package conditions

import "fmt"

func Conditions() {

	var v int = 100

	if v == 100 {
		fmt.Println("Executed if block")
	} else {
		fmt.Println("Executed else block")
	}
}
