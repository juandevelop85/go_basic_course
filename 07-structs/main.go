package main

import (
	"encoding/json"
	"fmt"

	"github.com/juandevelop85/go_basic_course/07-structs/structsInterfaces/structs"
)

func main() {
	var p1 structs.Product
	p1.Id = 2
	p1.Name = "Prueba"
	fmt.Println(p1)

	p2 := structs.Product{}
	p2.Id = 2
	p2.Name = "Prueba_2"
	fmt.Println(p2)

	p3 := structs.Product{1, "Prueba3", structs.Type{}, 4, 2.3, nil}
	fmt.Println(p3)

	// Forma correcta de inicializar el struct, indicando los key que se quieran inicializar
	p4 := structs.Product{
		Id:    1,
		Name:  "Prueba4",
		Price: 2000,
		Count: 4,
		Type: structs.Type{
			Id:          2,
			Code:        "23132",
			Description: "Una_descripcion",
		},
		Tags: []string{"a", "b", "c"},
	}
	fmt.Println(p4)

	// Retorna el json en ascii
	v, err := json.Marshal(p4)
	fmt.Println(string(v))
	fmt.Println(err)

	fmt.Println("precio total", p4.TotalPrice())

	p4.SetName("Actualizacion")
	v, _ = json.Marshal(p4)
	fmt.Println(string(v))

	p4.AddTags("F", "G", "H", "J")
	v, _ = json.Marshal(p4)
	fmt.Println(string(v))

	car := structs.NewCar(1)

	car.AddProducts(p3, p4)

	vCar, _ := json.Marshal(car)
	fmt.Println("---")
	fmt.Println("Mi carrito", string(vCar))
	fmt.Println("---")


}
