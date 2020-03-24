//结构体反射

package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name" range:"all"`
	Score int    `json:"score"`
}

func test1() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}

func (s student) Study() {
	msg := "好好学习，天天向上！"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() {
	msg := "早睡早起，方能养生！"
	fmt.Println(msg)
	return msg
}
func main() {
	test1()
}
