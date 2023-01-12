package conditionales

import (
	"fmt"
)

func main() {

	condicionales()
}

func condicionales() {
	{
		// Condicionales
		yearsOld := 20
		if yearsOld > 18 {
			fmt.Println("Es mayor")
		}

		boolVal := true
		if boolVal {
			fmt.Println("Es verdadero")
		} else {
			fmt.Println("Es falso")
		}

		// if value := fn(); value == true {

		// }

		if value := true; value {
			fmt.Println("Es verdadero")
		}

		number := 3
		if number == 1 {
			fmt.Println("Es 1")
		} else if number == 2 {
			fmt.Println("Es 2")
		} else {
			fmt.Println("Es 3")
		}

		switch number {
		case 1:
			fmt.Println("Es 1")
		case 2:
			fmt.Println("Es 2")
		case 3:
			fmt.Println("Es 3")
		default:
			fmt.Println("No definido")
		}

		switch number2 := 2; number2 {
		case 1:
			fmt.Println("Es 1")
		case 2:
			fmt.Println("Es 2")
		case 3:
			fmt.Println("Es 3")
		default:
			fmt.Println("No definido")
		}
	}
}
