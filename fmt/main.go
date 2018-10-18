package main

import "fmt"

func main () {
	fmt.Println("hello world")

	/*
		// 标准输出 不换行
		fmt.Print(a)
		// 标准输出 换行
		fmt.Println(a)
		// 写入到站位符当中
		fmt.Printf("%d",a)

		// 与上面三个相同 只不过是输出到 w 中去
		fmt.Fprint(w,a)
		fmt.Fprintln(w,a)
		fmt.Fprintf(w,a)

		// 与上面相同 只不过以字符串的形式输出
		fmt.Sprint()
		fmt.Sprintln()
		fmt.Sprintf()

	*/

	/*
		示例
	*/

	arr := []int{1,2,3,4,5,6,7,8,9}


	/*
		fmt.Print
	*/
	fmt.Print(arr)		// [1 2 3 4 5 6 7 8 9]	不换行


	/*
		fmtPrintln
	*/
	fmt.Println(arr)	// [1 2 3 4 5 6 7 8 9]	换行


	/*
		fmt.Printf
	*/
	fmt.Printf("ab %d %d %d cd\n", 1, 2, 3)


	/*
		fmt.Errorf
		错误输出
	*/
	if err := percent(30, 70, 90, 160); err != nil {
		fmt.Println(err)
	}


}

func percent(i ...int) error {
	for _, n := range i {
		if n > 100 {
			return fmt.Errorf("数值 %d 超出范围（100）", n)
		}
		fmt.Print(n, "%\n")
	}
	return nil
}