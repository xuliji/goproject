package main

import (
	"errors"
	"fmt"
	"goproject/src/go_code/customerManage/service"
)

type customView struct {
	key  string // 接收用户输入
	loop bool   //是否循环显示菜单
	// 增加一个CustomerService类型字段
	customService *service.CustomerService
}

func (cv *customView) List() {
	fmt.Println("--------------------" +
		"客户列表" + "--------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, v := range cv.customService.ReturnCustomers() {
		v.GetInfo()
	}
	fmt.Println("--------------------" +
		"客户列表完成" + "--------------------")
}

// AddCustomer 添加客户
func (cv *customView) AddCustomer() {
	var Age int
	var Name, Gender, Phone, Email string
	fmt.Println("--------添加客户--------")
	fmt.Println("姓名：")
	fmt.Scanln(&Name)
	fmt.Println("年龄：")
	fmt.Scanln(&Age)
	fmt.Println("性别：")
	fmt.Scanln(&Gender)
	fmt.Println("电话：")
	fmt.Scanln(&Phone)
	fmt.Println("电子邮件：")
	fmt.Scanln(&Email)

	// id需要系统分配
	cv.customService.AddCustomer(Age, Name, Gender, Phone, Email)

}

func (cv *customView) DeleteCustomer() {
	var id int
	var confirm string // 必须是string类型
	fmt.Println("--------删除客户--------")
	fmt.Println("请输入ID:")
	_, _ = fmt.Scanln(&id)
	fmt.Println("是否确认(Y or N):")
	_, _ = fmt.Scanln(&confirm)
	if confirm == "Y" {
		flag := cv.customService.DeleteCustomer(id)
		if flag {
			fmt.Println("--------删除成功--------")
		} else {
			fmt.Println("--------删除失败--------")
		}
	}

}

// 显示主菜单
func (view *customView) mainView() {
	for {
		fmt.Println("--------客户信息管理系统--------")
		fmt.Println("        1.添加客户信息         ")
		fmt.Println("        2.修改客户信息         ")
		fmt.Println("        3.删除客户         ")
		fmt.Println("        4.客户列表         ")
		fmt.Println("        5.退出         ")
		fmt.Print("请选择(1-5):")

		// 获取输入
		_, _ = fmt.Scanln(&(view.key))
		switch view.key {
		case "1":
			view.AddCustomer()
		case "2":
			fmt.Println("修改客户")
		case "3":
			view.DeleteCustomer()
		case "4":
			view.List()
		case "5":
			fmt.Println("退出")
			view.loop = false
		default:
			fmt.Println(errors.New("指令输入不正确!!!!"))

		}
		if !view.loop {
			break
		}
	}

}

func main() {
	// 实例化结构体
	view := customView{loop: true, key: "", customService: service.NewCustomerService()}
	view.mainView()

}
