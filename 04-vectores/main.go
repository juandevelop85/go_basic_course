package vectors

import (
	"fmt"
	"unsafe"
)

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
