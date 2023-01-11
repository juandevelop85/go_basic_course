package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {

	fors()
}

func variables() {
	var myIntVar int = -12
	fmt.Println("Mi variable es:", myIntVar)

	var myInt8Var int8 = 123
	fmt.Println("Mi variable es:", myInt8Var)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myInt8Var, unsafe.Sizeof(myInt8Var), unsafe.Sizeof(myInt8Var)*8)

	var myFloat32Var float32 = 123.1
	fmt.Println("Mi variable es:", myFloat32Var)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	var myFloat64Var float64 = 123.2
	fmt.Println("Mi variable es:", myFloat64Var)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)

	var myInt8VarNotDefined int8
	fmt.Println("Mi variable es:", myInt8VarNotDefined)

	var myUintVar uint8 = 255
	fmt.Println("Mi variable es:", myUintVar)

	var myStringVar string = "Juan Sebastian"
	fmt.Printf("Mi variable es: %s\n", myStringVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myStringVar, unsafe.Sizeof(myStringVar), unsafe.Sizeof(myStringVar)*8)

	var myStringVarNotDefined string
	fmt.Println("Mi variable es:", myStringVarNotDefined)

	var myBoolVar bool = true
	fmt.Println("Mi variable es:", myBoolVar)
	fmt.Println("Mi variable esta en:", &myBoolVar)

	var myBoolVarNotDefined bool
	fmt.Println("Mi variable es:", myBoolVarNotDefined)

	myBoolVar2 := false
	fmt.Println("Mi variable es:", myBoolVar2)

	const myIntConst = -12
	fmt.Println("Mi constante es:", myIntConst)

	const myStringConst = `Vernaza
	Lopez`
	fmt.Println("Mi constante es:", myStringConst)

	// Scope 1
	{
		var myFloat32Var float32 = 123.1233232
		fmt.Printf("Mi variable es: %.2f\n", myFloat32Var)
		fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

		intVal2, err := strconv.ParseInt("23232s", 0, 64)
		fmt.Println(err)
		fmt.Printf("Mi variable es: %d\n", intVal2)

		intVal3, _ := strconv.ParseInt("23232s", 0, 64)
		fmt.Printf("Mi variable es: %d\n", intVal3)
	}

	{
		var A byte = 'A'
		var a byte = 'a'

		var R byte = 82
		var s byte = 115
		vector := []byte{65, 34, 97, 82}

		fmt.Println()
		fmt.Println("Value in ASCII code")
		fmt.Println(A)
		fmt.Println(a)
		fmt.Println(R)
		fmt.Println(string(s))
		fmt.Println(vector)
		fmt.Println(string(vector))
	}
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

func vectores() {
	var myInt8Var int
	fmt.Println("Mi variable es:", myInt8Var)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myInt8Var, unsafe.Sizeof(myInt8Var), unsafe.Sizeof(myInt8Var)*8)

	var myArrayVar [6]int
	fmt.Println(myArrayVar)

	myArrayVar[1] = 2
	myArrayVar[2] = 3
	myArrayVar[3] = 4
	fmt.Println(myArrayVar)
	fmt.Println("size ", len(myArrayVar))
}

func slices() {
	// slices
	var slice1 []int
	fmt.Println("Mi variable es:", slice1)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", slice1, unsafe.Sizeof(slice1), unsafe.Sizeof(slice1)*8)
	fmt.Println("size ", len(slice1))

	slice1 = append(slice1, 12, 13, 45, 34, 47)
	fmt.Println("Mi variable es:", slice1)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", slice1, unsafe.Sizeof(slice1), unsafe.Sizeof(slice1)*8)
	fmt.Println("size ", len(slice1))

	var myArrayVar [6]int
	myArrayVar[1] = 2
	myArrayVar[2] = 3
	myArrayVar[3] = 4
	// pasar por referencia
	mySlice := myArrayVar[2:4]
	fmt.Println("Mi variable es:", mySlice)
	fmt.Println("size ", len(mySlice))
	fmt.Println(&myArrayVar[2])
	fmt.Println(&mySlice[0])

	mySlice[0] = 90
	fmt.Println(myArrayVar[2])
	fmt.Println(mySlice[0])

	fmt.Println(myArrayVar)
	// tomar desde el indice 0 hasta el 2
	fmt.Println(myArrayVar[:2])

	// tomar desde el indice 2 hasta el n
	fmt.Println(myArrayVar[2:])

	slice := make([]int, 10)

	fmt.Println(slice)
}

func maps() {
	m1 := make(map[int]string)
	m1[1] = "a"
	m1[2] = "b"
	m1[3] = "c"

	fmt.Println(m1)
	fmt.Println(m1[1])

	delete(m1, 2)
	fmt.Println(m1)

	m1[1] = "a2"
	fmt.Println(m1)

	v, exist := m1[1]
	fmt.Println(v)
	fmt.Println(exist)

	m2 := map[int]string{
		1: "hola",
		2: "mundo",
	}
	fmt.Println(m2)
}

func fors() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}

	fmt.Println(sum)

	for sum < 1000 {
		sum++
	}

	fmt.Println(sum)

	for {
		if sum > 1000 {
			break
		}
		sum++
	}

	fmt.Println(sum)

	arr := []int{1, 2, 3, 4, 5, 6}
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
