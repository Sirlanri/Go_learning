package main

import (
	"encoding/json"
	"fmt"
	"time"

	"math/rand"
	"os"
	"regexp"
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

func collectionTest() {
	strs := []string{"peach", "apple", "pear", "plum"} //最后这个梅子我竟然不认识
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

}

//正则表达式
func zhengzeTest() {
	//直接使用，是否符合表达式
	match1, _ := regexp.MatchString("p([a-z])ch", "peaaach")
	fmt.Println(match1)
	//预编译（Python里不建议这样使用
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peaaach"))
	fmt.Println(r.FindString("peaonech peatwoch"))

}

//json解析
type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func json1Test() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	//切片和map编码
	slinceD := []string{"apple", "peach", "pear"}
	slinceB, _ := json.Marshal(slinceD)
	fmt.Println(string(slinceB))

	mapD := map[int]string{1: "fisrt", 2: "second", 3: "third"}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var data map[string]interface{}
	//实际解码和错误检查
	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}
	fmt.Println(data)
	//将num转化为float64
	num2 := data["num"].(float64)
	fmt.Println(num2)
	//访问嵌套的值
	strs := data["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	//把json解析到结构体，这样方便多了~
	str2 := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str2), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	//使用OS，可以作为http相应体
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"one": 1, "two": 2}
	enc.Encode(d)
}

//随机数，测试的时候经常用到
func randomTest() {
	//返回一个0.0<=f<1.0的数
	fmt.Println(rand.Float64() * 10)
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Float64())
	}

	//随机种子
	//教程说，如果不随机种子，会产生一样的随机数，可是这边很正常啊
	seed1 := rand.NewSource(time.Now().UnixNano())
	res1 := rand.New(seed1)
	fmt.Println("随机种子", res1.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
}
func main() {
	randomTest()
}
