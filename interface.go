//接口相关
//1,接口可以嵌套
//2,接口可以被多个类型实现
//3,只要结构体或结构体内嵌套的结构体实现该接口，则可以使用该接口的方法
//4.任何类型都实现了空接口
package main

import "fmt"

//Sayer 说接口
type Sayer interface {
	say()
	// Move()
}

type dog struct{}
type cat struct{}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

//People 接口
type People interface {
	Speak(string) string
}

//Student 接口
type Student struct{}

//3,只要结构体或结构体内嵌套的结构体实现该接口，则可以使用该接口的方法

//WashingMachine 接口
type WashingMachine interface {
	wash()
	dry()
}

type dryer struct{}

func (d dryer) dry() {
	fmt.Println("甩干机甩一甩")
}

type haier struct {
	dryer
}

func (h haier) wash() {
	fmt.Println("洗衣机转圈圈")
}

//Speak 方法
func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

//空接口作为参数，可以支持任何类型
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

//使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量
//指针类则只能指针赋值
func interfaceDemo1() {
	var x Sayer
	d1 := dog{}
	d1.say()
	c1 := cat{}
	c1.say()
	x = d1 //仅当dog实现Sayer全部方法才可复制
	x.say()
	var wangcai = &dog{}
	x = wangcai
	x.say()
}

//测试指针类型接受者，是否可以通过值类型直接调用它的方法
func interfaceDemo2() {
	// var peo People = Student{}	//编译出错
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

func interfaceDemo3() {
	h1 := haier{}
	h1.dry()
	h1.wash()
}

//判断空接口变量的类型
func interfaceDemo4() {
	var x interface{}
	x = "hello world"
	v, ok := x.(string) //通过x.(T)可以预测空接口变量可能类型
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
func main() {
	interfaceDemo1()

	interfaceDemo2()

	interfaceDemo3()

	interfaceDemo4()

	a := make([]string, 2)
	a[0] = "北京"
	a[1] = "欢迎"
	a = append(a, "你")
	show(a)
}
