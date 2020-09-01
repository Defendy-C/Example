package main

import "fmt"

func main() {
	t := NewTree()
	arr := []int{3, 2, 1, 4, 5, 6, 7, 10, 9, 8}
	for i := 0; i < len(arr); i++ {
		t.Add(arr[i])
	}
	fmt.Println("del")
	t.Del(1)
	t.Del(3)
	fmt.Println("end")
}
