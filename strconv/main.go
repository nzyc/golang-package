package main

import (
	"fmt"
	"strconv"
)

func main () {
	fmt.Println("hello world")

	/*
		将其他格式的转化为 int 类型
	*/
	Atoi()
	// 错误就返回 0
	ParseInt()
	// 错误就返回 0 不允许有负值 有就报错
	ParseUint()
	ParseFloat()
	ParseBool()

	/*
		将其他类型的转化为 字符 类型
	*/
	Itoa()
	FormatInt()
	FormatUint()
	FormatFloat()
	FormatBool()
}

func Atoi () {
	num, _ := strconv.Atoi("1000")
	fmt.Printf("%T ,%v,%v\n", num, num,num+1)
	fmt.Println("-------------")
}
func ParseInt () {
	num , _:= strconv.ParseInt("4e00", 16, 64)
	fmt.Println(num)	// 19968
	num , _= strconv.ParseInt("01100001", 2, 64)
	fmt.Println(num)	// 97
	num , _= strconv.ParseInt("01100001", 10, 64)
	fmt.Println(num)	// 1100001
	num , _= strconv.ParseInt("4e00", 10, 64)
	fmt.Println(num)	// 0
}
func ParseUint () {
	num , _:= strconv.ParseUint("4e00", 16, 64)
	fmt.Println(num)	// 19968
	num , _= strconv.ParseUint("01100001", 2, 64)
	fmt.Println(num)	// 97
	num , _= strconv.ParseUint("01100001", 10, 64)
	fmt.Println(num)	// 1100001
	num , _= strconv.ParseUint("4e00", 10, 64)
	fmt.Println(num)	// 0
	num , _= strconv.ParseUint("-1000", 10, 64)
	fmt.Println(num)	// 0
}
func ParseFloat () {
	pi:="3.1415926"
	num, _ := strconv.ParseFloat(pi,64)
	fmt.Printf("%T, %v %v\n",num,num,num+num)
}
func ParseBool () {
	// 1 t T True TRUE
	// 0 f F False FALSE
	flag, _:= strconv.ParseBool("0")
	fmt.Printf("%T, %v \n", flag,flag)
}

func Itoa () {
	str := strconv.Itoa(1000)
	fmt.Printf("%T,%v\n", str,str)
}
func FormatInt () {
	str := strconv.FormatInt(19968,16)
	fmt.Printf("%T,%v\n", str,str)
}
func FormatUint () {
	str := strconv.FormatUint(19968,16)
	fmt.Printf("%T,%v\n", str,str)
}
func FormatFloat () {
	PI := 3.1415926
	str := strconv.FormatFloat(PI,'f',-1,64)
	fmt.Printf("%T,%v\n",str,str)
}
func FormatBool () {
	str := strconv.FormatBool(true)
	fmt.Printf("%T,%v\n",str,str)
}