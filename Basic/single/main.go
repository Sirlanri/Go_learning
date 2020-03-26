package main

import (
	"fmt"
	"math/rand"
	"time"
)

func database1() {
	fmt.Println("你好鸭GO，好久不见呢")

}
func changemain() {
	one := 1
	fmt.Println("数字的地址是", &one)
	a := 10
	change(&a)
	fmt.Println(a)
}
func change(a *int) {
	*a += 10 //唤醒了我那基础不牢的C语言回忆
}
func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

//接收一个列表
func getlist(l []int) int {
	sum := 0
	for _, single := range l {
		sum += single
	}
	return sum
}
func getlistmain() {
	//var listb [10]int //这个声明方式很独特，阴吹斯汀
	slice := []int{7, 9, 3, 5, 1}
	var x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)

	trytwo := []int{7, 8, 91, 2, 45, 56}
	var y = min(trytwo...)
	fmt.Println("第二次测试", y)
}

func plustest() {
	a := 1
	defer fmt.Println(a)
	a = 10
	fmt.Println("下面的")

}
func feibonaqie(num int) {
	one := 1
	two := 1
	var reslist [100]int
	for i := 0; i < num; i++ {
		one = two
		two = one + two
	}
	fmt.Println(reslist)
}
func feibo2(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = feibo2(n-1) + feibo2(n-2)
	}
	return
}

func feibo2test() {
	start := time.Now()
	result := 0
	for i := 0; i < 10; i++ {
		result = feibo2(i)
		fmt.Println(result)
	}
	end := time.Now()
	period := end.Sub(start)
	fmt.Printf("时长为%s", period)
}
func jiecheng(num int) (res int) {
	if num == 0 {
		res = 1
	} else {
		res = num * jiecheng(num-1)
	}
	return
}

func bibaotest() {
	//闭包方式的函数，更简洁
	he := func(a, b int) {
		fmt.Println(a + b)
	}
	he(1, 2)
}
func kaobei2(a *[3]int) {
	sum := 0
	for _, i := range a {
		sum += i
	}

}

//切片，Python里面已经很熟悉啦
func qiepiantest() {
	var arr1 [6]int
	//通过for赋值
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i + rand.Intn(100)
	}
	//用切片来传递参数，很不错~
	fmt.Println("求和为", qiepianRec(arr1[:]))

	//make用法,返回的是一个对象；new返回的是指针
	arr3 := make([]int, 5)
	for i := 0; i < 5; i++ {
		arr3[i] = rand.Intn(10)
	}
	fmt.Println("make的数组", arr3)

	arr4 := new([5]int)
	for i := 0; i < 5; i++ {
		arr4[i] = rand.Intn(10)
	}
	fmt.Println("new方法的数组：", arr4)

	//三个参数make，长度5，数量10
	arr5 := make([]int, 5, 10)
	for i := 0; i < 5; i++ {
		arr5[i] = rand.Intn(10)
	}
	fmt.Println("填5个值", arr5)

	//一个奇怪的想法
	arr6 := [3]int{1, 2, 3}
	fmt.Println(arr6[2:3])
	fmt.Println("测试copy")

	arr7 := []int{1, 2, 3}
	arr8 := make([]int, 5)
	copy(arr8, arr7)
	arr8[0] = 10
	fmt.Println("over", arr7)

}
func qiepianRec(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

//超级好用的map，也就是Python里的字典
func zidiantest() {
	map1 := make(map[int]string) //发现越来越喜欢用make了
	map2 := map[int]string{1: "one", 2: "two"}

	map1[1] = map2[1]
	map1[2] = "second"
	map1[3] = "third"

	//原来这if要这样写，好麻烦
	if value1, isok := map1[3]; isok {
		fmt.Println(value1)
	}
	delete(map1, 1)
	//copy(map1, map2)
	fmt.Println(map1)
}

func jiegoutitest() {
	type first struct {
		name string
		num  int
	}
	var one first
	one.name = "onename"
	one.num = 100

	//new，返回的是指针，但是用起来没啥区别
	second := new(first)
	second.name = "第二个"
	second.num = 101

	third := &first{"直接初始化，返回指针", 110}
	four := first{"返回本体，不是指针", 102}

	fmt.Println(second, third, four)

	//做个小栗子
	type Rectangle struct {
		chang int
		kuan  int
	}

	bus1 := new(Bus)
	bus1.num = 100
	bus1.SetName("USA")
	fmt.Println(bus1.GetName())

}

type Car struct {
	name string
	num  int
}
type Bus struct {
	Car
	rongliang int
}

//熟悉的set get方法
func (c *Car) SetName(getname string) {
	c.name = getname
}
func (c *Car) GetName() string {
	return c.name
}

//例子的复习啦！
func map2() {
	m1 := make(map[int]string)
	m1[1] = "第一条"
	m1[2] = "第二个！"
	m1[3] = "再来一个"
	m1[4] = "赠送"
	delete(m1, 2)

	for key, value := range m1 {
		fmt.Println("key", key, " ", value)
	}

}

//切片传入参数，一直是很懵逼的
func useslice2() {
	slice1 := make([]int, 0, 3)
	slice1 = append(slice1, 1, 2, 3)

	arr1 := []int{1, 2, 3}
	arr1 = append(arr1, 4)
	recarr(arr1)
	// 如果你有一个含有多个值的 slice，想把它们作为参数
	// 使用，你要这样调用 `func(slice...)`。
	recslice(slice1...)
}
func recarr(nums []int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("数组求和结果为 ", sum)
}
func recslice(nums ...int) {
	sum := 0
	fmt.Println("尝试输出第一个：", nums[0])
	for _, num := range nums {
		sum += num
	}
	fmt.Println("切片求和结果为 ", sum)

}

//闭包，教程说很重要，但是我也没怎么用过
//我的理解是，这个函数，返回值是一个函数
func bibao2(initial int) func() int {
	number := initial
	return func() int {
		number++
		return number
	}
}
func bibao2Test() {
	num1 := bibao2(5)

	for i := 0; i < 3; i++ {
		fmt.Println(num1())
	}
}

func test2() {
	selectTest1()

}
