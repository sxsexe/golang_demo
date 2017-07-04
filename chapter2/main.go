package main

import "fmt"

func getName() (firstName, middleName, lastName string) {
	return "May", "April", "June"
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
	const {
		Sunday = iota  	//大写字母开头，包外可见
		Monday
		Friday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays  //小写字母开头，包外不可见
	}

}
