package finds

import "fmt"

// 顺序查找
func OrderFind(name_slice []string, key string) {
	for i, s := range name_slice {
		if s == key {
			fmt.Printf("%v在列表中,下标为%d\n", key, i)
			break
		} else if i == len(name_slice)-1 {
			fmt.Printf("没有找到%v\n", key)
		}
	}

}
