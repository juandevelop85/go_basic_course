package variables

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {

	variables()
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
