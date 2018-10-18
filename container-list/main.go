package main

import (
	"container/list"
	"fmt"
)

func main () {
	/*
		声明方式
			1、var 变量 list.list
			2、变量名 := list.New()
			推荐使用 New()
		list 是值类型 不是 应用类型
		list 是值类型 不过采用 New方法声明的是一个指针。所以在拷贝和参数传递的时候具备了引用类型的特征
	*/

	// 声明容器类
	listOne := list.New()

	// 返回链表第一个元素
	listOne.Front()

	// 返回链表的最后一个元素
	listOne.Back()

	// 将新值添加到链表的第一个位置，范围新元素
	listOne.PushFront("1")
	listOne.PushFront("2")
	listOne.PushFront("3")

	// 将新值添加到链表的最后一个位置，范围新元素
	listOne.PushBack("4")
	listOne.PushBack("5")
	listOne.PushBack("6")

	// 返回链表的长度
	fmt.Println(listOne.Len())

	// 正序遍历链表
	for e := listOne.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v \t", e.Value)
	}

	// 清空链表
	listOne.Init()

	// 打印长度
	fmt.Println("")
	fmt.Println(listOne.Len())

	// 声明
	listTwo := list.New()
	listTwo.PushFront("one")
	listTwo.PushFront("two")
	listTwo.PushBack("three")
	listTwo.PushBack("four")

	// 给链表最前方添加链表
	listOne.PushFrontList(listTwo)

	// 给链链表最后方添加链表
	listOne.PushBack("-")
	listOne.PushBackList(listTwo)

	// 逆序排列链表
	for e := listOne.Back(); e != nil; e = e.Prev() {
		fmt.Printf("%v \t", e.Value)
	}

	// 声明
	listThree := list.New()
	one := listThree.PushFront("one")
	listThree.PushFront("two")
	element := listThree.PushBack("three")
	listThree.PushBack("four")
	five := listThree.PushBack("five")

	// 在链表某个元素前插入
	listThree.InsertBefore("HAHA",element)

	// 在链表某个元素后插入
	listThree.InsertAfter("XIXI",element)

	// 正序遍历元素
	fmt.Println("")
	for e := listThree.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v \t", e.Value)
	}

	// 将元素移动到链表的第一个元素 如果元素不是链表的元素 则不会被修改
	listThree.MoveToFront(five)

	// 将元素移动到链表的最后一个元素 如果元素不是链表的元素 则不会被修改
	listThree.MoveToBack(one)

	// 正序遍历元素
	fmt.Println("")
	for e := listThree.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v \t", e.Value)
	}

	// 将元素移动到链表 mark 的前面 如果元素和mark不是链表的元素 或者 e==mark 则不会被修改
	listThree.MoveBefore(one, element)

	// 将元素移动到链表 mark 的后面 如果元素和mark不是链表的元素 或者 e==mark 则不会被修改
	listThree.MoveAfter(five, element)

	// 正序遍历元素
	fmt.Println("")
	for e := listThree.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v \t", e.Value)
	}

	// 删除元素
	listThree.Remove(element)

	// 正序遍历元素
	fmt.Println("")
	for e := listThree.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v \t", e.Value)
	}



}
