package main

import "fmt"

func main() {
	var height, money float32
	var handsome bool
	fmt.Println("请输入身高,财富,帅:")
	fmt.Scanf("%f,%f,%t", &height, &money, &handsome)
	if height > 180.0 && money > 1.0e8 && handsome {
		fmt.Println("非他不嫁")
	} else if height > 180.0 || money > 1.0e8 || handsome {
		fmt.Println("比上不足,比下有余")
	} else {
		fmt.Println("不嫁")
	}
}
