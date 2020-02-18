package main

import "fmt"

type Goods struct {
	name  string
	price int
}

func (g *Goods) Getname() string {
	return g.name
}

func (g *Goods) Getprice() int {
	return g.price
}

func (g *Goods) Setname(name string) {
	g.name = name
}

func (g *Goods) Setprice(price int) {
	g.price = price
}

func TestObj() {
	good1 := Goods{
		"小米", 2,
	}
	good1.Setname("雷军的")
	good1.Setprice(998)

	fmt.Println("名字是", good1.Getname())
	fmt.Println("价格是", good1.Getprice())

}

//下面学学继承，嗯....go真奇怪，面向对象不好吗
type Apple struct {
	Goods
	color string
}

func (a *Apple) Setcolor(color string) {
	a.color = color
}

func TestApple() {
	iphont := Apple{
		Goods{"iPhone XS max", 9999},
		"土豪金",
	}
	fmt.Printf("名字是%s,价格%d,颜色%s", iphont.Getname(), iphont.Getprice(), iphont.color)
}

//再来个接口~这名字有点像java咯
type Salealbe interface {
	Sell()
}

func (Apple) Sell() {
	fmt.Println("Apple 实现了Sell接口~")
}
func TestInterface() {
	iphonex := Apple{
		Goods{"苹果11", 10099},
		"玫瑰金",
	}
	var i Salealbe
	i = &iphonex
	i.Sell()
}

func objtest() {
	TestInterface()
}
