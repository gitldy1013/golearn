package main

import (
	"fmt"
	"sort"
)

func main() {
	//testSlice_1()
	//testSlice_2()
	//testMake_1()
	//testMake_2()
	//testForSlice()
	//sliceFunction()
	//testWork_1()
	testWork_2()
}

func testSlice_1() {
	//切片再切片
	a := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	fmt.Printf("a:%v type:%T len:%d  cap:%d\n", a, a, len(a), cap(a))
	b := a[1:3]
	fmt.Printf("b:%v type:%T len:%d  cap:%d\n", b, b, len(b), cap(b))
	c := b[1:5]
	fmt.Printf("c:%v type:%T len:%d  cap:%d\n", c, c, len(c), cap(c))
}

func testSlice_2() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	//var d = []bool{false, true} //声明一个布尔切片并初始化
	//fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
}

func testMake_1() {
	//make([]T, size, cap)
	//T:切片的元素类型
	//size:切片中元素的数量
	//cap:切片的容量
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10

	//切片不能直接比较
	var s1 []int
	fmt.Println(len(s1), cap(s1), s1 == nil)
	s2 := []int{}
	fmt.Println(len(s2), cap(s2), s2 == nil)
	s3 := make([]int, 0)
	fmt.Println(len(s3), cap(s3), s3 == nil)
}

func testMake_2() {
	//切片的赋值拷贝
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println(s1) //[100 0 0]
	fmt.Println(s2) //[100 0 0]
}

func testForSlice() {
	//切片遍历
	s := []int{1, 3, 5}

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}

func sliceFunction() {
	//append()方法为切片添加元素
	var s []int
	s = append(s, 1)       // [1]
	s = append(s, 2, 3, 4) // [1 2 3 4]
	ss := []int{5, 6, 7}
	s = append(s, ss...) // [1 2 3 4 5 6 7]
	var citySlice []string
	// 追加一个元素
	citySlice = append(citySlice, "北京")
	// 追加多个元素
	citySlice = append(citySlice, "上海", "广州", "深圳")
	// 追加切片
	a := []string{"成都", "重庆"}
	citySlice = append(citySlice, a...)
	fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]

	//注意：通过var声明的零值切片可以在append()函数直接使用，无需初始化
	s1 := []int{} // 没有必要初始化
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)
	var s2 = make([]int, 0) // 没有必要初始化
	s2 = append(s2, 1, 2, 3)
	fmt.Println(s2)

	//扩容：append()添加元素和切片扩容
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
	/**
	结论：
	1.append()函数将元素追加到切片的最后并返回该切片。
	2.切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。
	*/

	//切片copy
	ac := []int{1, 2, 3, 4, 5, 5}
	bc := make([]int, 5, 5)
	cc := bc
	copy(bc, ac)
	bc[0] = 9
	fmt.Println(ac)
	fmt.Println(bc)
	fmt.Println(cc)

	//切片删除元素
	ac = append(ac[0:2], ac[3:]...)
	fmt.Println(ac)
	/**
	总结：要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
	*/
}

func testWork_1() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(fmt.Sprintf("%T", a[0]))
	fmt.Printf("%s", a)
	//前面五个空{"","","","","","0","1","2","3","4","5","6","7","8","9"}
}

func testWork_2() {
	a := [...]int{3, 7, 8, 9, 1}
	//b := a[:]
	//sort.Ints(b)
	//fmt.Println(a)
	//fmt.Printf("a type is %T \n",a)
	//fmt.Println(b)
	//fmt.Printf("b type is %T \n",b)
	//这里需要注意 a数组因为已经通过切片排序 切片所指向的底层数据就是被切的数组 所以a数组和b切片都是排序后的输出结果 所以亦可以简化为
	sort.Ints(a[:])
	fmt.Print(a)
}
