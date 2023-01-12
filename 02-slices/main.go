package slices

import (
	"fmt"
	"unsafe"
)

func main() {

	slices()
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
