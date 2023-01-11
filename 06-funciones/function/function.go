package function

import (
	"errors"
	"fmt"
)

type Operation int

// es como hacer un enum
const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Add(x int, y int) int {
	return x + y
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Print(value)
	}
}

func Calc(op Operation, x, y float64) (float64, error) {

	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("(y) debe ser mayor a cero")
		}
		return x / y, nil
	case MUL:
		return x * y, nil
	default:
		return 0, errors.New("operacion no valida")
		
	}
}

func Split(v int) (x, y int) {
	x = v * 4 / 9
	y = v - x
	return
}

func MSum(values ...float64) float64{
	var sum float64
	for _, v := range values {
		sum += v
	}

	return sum
}

func MOperations(op Operation, values ...float64) (float64, error) {

	if len(values) == 0 {
		return 0, errors.New("Debe ingresar valores")
	}

	sum := values[0]

	for _, v := range values[1:] {

		switch op {
		case SUM:
			sum += v
		case SUB:
			sum -= v
		case DIV:
			if v == 0 {
				return 0, errors.New("(y) debe ser mayor a cero")
			}
			sum /= v
		case MUL:
			sum *= v
		}
	}
	return sum, nil
}