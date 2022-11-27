package main

import "fmt"

// map使用for-range实现遍历
func main() {
	cities := map[string]string{
		"no1": "北京",
		"no2": "上海",
		"no3": "成都",
		"no4": "合肥",
	}
	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	students := make(map[string]map[string]string)
	students["01"] = map[string]string{
		"name": "xuliji",
		"sex":  "男",
	}

	students["02"] = map[string]string{
		"name": "marry",
		"sex":  "女",
	}

	for k1, v1 := range students {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v\n", k2, v2)
		}
	}
}
