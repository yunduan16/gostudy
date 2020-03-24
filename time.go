package main

import (
	"fmt"
	"time"
)

func timeDemo1() {
	now := time.Now()
	year, month, day := now.Date()
	fmt.Println(year, month, day)
	fmt.Println(now.Year(), now.Month(), now.Day(), now.Weekday())
}

func timeDemo2() {
	now := time.Now()
	fmt.Println(now.Format("2006年01月02日15时04分00秒")) //2006 1 2 3 4  2006年1月2号下午3点4分	GO诞生时间
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Println(timestamp1)
	fmt.Println(timestamp2)
}

func timeDemo3(timestamp int64, nextSecond int64) {
	timeObj := time.Unix(timestamp, nextSecond*1e9) //将时间戳转为时间格式,第二个参数表示增加的纳秒
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

//时间间隔方法
//time.Add()	time.Sub()求两个时间之差
func timeDemo4() {
	now := time.Now()
	currentTime := now.Add(time.Hour * 10)
	fmt.Println(currentTime)
}

//设置定时器Tick，本质是channel
func timeDemo5() {
	tick := time.Tick(time.Second * 2)
	for i := range tick {
		fmt.Println("定时输出一些东西", i)
	}
}

//解析时间字符串
func timeDemo6() {
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(now.Sub(timeObj))
}
func main() {
	timestamp := time.Now().Unix()
	timeDemo1()

	timeDemo2()

	timeDemo3(timestamp, 86400)

	timeDemo4()

	// timeDemo5()

	timeDemo6()
}
