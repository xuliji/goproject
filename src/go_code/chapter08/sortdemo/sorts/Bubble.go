package sorts

import "fmt"

// 冒泡排序
func Bubble(slice []int) {
	fmt.Println("排序前的切片：", slice)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	for j := 0; j < len(slice)-1; j++ {
		for i := 0; i < len(slice)-1-j; i++ {
			if slice[i] > slice[i+1] {
				temp := slice[i]
				slice[i] = slice[i+1]
				slice[i+1] = temp
			}
		}
	}
	fmt.Println("排序后的切片：", slice)

}
