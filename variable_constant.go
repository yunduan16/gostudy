package main

import (
	"fmt"
	"math"
	"strings"
)

var (
	a  string
	b  int
	c  bool
	d  float64
	s1 = `这是第一行
第二行
	插在中间位置
最后一行
	`  //贴壁就没有tab
)

const (
	pi = 3.1415926
	e  = 2.17
)

const (
	ios     = -3
	tp      = 1
	andorid = iota      //2
	h5                  // 3
	pc      = 1 << iota //16
)

func main() {
	fmt.Println(andorid)
	fmt.Println(h5)
	fmt.Println(pc)

	var a int = 10
	fmt.Printf("%b \n", a) //用2进制表示
	fmt.Printf("%x \n", a) //用2进制表示

	var b int = 0o77       //077也可以表示8进制
	fmt.Printf("%o \n", b) //用8进制表示

	// var c int = 077
	// fmt.Printf("%o \n", c) //用8进制表示

	var d int = 0xa2
	fmt.Printf("%x \n", d) //用16进制表示 a2
	fmt.Printf("%X \n", d) //用16进制，大写 A2
	fmt.Println(d)         //十进制 162

	fmt.Println(s1)
	s1Arr := strings.Split(s1, "第")
	fmt.Println(len(s1Arr))

	traversalString()

	changeString()

	sqrtDemo()
}

func traversalString() {
	// 因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println("\n")
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func changeString() {
	s1 := "英文big猪"
	// s2 := 'a' //单引号的只能有一个字符，否则报错

	byteS1 := []byte(s1)
	byteS1[0] = 'p' //可以替换，但是中文会被破坏
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红' // runeS2[] = '好' 错误，数组长度固定
	fmt.Println(string(runeS2))
	fmt.Printf("%p \n", "你好") //字符串等没有地址，返回报错
}

func sqrtDemo() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println("c=", c)
}
