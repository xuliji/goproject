package finds

import (
	"fmt"
)

// 二分查找:查找有序序列
func BinaryFind(int_slice []int, find_val int, left_index int, right_index int) {
	middle := (left_index + right_index) / 2
	if left_index > right_index {
		fmt.Println("找不到")
		return
	}

	if int_slice[middle] > find_val {
		BinaryFind(int_slice, find_val, left_index, middle-1)
	} else if int_slice[middle] < find_val {
		BinaryFind(int_slice, find_val, middle+1, right_index)
	} else if int_slice[middle] == find_val {
		fmt.Printf("找到n=%d, 下标为%d\n", find_val, middle)
	}

}
