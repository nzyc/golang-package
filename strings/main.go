package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main () {
	fmt.Println("hello world")

	s := "我爱go语言"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i,ch := range s {
		fmt.Printf("%d:%c \n",i,ch)
	}
	for i,ch := range []byte(s){
		fmt.Printf("%d:%x \n",i,ch)
	}
	for i,ch := range []rune(s) {
		fmt.Printf("%d:%c \n",i,ch)
	}


	/*
		 比较
	*/
	// 比较 a 和 b 的大小
	Compare()
	// 比较俩个 utf8 字符串 是否相等  忽略大小写
	EqualFold()
	// 将字符串重复 count 次后返回
	Repeat()
	// 将 str 字符串中 old 替换为 new n < 0 替换所有的old
	Replace()
	// 将 所有字符串 连接成一个字符串 并以 sep 作为分割符
	Join()


	/*
		转换
	*/
	// 每个单词首字母大写
	Title()
	// 全小写
	ToTitle()
	// 转大写
	ToUpper()
	// 转小写
	ToLower()


	/*
		清理
	*/
	// 清除 str 符合 h 的字符
	Trim()
	// 清除 str 符合 f 函数 的字符
	TrimFunc()
	// 清除 str 符合 h 左边 的字符
	TrimLeft()
	// 清除 str 符合 f 函数 左边 的字符
	TrimLeftFunc()
	// 清除 str 符合 h 右边 的字符
	TrimRight()
	// 清除 str 符合 f 函数 右边 的字符
	TrimRightFunc()
	// 清除俩边的空格
	TrimSpace()
	// 清除 符合 字符串的 前缀
	TrimPrefix()
	// 清除 符合 字符串的 后缀
	TrimSuffix()


	/*
	 	检索
	*/
	Contains()
	ContainsAny()
	ContainsRune()
	Count()
	HasPrefix()
	HasSuffix()
	Index()
	IndexAny()
	IndexByte()
	IndexRune()
	IndexFunc()
	LastIndex()
	LastIndexAny()
	LastIndexByte()
	LastIndexFunc()


	/*
		分割
	*/
	// 以空白字符分割 返回一个切片
	Fields()
	// 将字符串以 f(r) == true 分割 返回一个切片
	FieldsFunc()
	// 将字符串以 sep 分割 分割字符串后去掉 sep
	Split()
	// 将字符串以 sep 分割 分割字符串后去掉 sep N决定返回的切片数
	SplitN()
	// 将字符串以 sep 分割 分割字符串后附加 sep
	SplitAfter()
	// 将字符串以 sep 分割 分割字符串后附加 sep N决定返回的切片数
	SplitAfterN()
}

/*
	比较
*/
func Compare () {
	one := "hello"
	two := "world"
	three := strings.Compare(one,two)
	fmt.Println(three)				// 相等返回 0 	不等返回 -1  大于号 返回 1
}
func EqualFold () {
	one := "我是"
	two := "傻逼"
	three := strings.EqualFold(one,two)
	fmt.Println(three)				// false
}
func Repeat () {
	fmt.Println(strings.Repeat("simawanyi", 3))								// simawanyisimawanyisimawanyi
}
func Replace () {
	fmt.Println(strings.Replace("hello world", "world", "go", -1))		// hello go
}
func Join () {
	fmt.Println(strings.Join([]string{"asd","asd","asd","asd","asd"}, "~"))		// asd~asd~asd~asd~asd
}
/*
	转换
*/
func Title(){
	fmt.Println(strings.Title("hello world"))
}
func ToTitle () {
	one := "hellO world"
	two:= strings.ToTitle(one)
	fmt.Println(two)				// HELLO WORLD
}
func ToUpper () {
	one := "hello"
	two:= strings.ToUpper(one)
	fmt.Println(two)				// HELLO
}
func ToLower () {
	one := "HELLO"
	two:= strings.ToLower(one)
	fmt.Println(two)				// hello
}
/*
	清除
*/
func Trim () {
	one := "hello world 1h"
	two := "h"
	three := strings.Trim(one,two)
	fmt.Println(three)				// ello world 1
}
func TrimFunc() {
	f:= func(c rune) bool {
		if c == 'h' {
			return true
		}else {
			return false
		}
	}
	fmt.Println(strings.TrimFunc("hello world", f))
}
func TrimLeft () {
	one := "hello world 1h"
	two := "h"
	three := strings.TrimLeft(one,two)
	fmt.Println(three)				// eello world 1h
}
func TrimLeftFunc () {}
func TrimRight () {
	one := "hello world 1h"
	two := "h"
	three := strings.TrimRight(one,two)
	fmt.Println(three)				// hello world 1
}
func TrimRightFunc () {}
func TrimSpace () {
	one := " hello world "
	two := strings.TrimSpace(one)
	fmt.Println(two)
}
func TrimPrefix()  {
	fmt.Println(strings.TrimPrefix("hello world", "he"))	// llo world
}
func TrimSuffix () {
	fmt.Println(strings.TrimSuffix("hello world", "ld"))	// hello wor
}
/*
	 检索
*/
func Contains () {
	fmt.Println(strings.Contains("seafood","foo"))	// true
	fmt.Println(strings.Contains("seafood","boo"))	// false
	fmt.Println(strings.Contains("seafood",""))		// true
	fmt.Println(strings.Contains("",""))				// true
	fmt.Println(strings.Contains("steven王","王"))		// true
}
func ContainsAny() {
	fmt.Println(strings.ContainsAny("seafood","foo"))	// true
	fmt.Println(strings.ContainsAny("seafood","f%s"))	// true
	fmt.Println(strings.ContainsAny("",""))			// false
	fmt.Println(strings.ContainsAny("seafood",""))		// false
}
func ContainsRune() {
	fmt.Println(strings.ContainsRune("定定定",'定'))		// true
	fmt.Println(strings.ContainsRune("一定定",19968))		// true
}
func Count () {
	fmt.Println(strings.Count("hello", "l"))			// 2
	fmt.Println(strings.Count("hello", ""))			// 6
}
func HasPrefix() {
	fmt.Println(strings.HasPrefix("hello", "he"))		// true
}
func HasSuffix() {
	fmt.Println(strings.HasSuffix("hello", "lo"))		// true
}
func Index () {
	fmt.Println(strings.Index("hello", "lo"))			// 3
}
func IndexAny () {
	fmt.Println(strings.IndexAny("hello", "基地h"))	// 0
}
func IndexByte() {
	fmt.Println(strings.IndexByte("hello基", 'a'))		// -1
}
func IndexRune () {
	fmt.Println(strings.IndexRune("hello基", '基'))		// 5
}
func IndexFunc () {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han,c)
	}
	fmt.Println(strings.IndexFunc("hello,基地", f))			// 6
}
func LastIndex(){}
func LastIndexAny(){}
func LastIndexByte(){}
func LastIndexFunc(){}
/*
	分割
*/
func Fields () {
	fmt.Println(strings.Fields("hello"))	// [hello]
}
func FieldsFunc () {}
func Split () {
	fmt.Println(strings.Split("hello world","o"))				// [hell  w rld]
}
func SplitN () {
	fmt.Println(strings.SplitN("hello world","o", -1))		// [hell  w rld]
}
func SplitAfter() {
	fmt.Println(strings.SplitAfter("hello world","o"))			// [hello  wo rld]
}
func SplitAfterN() {
	fmt.Println(strings.SplitAfterN("hello world","o",-1))	// [hello  wo rld]
}