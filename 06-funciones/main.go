package main

import (
	"fmt"

	"github.com/juandevelop85/go_basic_course/06-funciones/function"
)

func main() {
	fmt.Println(function.Add(3,4))
	function.RepeatString(10, "juan")

	v, err := function.Calc(function.SUM, 3, 6)
	fmt.Println("value: ", v, " Error: ", err)

	v, err = function.Calc(function.DIV, 3, 0)
	fmt.Println("value: ", v, " Error: ", err)

	v1, v2 := function.Split(20)
	fmt.Println("value1: ", v1, " value2: ", v2)
	
	mSumResult := function.MSum(2,4,6,5,4,8,9,10,1)
	fmt.Println("value1: ", mSumResult)

	v, err = function.MOperations(function.SUM, 3, 6, 5, 7, 8)
	fmt.Println("value: ", v, " Error: ", err)

	v, err = function.MOperations(function.DIV, 3, 5, 2, 6, 8)
	fmt.Println("value: ", v, " Error: ", err)
	
	v = function.FactoryOperation(function.SUM)(3, 4)
	fmt.Println("value: ", v)

	v = function.FactoryOperation(function.MUL)(3, 4)
	fmt.Println("value: ", v)
}
