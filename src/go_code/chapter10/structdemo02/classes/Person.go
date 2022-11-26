package classes

// 结构体中指针, slice, map都需要先make再使用
type Person struct {

	// 字段首字母大写才能跨包使用
	Ptr   *int
	Map1  map[string]string
	Slice []int
}
