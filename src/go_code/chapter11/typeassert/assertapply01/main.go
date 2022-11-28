package main

import (
	"fmt"
	"goproject/src/go_code/chapter11/polydemo/utils"
)

// 类型断言应用 重点
func main() {
	var usbArr [3]utils.Usb

	usbArr[0] = utils.Phone{Name: "小米"}
	usbArr[1] = utils.Creame{Name: "ikon"}
	usbArr[2] = utils.Phone{Name: "华为"}

	com := utils.Computer{}
	for _, val := range usbArr {
		com.Working(val)
		fmt.Println()
	}
}
