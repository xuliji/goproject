package classes

// 1. 声明Heros结构体
type Heros struct {
	Name string
	Age  int
}

// 2.申明Heros结构体切片
type HeroSlice []Heros

// 3.实现Interface接口
// Len方法返回集合中的元素个数
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less方法报告索引i的元素是否比索引j的元素小
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

// Swap方法交换索引i和j的两个元素
func (hs HeroSlice) Swap(i, j int) {
	//temp := Heros{}
	//temp = hs[i]
	//hs[i] = hs[j]
	//hs[j] = temp
	// 等价于
	hs[i], hs[j] = hs[j], hs[i]
}
