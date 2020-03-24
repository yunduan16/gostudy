Go语言

类型
    整型 int int8(byte) int16 int32(rune) int64 uint uint8 uint16 uint32 uint64
    string
    浮点型 float32 float64 
    布尔
    数组 [...]int [2]int
    切片 []int []*int
    map map[T]T
    通道 chan T     chan<- T    <-chan T
    结构体 type X struct 
    接口  type X interface
    函数 func(参数...) (返回值...)
    方法 (接受者) func(参数...) (返回值...)


常量


概念
    包管理(mod最新)
    指针(&取地址 *地址类型取值)
    接口
    函数(匿名函数,闭包函数,函数结构)
    方法(继承)
    结构体(初始化，值|指针)
    Goroutine
    通道(接受|发送)
    切片-数组(底层实现)
    锁（读写锁）
    defer(延迟处理)
    goto(跳转代码快)
    ifelse(条件判断)
    switch(多条件判断, gothrough)
    for(循环，死循环，遍历, continue, break)
    select(io的switch, continue, break)


常用的包

    打印(fmt)
    网络(net)
    缓存io(bufio)
    系统相关(os)
    字符串(strings)
    时间(time)
    异步并发-锁相关(sync)
    二进制文件(encoding/binary)
    字节(bytes)
    错误(errors)
    页面(html)
    图片(image)
    输入输出(io)
    数学(math)
    反射(reflect)
    运行(runtime)
    


