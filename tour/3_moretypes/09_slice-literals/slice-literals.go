/* 切片文法
切片文法类似于没有长度的数组文法。

这是一个数组文法：

[3]bool{true, true, false}
下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：

[]bool{true, true, false} */


package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)	// [2 3 5 7 11 13]

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)	// [true false true true false true]

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)	// [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
}
