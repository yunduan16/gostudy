//本节关注结构体的嵌套、继承、序列化
package main

import (
	"encoding/json"
	"fmt"
)

//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet int8
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

// Student 学生
// 结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。键值对之间使用一个空格分隔
type Student struct {
	ID     int    `json:"id"` //定义json序列化的Tag
	Gender string `key1:"value1" key2:"value2"`
	Name   string
}

//Class 课程
type Class struct {
	Title    string
	Students []*Student
}

func structDemo2() {
	c := &Class{
		Title:    "高三2班",
		Students: make([]*Student, 0, 20),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			ID:     i,
			Gender: "男",
			Name:   fmt.Sprintf("stu%02d", i),
		}
		c.Students = append(c.Students, stu)
	}
	//json序列化
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//json反序列化
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%$v\n", c1)
}
func structDemo1() {
	d1 := Dog{
		Feet: 4,
		Animal: &Animal{ //初始化时引用类型key不带*
			name: "乐乐",
		},
	}
	d1.move()
	d1.wang()
}

func main() {
	structDemo1()

	structDemo2()
}
