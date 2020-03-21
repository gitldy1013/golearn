package main

import "fmt"

func main() {
	//pointer_1()
	//pointer_2()
	//pointer_3()
	//pointer_4()
	//pointer_4pro()
	//pointer_5()
	pointer_5pro()

}

//指针地址与指针类型
func pointer_1() {
	a := 10
	b := &a
	fmt.Printf("a:%v \t\t\t a:%d \t\t\t\t a-type:%T \t a-ptr:%p \t\t &a-ptr:%p\n", a, a, a, a, &a)
	fmt.Printf("b:%v \t b:%d \t b-type:%T \t b-ptr:%p \t &b-ptr:%p\n", b, b, b, b, &b)
	c := &b
	d := &c
	fmt.Printf("c:%v \t c:%d \t c-type:%T \t c-ptr:%p \t &c-ptr:%p\n", c, c, c, c, &c)
	fmt.Printf("d:%v \t d:%d \t d-type:%T \t d-ptr:%p \t &d-ptr:%p\n", d, d, d, d, &d)
}

//指针取值
func pointer_2() {
	a := 10
	b := &a //取变量a的地址 将指针的值保存在b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
	/*
		总结： 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
		变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
			>* 对变量进行取地址（&）操作，可以获得这个变量的指针变量。
			>* 指针变量的值是指针地址。
			>* 对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
	*/
}

//指针传值
func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func pointer_3() {
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)    // *int类型变量&a
	fmt.Println(a) // 100
}

//new和make
func pointer_4() {
	var a *int
	*a = 100 //这句还报错 注释掉后输出结果为<nil> 可以看出a已经是nil了 所以*a必然报空指针
	fmt.Println(a)
	fmt.Printf("%#v %#v\n", a)
	/*
		此时因为还没有在内存中分配内存地址 无法直接声明*int类型变量进行赋值
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x109cecf]
	*/
}
func pointer_4pro() {
	var a *int
	a = new(int) //此时通过new初始化为对应类型的默认值
	fmt.Println(*a)
	*a = 100
	fmt.Printf("%#v %#v\n", a, *a)
	/*
		new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。
		var a *int 只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值。
	*/
	an := new(int)
	bn := new(bool)
	fmt.Printf("%T\n", an) // *int
	fmt.Printf("%T\n", bn) // *bool
	fmt.Println(*an)       // 0
	fmt.Println(*bn)       // false
}

func pointer_5() {
	var b map[string]int //只是声明变量b是一个map类型的变量，此处没有初始化变量b
	b["小桃子"] = 100       //此时b也是nil 注释此行代码输出结果为map[string]int(nil)
	fmt.Println(b)
	fmt.Printf("%#v\n", b)
	/*
		panic: assignment to entry in nil map
	*/
}

func pointer_5pro() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["小桃子"] = 100
	fmt.Printf("%#v\n", b)
	/*
		make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，
		而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
	*/
}

/*
new与make的区别:
	>* 二者都是用来做内存分配的。
	>* make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	>* 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
*/
