package service

import "goproject/src/go_code/customerManage/model"

// CustomerService 完成对customer的操作    包括:增删改查
type CustomerService struct {
	customers []model.Customer
	// 表示当前有多少个客户
	customerNum int
}

// Customers 返回customer切片
func (c CustomerService) Customers() []model.Customer {
	return c.customers
}

// NewCustomerService 工厂模式
func NewCustomerService() *CustomerService {
	cs := CustomerService{}
	cs.customerNum = 1
	customer := model.NewCustomer(0, 23, "xuliji", "男", "123456",
		"123456@qq.com")
	cs.customers = append(cs.customers, customer)
	return &cs

}

// 返回客户切片
