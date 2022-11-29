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

// GetLoop SetLoop get和set方法
func (view *customView) GetLoop() bool {
	return view.loop
}

func (view *customView) SetLoop(loop bool) {
	view.loop = loop
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
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			view.List()
		case "5":
			fmt.Println("退出")
			view.SetLoop(false)
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
