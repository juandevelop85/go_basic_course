package conditionales

import (
	"fmt"
)

func main() {
	maps()
	fors()
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
