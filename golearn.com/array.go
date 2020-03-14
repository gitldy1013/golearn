package main

import "fmt"

func main() {
	fmt.Printf("%T", "hello")
	var arr = [...]int{1, 3, 5, 7, 8}
	test1(arr)
	test2(arr)
}

func test1(arr [5]int) {
	//求数组[1, 3, 5, 7, 8]所有元素的和
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Printf("数组之和为 %d \n", sum)
}

func test2(arr [5]int) {
	//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
	for index, val := range arr {
		for i := index + 1; i < len(arr); i++ {
			if (val + arr[i]) == 8 {
				fmt.Printf("(%d,%d)", index, i)
			}
		}
	}
}
