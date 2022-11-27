package main

import "fmt"

// Usb 申明一个接口
type Usb interface {
	// Start 申明了两个没有实现的方法
	Start()
	Stop()
}

// Phone 手机类
type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作....")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作....")
}

// Creame 相机类
type Creame struct {
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
func (com Computer) working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	// 先创建结构体变量
	com := Computer{}
	phone := Phone{}
	camera := Creame{}
	com.working(phone)
	com.working(camera)

}
