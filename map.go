package main

import (
	"fmt"
	"strings"
)

// map[KeyType]ValueType，初始置nil

func mapDemo() {
	m := map[string]int{
		"user_id": 1000,
		"age":     22,
	}
	fmt.Println(m) //打印时key值按照字母顺序，而非自定义填充顺序
	v, ok := m["age"]
	if ok {
		fmt.Println("key:age is exist, value=", v)
	} else {
		fmt.Println("查无此人")
	}
}

//map遍历
func mapDemo2() {
	m := make(map[string]int, 0)
	m["张三"] = 55
	m["李四"] = 80
	m["薛宝钗"] = 89
	m["葛优"] = 99
	//遍历顺序随机不固定
	for i, v := range m {
		fmt.Printf("index=%v value=%v\n", i, v)
	}
}

//删除指定key
func mapDemo3() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	delete(scoreMap, "小明") //将小明:100从map中删除
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}

//外层为切片，内层为map
func mapDemo4() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("outer slice inner map index=%v value=%v\n", index, value)
	}
	fmt.Println("after init")
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["birth"] = "2019-01-01"
	fmt.Println(mapSlice)
}

//外层map，内层切片
func mapDemo5() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap, "after init")
	key := "中国"
	if value, ok := sliceMap[key]; !ok {
		value = make([]string, 0, 2)
		value = append(value, "北京", "上海", "安徽")
		sliceMap[key] = value
	}
	fmt.Println(sliceMap)
}

//写出一个字符串中单词重复次数
func mapTest1(s string) map[string]int {
	fmt.Println("before of starting statistics, print input string:", s)
	arr := strings.Split(s, " ")
	expectArr := []string{",", ".", "!", "，", "。"}
	var m = make(map[string]int, 5)
	for _, v := range arr {
		// if inArray([]interface{}{",", ".", "!"}, v) {
		if inArray(v, expectArr) {
			continue
		}
		fmt.Println(v)
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	return m
}

//inArray 是否在数组中
func inArray(s interface{}, d []string) bool {
	for _, v := range d {
		if s == v {
			return true
		}
	}
	return false
}

func mapTest2() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 13)
	fmt.Printf("%v\n", s)
	m["aiqiyi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", m["aiqiyi"])
}

func main() {
	/*
		mapDemo()

		mapDemo2()

		mapDemo3()
	*/
	mapDemo4()

	mapDemo5()

	params1 := "how do you do"
	params2 := "dear my friend tina, you are my best friend . i love you tina very much . good nice to you !"
	fmt.Println(mapTest1(params1))
	fmt.Println(mapTest1(params2))

	mapTest2()
}
