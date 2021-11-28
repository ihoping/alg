package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := []int{4,6,8,5,9}
	//fmt.Println(sort.QuickSort(s))
	//lc := link.Constructor(2)
	//lc.Put(1, 1)
	//lc.Put(2, 2)
	//lc.Print()
	//
	//fmt.Println()
	//
	//lc.Put(3, 3)
	//lc.Put(4, 4)
	//lc.Print()
	//
	//fmt.Println("------")
	//
	//fmt.Println(lc.Get(1))
	//fmt.Println("------")
	//lc.Print()
	//
	//fmt.Println()
	//
	//lc.Put(3, 3)
	//lc.Print()
	//
	//fmt.Println()
	//
	//lc.Get(2)
	//lc.Print()
	//
	//fmt.Println()
	//
	//lc.Put(4, 4)
	//lc.Print()
	//
	//fmt.Println()
	//lc.Get(1)
	//lc.Print()
	//
	//fmt.Println()
	//lc.Get(3)
	//lc.Print()
	//
	//fmt.Println()
	//lc.Get(4)
	//lc.Print()

	//fmt.Println(lc.Get(1))
	s1 := "ddddsdsd"
	var build strings.Builder

	for i := len(s1) - 1; i >= 0; i-- {
		build.WriteString(string(s1[i]))
	}
	fmt.Println(build.String())
}

func maxsumofSubarray(arr []int) int {
	// write code here
	if arr == nil {
		return 0
	}

	length := len(arr)
	if length == 1 {
		return arr[0]
	}

	m := make(map[int]int)
	j := length - 1
	m[j] = arr[j]
	j--

	for j >= 0 {
		m[j] = max(arr[j]+m[j+1], arr[j])
		j--
	}

	fmt.Println(m)
	return k

}

var k int

func max(a, b int) int {
	if a > b {
		if a > k {
			k = a
		}
		return a
	}
	return b
}
