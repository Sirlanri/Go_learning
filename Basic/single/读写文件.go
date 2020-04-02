package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//检查函数，可以少敲很多下键盘哦
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readTest() {
	//相对路径，开头不要加/
	data1, err := ioutil.ReadFile("data/text1.txt")
	check(err)
	fmt.Println(string(data1))

	//稍微高级点的做法
	file1, err := os.Open("data/text1.txt")
	check(err)
	b1 := make([]byte, 5) //最多读取5个字符
	n1
}

func main() {
	readTest()
}
