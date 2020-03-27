package main

import (
	"fmt"
	"sort"
)

//排序
func sortTest() {

	//直接改变给定的数组，不返回新值
	strs := []string{"i", "l", "y", "f"}
	sort.Strings(strs)
	fmt.Println(strs)

	nums := []int{5, 7, 9, 1, 4}
	sort.Ints(nums)
	fmt.Println(nums)

	//是否已经排序
	result := sort.IntsAreSorted(nums)
	fmt.Println(result)

	//突发奇想，切片能不能排序呢？
	slice1 := make([]int, 0, 5)
	slice1 = append(slice1, 5, 9, 7, 1, 5)
	sort.Ints(slice1)
	fmt.Println(slice1)
}

//自定义排序
//排序函数
type Bylength []string

func (a Bylength) Len() int      { return len(a) }
func (a Bylength) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Bylength) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

func sortMyTest() {
	text := []string{"one", "second", "them"}
	sort.Sort(Bylength(text))
	fmt.Println(text)
}

func main() {
	sortMyTest()
}
