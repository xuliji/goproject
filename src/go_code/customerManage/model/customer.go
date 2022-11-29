package model

import "fmt"

// Customer 表示客户信息
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func NewCustomer(id, Age int, Name, Gender, Phone, Email string) Customer {
	return Customer{
		Id:     id,
		Name:   Name,
		Gender: Gender,
		Age:    Age,
		Phone:  Phone,
		Email:  Email,
	}
}

func (cus Customer) GetInfo() {
	fmt.Printf("%v\t%v\t%v\t%d\t%v\t%v\n", cus.Id, cus.Name, cus.Gender, cus.Age, cus.Phone, cus.Email)
}
