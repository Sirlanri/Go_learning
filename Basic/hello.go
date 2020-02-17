package main

import "fmt"

func main() {
	fmt.Println("hello matherfucker")
	var i int = 1
	x, y := 2, 3
	qiuhe := i + x + y
	fmt.Println(qiuhe)
	var arr1 [10]int
	plus := 10
	for i := 0; i < 10; i++ {
		arr1[i] = plus
		plus += 10
	}
	println(arr1[9])
	fmt.Println(arr1[8])

	m := make(map[string]string)
	m["hello"] = "fuckoff"
	m1 := m
	println(m1["hello"])

	n := make(map[int]string)
	n[1] = "first"
	n1 := n
	println(n1[1])
}
