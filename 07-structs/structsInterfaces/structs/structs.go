package structs

type Product struct {
	Id    uint     `json:"id"`
	Name  string   `json:"name"`
	Type  Type     `json:"type"`
	Count uint16   `json:"count"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

type Type struct {
	Id          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (p Product) TotalPrice() float64 {
	return float64(p.Count) * p.Price
}

func (p *Product) SetName(name string) {
	p.Name = name
}

func (p *Product) AddTags(name ...string) {
	p.Tags = append(p.Tags, name...)
}

type Car struct {
	Id       uint      `json:"id"`
	UserId   uint      `json:"user_id"`
	Products []Product `json:"products"`
}

func (c *Car) AddProducts(products ...Product) {
	c.Products = append(c.Products, products...)
}

func NewCar(userId uint) Car {
	return Car{UserId: userId}
}
