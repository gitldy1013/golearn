package main

import "fmt"

func main() {
	test_1()
	test_2()
}

func test_1() {
	var i int
	var f32 float32
	var f64 float64
	var bl bool
	var s string
	var ss string = "小桃子"
	var d int8  //int8
	var bt byte //uint8
	var ru rune //int32
	fmt.Printf("%v %T \n", i, i)
	fmt.Printf("%v %T \n", f32, f32)
	fmt.Printf("%v %T \n", f64, f64)
	fmt.Printf("%v %T \n", bl, bl)
	fmt.Printf("%v %T \n", s, s)
	fmt.Printf("%v %T \n", ss, ss)
	fmt.Printf("%v %T \n", d, d)
	fmt.Printf("%v %T \n", bt, bt)
	fmt.Printf("%v %T \n", ru, ru)
}

func test_2() {
	s := "hello小桃子"
	for i := 0; i < len(s); i++ {
		fmt.Printf("普通for-%v：类型：%T %v \n", i, s[i], s[i])
	}
	count := 0
	for index, value := range s {
		fmt.Printf("超级for-%v：类型：%T %v \n", index, value, value)
		if string(byte(value)) != string(value) {
			count++
		}
	}
	fmt.Println(count)
}
