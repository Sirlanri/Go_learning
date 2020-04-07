package main

import (
	"bufio"
	"fmt"
	"io"
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
	n1, err := file1.Read(b1)
	check(err)
	fmt.Println("读取5个字符后", n1)

	//到一个已知的位置，开始读取
	o2, err := file1.Seek(6, 0)
	check(err)
	b2 := make([]byte, 4)
	n2, err := file1.Read(b2)
	check(err)
	fmt.Println("seek结果：", n2, o2)

	//更健壮的实现
	o3, err := file1.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(file1, b3, 2)
	check(err)
	fmt.Println("更强的方法", n3, o3)

	//带缓冲的
	r4 := bufio.NewReader(file1)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Println("带缓冲的", b4)
}

func writeTest() {
	test2 := []byte("wdnmd写入")
	err := ioutil.WriteFile("data/text2.text", test2, 0644)
	if err != nil {
		println(err)
	}

	//更细颗粒的写入
	file2, err := os.Create("data/text3.txt")
	if err != nil {
		println(err)
	}
	defer file2.Close()
	intarr := []byte{1, 2, 3, 4, 5} //字节切片
	lenint, err := file2.Write(intarr)
	if err != nil {
		println(err)
	}
	fmt.Println("写入的数据长度", lenint)
	//将缓冲区的信息写入硬盘
	file2.Sync()

}

func main() {
	writeTest()
}
