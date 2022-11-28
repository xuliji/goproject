package main

import "fmt"

type customView struct {
	key  string // 接收用户输入
	loop bool   //是否循环显示菜单
}

// GetLoop SetLoop get和set方法
func (view *customView) GetLoop() bool {
	return view.loop
}

func (view *customView) SetLoop(loop bool) {
	view.loop = loop
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
		fmt.Scanln(&(view.key))
		switch view.key {
		case "1":
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			fmt.Println("客户列表")
		case "5":
			fmt.Println("退出")
			view.SetLoop(false)
		default:
			fmt.Println("命令输出不正确")

		}
		if !view.loop {
			break
		}
	}

}

func main() {
	// 实例化结构体
	view := customView{loop: true}
	view.mainView()

}
