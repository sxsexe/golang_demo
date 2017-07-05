package main

import "fmt"
import "math"

func getName() (firstName, middleName, lastName string) {
	return "May", "April", "June"
}

/**
*	p: 自定义精度
**/
func isEqual(f1, f2, p float64) bool {
	return math.Dim(f1, f2) < p
}

//视图修改传入的数组内容，实际上修改的是临时变量的内容
func modifyArray(array [10]int) {
	array[0] = array[0] + 1
	fmt.Println("In modify(), array values:", array)
}

// 接受任意参数类型的不定参数
func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, " is an int value")
		case float32:
			fmt.Println(arg, " is a float32 value")
		case float64:
			fmt.Println(arg, " is a float64 value")
		case string:
			fmt.Println(arg, " is a string value")
		default:
			fmt.Println(arg, " is a unknown type")
		}
	}
}

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	i := 3
	j := 4
	i, j = j, i // 快速交换，多重赋值
	fmt.Println("i = ", i, ", j = ", j)

	_, _, lastName := getName() // 忽略掉前两个返回值
	fmt.Println("lastName = ", lastName)

	const ( //const关键字出现，iota被置为0
		c0 = iota //c0 == 0 然后iota自增1
		c1 = iota //c1 == 1 然后iota自增1
		c2 = iota //c2 == 2 然后iota自增1
	)
	fmt.Printf("c0 = %d, c1 = %d, c2 = %d\n", c0, c1, c2)

	const c3 = iota //iota被置为0
	fmt.Printf("c3 = %d\n", c3)

	const (
		a = 1 << iota // 此时iota为0，1左移0位为1，iota自增1
		b = 1 << iota // 此时iota为1，1左移1位为2，iota自增1
		c = 1 << iota // 此时iota为2，1左移2位为4，iota自增1
	)
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	//如果两个const赋值语句是一样的，后一句可以省略赋值表达式
	const (
		aa = iota
		bb
		cc
	)
	fmt.Printf("aa = %d, bb = %d, cc = %d\n", aa, bb, cc)

	//golang中没有enum，可以用const实现
	const (
		Sunday = iota //大写字母开头，包外可见
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays //小写字母开头，包外不可见
	)

	//bool类型不能接受其他类型的赋值，不支持自动或强制类型转换
	//	var v1 bool
	//	v1 = 1 // 编译出错  cannot use 1 (type int) as type bool in assignment

	// int int32是两张不同的类型，不能互相赋值
	var value2 int32
	value1 := 64 // 自动推导为int
	//	value1 = value2 // 编译出错 cannot use value2 (type int32) as type int in assignment
	value2 = int32(value1) // 编译通过
	fmt.Println("value2=", value2)

	fmt.Println("5 % 3 = ", 5%3)

	// 不同类型的整型不能直接比较
	/**/

	//	var ii int32
	//	var jj int64
	//	ii, jj = 1, 2

	//编译出错
	//	if ii == jj { // invalid operation: ii == jj (mismatched types int32 and int64)
	//		fmt.Println("ii and jj are not equal")
	//	}

	// float32相当于C语言里的float  float64相当于double
	// 浮点数比较不能直接用==判断，不够准确 参考fun isEqual
	v1 := 12.0000002
	v2 := 12.0000001
	ret := isEqual(v1, v2, 0.000001)
	fmt.Println("isEqual ", ret) // true

	// 字符串遍历方法一
	str := "Hello, 世界"
	n := len(str)
	fmt.Println("len(str) = ", n)
	for i := 0; i < n; i++ {
		ch := str[i] //依据下标取字符串中的字符,类型为byte
		fmt.Println(i, ch, str[i])
	}

	// 字符串遍历方法二
	for i, ch := range str {
		fmt.Println(i, ch) //ch的类型为rune(字符)
	}

	array := [10]int{1, 2, 3, 4, 5}
	modifyArray(array) // 数组是值类型，传递会产生一份副本

	fmt.Println("Out modify(), array values:", array)

	// 切片
	mySlice := make([]int, 5, 10) // len = 5, cap = 10
	fmt.Println("len = ", len(mySlice), ", cap = ", cap(mySlice))
	mySlice2 := []int{7, 8, 9}
	mySlice = append(mySlice, mySlice2...) // 必须要有省略号，相当于把mySlice2打散后传入
	fmt.Println("len = ", len(mySlice), ", cap = ", cap(mySlice))
	fmt.Println("mySlice = ", mySlice)

	// map
	var personDB map[string]PersonInfo     // 声明一个map
	personDB = make(map[string]PersonInfo) // 创建一个map，如果不创建而直接使用会导致panic
	personDB["12345"] = PersonInfo{"12345", "Tom", "USA"}
	personDB["1"] = PersonInfo{"1", "Jack", "China"}

	person, ok := personDB["12345"] // ok表示是否找到对应的数据
	if ok {
		fmt.Println("Found person", person, "with ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}

	// 小写字母开头的函数和变量只是包内可见，首字母大写的才是包外可见

	//不定参数列表demo，传入类型为任意类型
	var vv1 int = 1
	vv2 := 2.0
	var vv3 int64 = 3
	var vv4 string = "hello"
	var vv5 float32 = 2.3
	var vv6 float64 = 6
	MyPrintf(vv1, vv2, vv3, vv4, vv5, vv6)

	// 闭包
	var cj int = 2
	ca := func() func() {
		var i int = 10
		return func() {
			fmt.Println("i = ", i, ", cj = ", cj)
		}
	}()

	ca()
	cj *= 2
	ca()

}
