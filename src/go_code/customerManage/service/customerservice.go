package service

import (
	"goproject/src/go_code/customerManage/model"
)

// CustomerService 完成对customer的操作    包括:增删改查
type CustomerService struct {
	customers []model.Customer
	// 表示当前有多少个客户
	customerNum int
}

// Customers 返回customer切片
func (c *CustomerService) ReturnCustomers() []model.Customer {
	return c.customers
}

// NewCustomerService 工厂模式
func NewCustomerService() *CustomerService {
	cs := CustomerService{}
	cs.customerNum = 1
	customer := model.NewCustomer(0, 23, "xuliji", "男", "123",
		"123@qq.com")
	cs.customers = append(cs.customers, customer)
	return &cs

}

// AddCustomer 添加客户
func (c *CustomerService) AddCustomer(Age int, Name, Gender, Phone, Email string) {
	id := c.customerNum
	cus := model.NewCustomer(id, Age, Name, Gender, Phone, Email)
	c.customers = append(c.customers, cus)
	c.customerNum += 1
}

// FindByid 根据id返回切片下标
func (c *CustomerService) FindByid(Id int) int {
	for i, v := range c.customers {
		if v.Id == Id {
			return i
		}
	}
	return -1
}

// DeleteCustomer 删除客户
func (c *CustomerService) DeleteCustomer(id int) bool {
	Index := c.FindByid(id)
	if Index == -1 {
		return false
	}
	// 如何从切片删除一个元素
	c.customers = append(c.customers[:Index], c.customers[Index+1:]...)
	return true
}

// Updete 更新用户信息
func (c *CustomerService) Updete(Index int) bool {
	if Index == -1 {
		return false
	}

	return true
}
