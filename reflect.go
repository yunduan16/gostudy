//反射
package main

import (
	"fmt"
	"reflect"
)

//User 用户结构体
type User struct {
	ID       int
	UserName string
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type=%v, kind=%v\n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is %d\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %d\n", float64(v.Float()))
	}
}

func reflectDemo1() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)
	c := &a
	reflectType(c)
	d := &User{ID: 100, UserName: "我爱罗"}
	reflectType(d)

	e := User{}
	reflectType(e)
}

func reflectDemo2() {
	var a float32 = 3.14
	var b int64 = 100
	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b) // type is int64, value is 100
	// 将int类型的原始值转换为reflect.Value类型
	c := reflect.ValueOf(10)
	fmt.Printf("type c :%T value=%v\n", c, c) // type c :reflect.Value
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		// v.SetInt(200) //修改的是副本会议，必须用地址更改
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		//反射中使用Elem获取指针对应的值
		v.Elem().SetInt(200)
	}
}
func main() {
	reflectDemo1()

	reflectDemo2()

	var a int64 = 1000
	reflectSetValue1(a)
	fmt.Println(a)
	reflectSetValue2(&a) //必须传指针类型
	fmt.Println(a)
}
