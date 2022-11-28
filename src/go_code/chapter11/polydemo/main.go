package main

import (
	"fmt"
	"goproject/src/go_code/chapter11/polydemo/utils"
)

// 多态数组
func main() {
	var usbArr [3]utils.Usb

	usbArr[0] = utils.Phone{Name: "小米"}
	usbArr[1] = utils.Creame{Name: "ikon"}
	usbArr[2] = utils.Phone{Name: "华为"}
	fmt.Println(usbArr)
}
