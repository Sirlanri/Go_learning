package main

//有点不理解接口，准备自己敲一遍

//薪资计算器接口
type SalaryCalculator interface {
	CalculateSalary() int
}

//普通挖掘机员工
type Contract struct {
	empId    int
	basicpay int
}

//蓝翔毕业的
type Permanent struct {
	empId    int
	basicpay int
	jj       int //多余的奖金
}

func CalculateSalary(p Permanent) int {
	return p.basicpay + p.jj
}
func CalculateSalary(c Contract) int {
	return c.basicpay
}

//总开支
func totalExpense(s []SalaryCalculator) { //这...不就是java里的多态吗？
	expense := 0
	for _, v := range s { //range就相当于js里的ForEach吧
		expense = expense + v.CalculateSalary()
	}
	fmt.Println("总开支是", expense)
}

func jiekou() {

}
