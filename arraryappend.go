package main

import (
	"fmt"
)



func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	// s1 := arr1[1:2] // 切片容量足够时不会分配新的数组，而是直接修改原数组。
	s1 := arr1[4:5] // 切片 s1 不再引用原数组 arr1 的底层数组，而是引用新的底层数组。
	s1 = append(s1, 7) // 追加一个元素，根据切片的容量，是否会分配一个新的底层数组。

	fmt.Println("arr1", arr1)
	fmt.Println("s1", s1)
	s1[0] = 6
	fmt.Println("arr1", arr1)
	i := make([]int, 4, 5)
	fmt.Println(i)
}