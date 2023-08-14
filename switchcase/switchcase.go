package switchcase

import "fmt"

func SwitchCase() {
	switch value := 3; value {
	case 1:
		fmt.Println("in case 1")
	case 2:
		fmt.Println("in case 2")
	case 3:
		fmt.Println("in case 3")
	}
}
