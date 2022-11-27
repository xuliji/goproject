package model

import "errors"

type person struct {
	Name   string
	age    int
	salary float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

/*
set()和get()方法
*/
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		err := errors.New("年龄范围不正确")
		panic(err)
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSalary(salary float64) {
	if salary > 0 {
		p.salary = salary
	} else {
		err := errors.New("年龄范围不正确")
		panic(err)
	}
}

func (p *person) GetSalary() float64 {
	return p.salary
}
