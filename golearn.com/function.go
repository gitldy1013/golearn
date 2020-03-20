package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(intSum_1(999,999))
	//fmt.Println(intSum_2(999,999))
	//fmt.Println(intSum_3(999,999))
	//fmt.Println(intSum_4(999,999,998))
	//fmt.Println(intSum_4())//可变参数也可以不传
	//fmt.Println(intSum_5(998,999,1000))
	//fmt.Println(intSum_5(888))//固定参数和可变参数同时出现时固定参数部分必须传入
	//fmt.Printf("%#v\n",someFunc(""))
	//fmt.Println(num)
	//testGlobalVar()
	//testLocalVar(1,2)
	//funcType()
	//fmt.Println(calc(1,2,intSum_1)) //函数作为参数

	//var f = adder_1()
	//fmt.Println(f(10)) //10
	//fmt.Println(f(20)) //30
	//fmt.Println(f(30)) //60
	//f1 := adder_1()
	//fmt.Println(f1(40)) //40 (此时不再累加)
	//fmt.Println(f1(50)) //90
	//注意: 闭包指的是一个函数和与其相关的引用环境组合而成的实体。
	//变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。 在f的生命周期内，变量x也一直有效。

	//var f2 = adder_2(10)
	//fmt.Println(f2(10)) //20
	//fmt.Println(f2(20)) //40
	//fmt.Println(f2(30)) //70

	//f3 := adder_2(20)
	//fmt.Println(f3(40)) //60
	//fmt.Println(f3(50)) //110

	//jpgFunc := makeSuffixFunc(".jpg")
	//txtFunc := makeSuffixFunc(".txt")
	//fmt.Println(jpgFunc("test")) //test.jpg
	//fmt.Println(txtFunc("test")) //test.txt

	//fc1, fc2 := calcPro(10)
	//fmt.Println(fc1(1), fc2(2)) //11 9
	//fmt.Println(fc1(3), fc2(4)) //12 8
	//fmt.Println(fc1(5), fc2(6)) //13 7

	//deferTest()
	deferDemo()
	defer_case()
	testFun()
}

func intSum_1(x int, y int) int {
	return x + y
}

func intSub_1(x int, y int) int {
	return x - y
}

func intSum_2(x, y int) int {
	return x + y
}

//返回值定义参数时需要用小括号包裹
func intSub_2(x, y int) (ret int) {
	ret = x - y
	return //此时可以不写返回值 但return关键字不可省略
}

func intSum_3(x, y int) (int, int) {
	return x + y, x
}

//可变参数
func intSum_4(x ...int) (int, int) {
	sum := 0
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum, len(x)
}

//可变参数必须放在最后 参数和返回值不能冲突定义 返回值也支持类型简写
func intSum_5(y int, x ...int) (sum, yr int) {
	//sum := 0
	for _, v := range x {
		sum += v
	}
	yr = y
	//return sum, yr
	return
}

//返回切片
func someFunc(x string) []int {
	if x == "" {
		return nil //可以返回nil值
	} else {
		return []int{len(x)}
	}
}

//定义全局变量
var num int64 = 10

func testGlobalVar() {
	//定义局部变量
	var numx int = 100
	fmt.Printf("numx=%d\n", numx)
	fmt.Printf("num=%d\n", num)
	num := 200
	//优先访问局部变量
	fmt.Println(num)
}

func testLocalVar(x, y int) {
	//fmt.Printf("numx=%d\n",numx) //这里无法获取numx
	fmt.Println(x, y) //函数的参数也是只在本函数中生效
	if x > 0 {
		z := 100 //变量z只在if语句块生效
		fmt.Println(z)
	}
	//fmt.Println(z)//此处无法使用变量z
	for i := 0; i < 10; i++ {
		fmt.Println(i) //变量i只在当前for语句块中生效
	}
	//fmt.Println(i) //此处无法使用变量i
}

func funcType() {
	//函数类型
	type calculation func(int, int) int // 声明一个calculation类型的变量c
	var c calculation
	c = intSum_1                     // 把intSum_1赋值给c
	c(1, 2)                          // 像调用intSum_1一样调用c
	f := intSum_1                    // 将函数intSum_1赋值给变量f
	f(1, 2)                          // 像调用intSum_1一样调用f
	fmt.Printf("type of c:%T \n", c) //type of c:main.calculation
	fmt.Printf("type of c:%T \n", f) //type of c:func(int, int) int
}

//函数作为参数
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

//函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return intSum_1, nil
	case "-":
		return intSub_1, nil
	default:
		err := errors.New("无法识别操作符")
		return nil, err
	}
}

//匿名函数: 定义在函数内部
func testInnerFunc() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	// 将匿名函数保存到变量
	add(1, 2)

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}

//注意: 闭包=函数+引用环境
func adder_1() func(int) int {
	var x int
	fmt.Println("x:", x)
	return func(y int) int {
		fmt.Println("y:", y)
		x += y
		fmt.Println("x+y:", x)
		return x
	}
}

//闭包进阶1
func adder_2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

//闭吧进阶2
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

//闭包进阶3
func calcPro(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func deferTest() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
	//总结:
	// return语句执行底层实现: 返回值=x -> RET 指令
	// defer语句执行的时机: 返回值=x —> 运行defer -> RET指令
}

//defer案例 ==start==
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func deferDemo() {
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5
}

//defer案例 ==end==

//defer 面试题 ==start==
func calc_defer(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func defer_case() {
	x := 1
	y := 2
	defer calc_defer("AA", x, calc_defer("A", x, y))
	x = 10
	defer calc_defer("BB", x, calc_defer("B", x, y))
	y = 20
}

//defer 面试题 ==end==

//panic/recover
func funcAp() {
	fmt.Println("func A")
}

func funcBp() {
	panic("panic in B")
}

func funcCp() {
	fmt.Println("func C")
}
func panicDemo() {
	funcAp()
	funcBp()
	funcCp()
}

//程序运行期间funcB中引发了panic导致程序崩溃，异常退出了。这个时候我们就可以通过recover将程序恢复回来，继续往后执行。
/* 输出:
func A
panic: panic in B

goroutine 1 [running]:
main.funcB(...)
.../code/func/main.go:12
main.main()
.../code/func/main.go:20 +0x98
*/

func funcAr() {
	fmt.Println("func A")
}

func funcBr() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcCr() {
	fmt.Println("func C")
}
func recoverDemo() {
	funcAr()
	funcBr()
	funcCr()
}

/*
注意:
	1.recover()必须搭配defer使用。
	2.defer一定要在可能引发panic的语句之前定义。
*/
//panic/recover

/*
分金币:
	  你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
	  分配规则如下:
			  a. 名字中每包含1个'e'或'E'分1枚金币
			  b. 名字中每包含1个'i'或'I'分2枚金币
			  c. 名字中每包含1个'o'或'O'分3枚金币
			  d: 名字中每包含1个'u'或'U'分4枚金币
	  写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
	  程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func testFun() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

func dispatchCoin() (res int) {
	//standard := make(map[string]int,4)
	standard := map[string]int{"e": 1, "i": 2, "o": 3, "u": 4}
	for i := 0; i < len(users); i++ {
		distribution[users[i]] = 0
		lower := strings.ToLower(users[i])
		for index, value := range standard {
			if strings.Contains(lower, index) {
				count := strings.Count(lower, index)
				distribution[users[i]] += value * count
				coins -= value * count
				fmt.Println(users[i]+" "+index+" 当前金币:", coins)
			}
		}
	}
	fmt.Printf("%#v\n", distribution)
	return coins
}

/*
python的实现

namestr='Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth'
namelist=namestr.split(',')
coindir={}

for i in namelist:
    coindir[i]=0
# print(coindir)
new_coindir=coindir


def addcoin(namelist):
    zimustr = 'e,i,o,u'
    zimulist=zimustr.split(',')
    ZIMUstr = 'E,I,O,U'
    ZIMUlist = ZIMUstr.split(',')
    for i in namelist:
        for j in i :
            # print(j)
            if j ==zimulist[0] or j ==ZIMUlist[0] :
                new_coindir[i]+=1
            elif j ==zimulist[1] or j ==ZIMUlist[1] :
                new_coindir[i]+=2
            elif j == zimulist[2] or j == ZIMUlist[2]:
                new_coindir[i] += 3
            elif j == zimulist[3] or j == ZIMUlist[3]:
                new_coindir[i] += 4


def leftcoin(new_coindir):
    all_coin=50
    now_coin=0
    for value in new_coindir:
        now_coin=all_coin-new_coindir[value]
        all_coin=now_coin
        print(now_coin)
    return now_coin


addcoin(namelist)
print(new_coindir)
left_coin=leftcoin(new_coindir)
print('剩余的金币数是：'+ str(left_coin))
*/
