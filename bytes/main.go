package main

import (
	"bytes"
	"fmt"
)

func main () {
	fmt.Println("hello world")

	/*
		比较
	*/
	Compare()
	Equal()
	EqualFold()


	/*
		清理
	*/
	Trim()
	TrimLeft()
	TrimRight()
	TrimSpace()
	TrimPrefix()
	TrimSuffix()

	/*
		拆合
	*/
	Split()
	SplitAfter()
	SplitAfterN()
	Fields()
	FieldsFunc()
	Join()
	Repeat()

	/*
		子串
	*/
	HasPrefix()
	HasSuffix()
	Contains()
	Index()
	LastIndex()
	Count()


	/*
		替换
	*/
	Replace()
	Map()
}


// 比较俩个 byte[] 的切片是否一样
// 返回 int
// 相等返回 0 	小于号返回 -1  大于号 返回 1
func Compare () {
	one := []byte{'a','b'}
	two := []byte{'b','e'}
	num := bytes.Compare(one,two)
	fmt.Println(num)		// -1
}

// 比较俩个切片 byte[] 切片 是否相同
// 返回 bool
func Equal () {
	one := []byte{'a','b'}
	two := []byte{'b','e'}
	bool1 := bytes.Equal(one, two)
	fmt.Println(bool1)	// false
}

// 比较 俩个UTF-8的切片是否相似
// 返回 bool
func EqualFold() {}


// 清理
// 清理 两边 符合 hello的 切片
func Trim () {
	one := []byte{'a','b','c','d'}
	two := bytes.Trim(one,"hello")
	fmt.Printf("%q\n",two)
}
// 清理 左边 符合 hello的 切片
func TrimLeft () {
	one := []byte{'a','b','c','d'}
	two := bytes.TrimLeft(one,"hello")
	fmt.Printf("%q\n",two)
}
// 清理 右边 符合 hello的 切片
func TrimRight () {
	one := []byte{'a','b','c','d'}
	two := bytes.TrimRight(one,"hello")
	fmt.Printf("%q\n",two)
}
// 清理俩边的空白
func TrimSpace () {
	one := []byte{' ', 'a','b','c','d',' '}

	two := bytes.TrimSpace(one)
	fmt.Printf("%q\n", two)
}
// 清理 前缀 prefix  返回 切片
func TrimPrefix () {
	one := []byte{'a','b','c','d',' '}
	two := []byte{'a'}
	three := bytes.TrimPrefix(one,two)
	fmt.Printf("%q\n",three)
}
// 清理 后缀 suffix 返回 切片
func TrimSuffix () {
	one := []byte{'a','b','c','d',' ','e'}
	two := []byte{'e'}
	three := bytes.TrimSuffix(one, two)
	fmt.Printf("%q\n",three)
}


// 拆合
func Split () {
	one := []byte{'a','b','c','a','d',' ','e'}
	two := []byte{'a'}
	three := bytes.Split(one,two)
	fmt.Printf("%q\n", three)	// ["" "bc" "d e"]
}

func SplitAfter () {
	one := []byte{'a','b','c','a','d',' ','e'}
	two := []byte{'a'}
	three := bytes.SplitAfter(one,two)
	fmt.Printf("%q\n", three)	// ["a" "bca" "d e"]
}

func SplitAfterN () {
	one := []byte{'a','b','c','a','d',' ','e'}
	two := []byte{'a'}
	three := bytes.SplitAfterN(one,two,3)
	fmt.Printf("%q\n", three)	// ["a" "bca" "d e"]
}

func Fields () {
	one := []byte{'a','b','c','a','d',' ','e'}
	two := bytes.Fields(one)
	fmt.Printf("%q\n",two)		// ["abcad" "e"]
}

func FieldsFunc () {
	one := []byte{'a','b','c','a','d',' ','e'}
	two := bytes.FieldsFunc(one, func(r rune) bool {
		if r == 'a' {
			return true
		}
		return false
	})
	fmt.Printf("%q\n",two)		// ["bc" "d e"]
}

func Join () {
	one := [][]byte{[]byte{'a', 'b', 'c', 'd', 'e'}}
	two := bytes.Join(one,[]byte{'-'})
	fmt.Printf("%q\n",two)		// "abcde"
}

func Repeat (){
	one:=[]byte{'a'}
	two := bytes.Repeat(one,5)
	fmt.Printf("%q\n",two)		// "aaaaa"
}


func HasPrefix () {
	one := []byte{'a', 'b', 'c', 'd', 'e'}
	two := bytes.HasPrefix(one,[]byte{'a'})
	fmt.Println(two)					// true
}

func HasSuffix () {
	one := []byte{'a', 'b', 'c', 'd', 'e'}
	two := bytes.HasSuffix(one,[]byte{'e'})
	fmt.Println(two)					// true
}

func Contains (){
	one := []byte{'a', 'b', 'c', 'd', 'e'}
	two := bytes.Contains(one,[]byte{'c'})
	fmt.Println(two)					// true
}

func Index () {
	one := []byte{'a', 'b', 'c', 'd', 'e'}
	two := bytes.Index(one,[]byte{'c'})
	fmt.Println(two)					// 2 找到返回下标 没找到 返回 -1
}

func LastIndex () {
	one := []byte{'a', 'b', 'c', 'c', 'd', 'e'}
	two := bytes.LastIndex(one,[]byte{'c'})
	fmt.Println(two)					// 3 找到返回下标 没找到 返回 -1
}

func Count () {
	one := []byte{'a', 'b', 'c', 'c','c','c', 'd', 'e'}
	two := bytes.Count(one,[]byte{'c'})
	fmt.Println(two)					// 4 返回出现次数
}

func Replace () {
	one:= []byte{'a','a','a','c','c','c','c','d','q','r'}
	two := bytes.Replace(one,[]byte{'a'},[]byte{'w'},1)
	fmt.Printf("%q\n",two)		// "waaccccdqr"
}

func Map () {
	one:= []byte{'a','a','a','c','c','c','c','d','q','r'}
	two := bytes.Map(func(r rune) rune {
		if r == 'a' {
			return 's'
		}
		return r
	},one)
	fmt.Printf("%q\n",two)		// "sssccccdqr"
}