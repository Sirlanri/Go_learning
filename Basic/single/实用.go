package main

import (
	"fmt"
	"sort"
	"strings"
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

//组合函数练习
//Index返回t出现的第一个位置，没有返回-1
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

//Include 如果存在，就返回true
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

//Any 有一个满足，就返回true
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		//f就是引用的func
		if f(v) {
			return true
		}
	}
	return false
}

//Filter 返回一个包含所有切片中满足条件f的字符串新切片
func Filter(vs string, f func(string) bool) []string {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//Map 返回一个对原始切片中所有字符串执行函数f后的新切片
func Map(vs []string, f func(string) bool) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
func collectionTest() {
	strs := []string{"peach", "apple", "pear", "plum"} //最后这个梅子我竟然不认识
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
}

func main() {
	sortMyTest()
}
