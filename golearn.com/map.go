package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func main() {
	//base()
	//iterator()
	//iteratorBySort()
	//deleteFunc()
	//mapSlice()
	//sliceMap()
	//fmt.Println(testMap_1("how do you do"))
	testMap_2()
}

func base() {
	//基础用法
	var scoreMap map[string]int
	fmt.Println(scoreMap == nil)
	fmt.Printf("%#v \n", scoreMap) //友好打印
	scoreMap = make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("%v \n", scoreMap)
	fmt.Printf("%#v \n", scoreMap)
	fmt.Printf("type of a: %T \n", scoreMap)

	userInfo := map[string]string{
		"username": "小桃子",
		"password": "01234567", //和java不同换行后无论是不是最后一个值都要加行逗号 或者像下面一样直接写到一行
	}
	//userInfo := map[string]string{"username": "小桃子", "password": "01234567"}
	fmt.Println(userInfo)

	v, ok := scoreMap["李四"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}

func iterator() {
	//遍历迭代
	scoreMap := make(map[string]int)
	scoreMap["娜扎"] = 60
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["阿强"] = 100
	scoreMap["赵六"] = 100
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	fmt.Println(scoreMap)
	for x := range scoreMap {
		fmt.Println(x) //只取一个时 x是key的值
		fmt.Println(scoreMap[x])
	}
	for _, v := range scoreMap {
		fmt.Println(v) //只遍历value
	}
	//注意: 遍历map时的元素顺序与添加键值对的顺序无关
	fmt.Println(scoreMap)

}

func iteratorBySort() {
	//指定顺序遍历
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for k := range scoreMap {
		keys = append(keys, k)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

func deleteFunc() {
	//var scoreMap map[string]int
	scoreMap := map[string]int{} //此处初始化为空map 不要理解为和上面一样未初始化
	scoreMap["娜扎"] = 60
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["阿强"] = 100
	scoreMap["赵六"] = 100
	fmt.Println(len(scoreMap))
	delete(scoreMap, "张三")
	for k, v := range scoreMap {
		//迭代过程中也可以删除元素
		//delete(scoreMap,k)
		fmt.Println(k, v)
	}
	fmt.Println(len(scoreMap))
}

func mapSlice() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index: %d value:%#v \n", index, value)
	}
	fmt.Println("after init")
	//对切片中的map元素进行初始化 这里值得注意的是切片中的map必须初始化才能进行赋值操作
	for index, value := range mapSlice {
		fmt.Printf("%T %#v %v %+v ", index, index, index, index)
		mapSlice[index] = make(map[string]string, 10)
		//注意: 字符串和int类型转换不能使用 string(index)
		mapSlice[index]["name"] = "小王子" + fmt.Sprintf("%d", index)
		mapSlice[index]["password"] = "123456" + fmt.Sprintf("%d", index)
		mapSlice[index]["address"] = "辽宁" + fmt.Sprintf("%d", index)
		fmt.Printf("index: %d value:%#v \n", index, value) //这里打印的仍为nil 因为value被赋值之前切片里的map还没有被初始化
	}
	for index, value := range mapSlice {
		fmt.Printf("index: %d value:%v \n", index, value)
		fmt.Printf("index: %d value:%#v \n", index, value) //啧啧啧
	}
}

func sliceMap() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Printf("%#v\n", sliceMap)
	fmt.Println("after init")
	sliceMap["中国"] = []string{"加油", "必胜"}
	//sliceMap["美国"][0] = "得瑟"//此处因为没有初始化 会报错
	key := "日本"
	//key = "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 2) //初始化两个空串 这里的一个细节是初始化时切片长度不能是0 0相当于没有初始化
		fmt.Printf("%#v\n", value[0])
	} else {
		for i, v := range value {
			fmt.Println(i, v)
		}
	}
	value = append(value, "灭绝", "作死")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

func testMap_1(s string) map[string]int {
	sts := strings.Split(s, " ")
	var res = make(map[string]int, len(sts)*2)
	for _, str := range sts {
		if res[str] != 0 {
			res[str] += 1
		} else {
			res[str] = 1
		}
	}
	return res
}

func testMap_2() {
	type info struct {
		name string
		id   int
	}
	v := info{"Nan", 33}
	fmt.Printf("%v\n", v)
	fmt.Printf("%+v\n", v)
	//	{Nan 33}
	//	{name:Nan id:33}

	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s) //{1 2 3} 切片s中追加一个元素
	m["太阳"] = s            //此时m中的切片已经被赋值为{1,2,3} cap为2的倍数 此时为4
	//fmt.Println(len(s))
	//fmt.Println(cap(s))
	//fmt.Println(len(m["太阳"]))
	//fmt.Println(cap(m["太阳"]))
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)       //{1 3} 切片是中索引为1的元素被删除 影响了原来存储数字2变为3
	fmt.Printf("%+v\n", m["太阳"]) //{1,3,3} 切片改变的内存数据

	//不信我们就再试试
	println(s)
	s = append(s, s...)
	println(s)
	s[2] = 10
	s[3] = 9
	fmt.Printf("试试：%+v\n", s)
	fmt.Printf("试试：%+v\n", m["太阳"]) //想想为啥子这里不是{1,3,10,9}
	//fmt.Println(len(s))
	//fmt.Println(cap(s))
	//fmt.Println(len(m["太阳"]))
	//fmt.Println(cap(m["太阳"]))
	println(m["太阳"])
	m["太阳"] = append(m["太阳"][:1], m["太阳"][1:]...) //只要不扩容
	println(m["太阳"])
	s[0] = 888
	m["太阳"][0] = 777
	fmt.Printf("试试这个：%+v\n", s)
	fmt.Printf("试试这个：%+v\n", m["太阳"])
	m["太阳"][0] = 333
	s[0] = 444
	//s = append(s[:0], s[4:]...)//清空切片数据试试
	fmt.Printf("试试那个：%+v\n", s)
	fmt.Printf("试试那个：%+v\n", m["太阳"])

	//扩展测试切片再次赋值
	s = []int{4, 5, 6}
	//我们再看看m["太阳"]的值
	fmt.Printf("再试试：%+v\n", s)
	fmt.Printf("再试试：%+v\n", m["太阳"]) //此时切片已经指向一块新的内存地址 所以不会影响原来的数据
}
