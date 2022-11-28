package utils

import "fmt"

// Usb 申明一个接口
// 接口不能包含任何变量  让结构体实现方法
type Usb interface {
	// Start 申明了两个没有实现的方法
	Start()
	Stop()
}

// Phone 手机类
type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println("手机开始工作....")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作....")
}

func (p Phone) Call() {
	fmt.Println("打电话给xxxxxx")
}

// Creame 相机类
type Creame struct {
	Name string
}

func (c Creame) Start() {
	fmt.Println("相机开始工作....")
}

func (c Creame) Stop() {
	fmt.Println("相机停止工作....")
}

// Computer 电脑类
type Computer struct {
}

// 接受Usb接口变量
func (com Computer) Working(usb Usb) {
	usb.Start()
	// 如果usb是指向一个Phone变量的,则还需要调用一个Call方法
	phone, flag := usb.(Phone)
	if flag {
		phone.Call()
	}
	usb.Stop()
}
