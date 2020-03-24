package main

import (
	"fmt"
	"strings"
)

//MyCode64 int64别名
type MyCode64 = int64

//MyInt 自定义类型
type MyInt int64

type person struct {
	name string
	city string
	age  int
}

type test struct {
	a int8
	b int8
	c int8
}

type student struct {
	name string
	age  int
}

//NiMing 匿名结构体，一个类型只能有一个
type NiMing struct {
	string
	int8
	int32
}

func newPerson(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

//方法与函数不同，是特定的类型，p即接受者，分为值类型和指针类型
//非本包的类型不能定义方法
func (p person) Dream() {
	fmt.Printf("%s梦想是学号GO语言!\n", p.name)
}

func (p *person) SetAge(newAge int) {
	p.age = newAge
}
func (p person) Set2Age(newAge int) {
	p.age = newAge
}

//SayHello 测试方法
func (m MyInt) SayHello() {
	fmt.Println("新类型的值", m)
}

func structTest1() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
		{name: "哈哈", age: 0},
	}

	// fmt.Println(stus)
	// stu 是切片中的一个数据
	for _, stu := range stus {
		fmt.Printf("%#v , %#p , name=%#v\n", stu, &stu, stu.name) //遍历正确
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	//stus外层是切片，每次循环的地址是最后一个student
	/*
		大王八 => 大王八
		小王子 => 大王八
		娜扎 => 大王八
	*/
	fmt.Println(m)
}
func main() {
	//1,使用new初始化，指针
	p := new(person)
	p.name = "测试者"
	//2,使用var 一样初始化, 非指针
	var p1 person
	p1.name = "大海"
	p1.age = 28
	p1.city = "北京"

	p3 := &person{} //地址

	var user struct {
		Name string
		Age  int
	}
	user.Name = "华为"
	user.Age = 100

	p4 := person{
		name: "小王子",
		city: "上海",
		age:  10,
	}
	p5 := &person{
		name: "小可爱",
		city: "安徽",
		age:  1,
	}
	fmt.Printf("value=%#v, type=%T\n", p, p)
	fmt.Printf("value=%#v, type=%T\n", p1, p1)
	fmt.Printf("value=%#v, type=%T\n", p3, p3)
	fmt.Printf("value=%#v, type=%T\n", user, user)
	fmt.Printf("value=%#v, type=%T\n", p4, p4)
	fmt.Printf("value=%#v, type=%T\n", p5, p5)

	fmt.Println(strings.Repeat("---分割线--", 10))
	t := test{
		1,
		2,
		3,
	}
	fmt.Printf("t.a pointer's address : %p\n", &t.a) //地址是连续的块
	fmt.Printf("t.b pointer's address : %p\n", &t.b)
	fmt.Printf("t.c pointer's address : %p\n", &t.c)

	fmt.Println(strings.Repeat("---分割线--", 10))

	structTest1()

	p6 := newPerson("我爱罗", "海淀", 100)
	fmt.Println(p6)
	p6.Dream()
	p6.SetAge(20)
	fmt.Println(p6.age)
	p6.Set2Age(50) //值类型的接受者，值复制的修改只影响方法内
	fmt.Println(p6.age)

	var myint MyInt
	myint.SayHello()

	var nm NiMing = NiMing{
		"匿名者",
		127, //int8 有符号最大值
		5678,
	}
	fmt.Println(nm)
}
