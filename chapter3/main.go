package main // main.go

import "fmt"

// 增添新类型Integer,并且增加方法Less, Add
type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

// 返回一个指针
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

//继承(组合)
type Base struct {
	Name string
}

func (b *Base) Foo() {
	fmt.Println("Base.Foot")
}
func (b *Base) Bar() {
	fmt.Println("Base.Bar")
}

//虽然是小写,但是仍然可以被继承, 小写是限制了包外的访问
func (b *Base) foo() {
	fmt.Println("Base.foo")
}

//继承自Base
type Foo struct {
	Base
}

//重写了Base.Bar
func (foo *Foo) Bar() {
	foo.Base.Bar() // 调用了Base的Bar方法
	fmt.Println("Foo.Bar")
}

func main() {

	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, " less 2")
	}

	a.Add(2)
	fmt.Println("a Add 2 = ", a)

	// Go中的数组是值类型,赋值会复制
	var array [3]int = [3]int{1, 2, 3}
	var bb = array // 发生数据拷贝
	bb[1]++
	fmt.Println(array, bb)
	// 如果希望bb和array指向同一串数据,需要用到指针
	var bbb = &array
	bbb[1]++
	fmt.Println(array, bbb)
	//或者使用切片. 注意切片之间的赋值为值复制
	var ages []int = []int{1, 2, 3, 5}
	var b = ages // 指向相同的数组
	b[0]++
	fmt.Println(ages, b)

	// struct
	// 实例化的方法有以下几种
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}

	area1 := rect1.Area()
	area2 := rect2.Area()
	area3 := rect3.Area()
	area4 := rect4.Area()
	fmt.Printf("area1=%f, area2=%f, area3=%f, area4=%f \n", area1, area2, area3, area4)

	//GO中没有构造函数的概念,一般使用一个全局函数NewXXX完成
	rect5 := NewRect(0, 0, 200, 13)
	area5 := rect5.Area()
	fmt.Println("area5=", area5)

	//继承
	base := &Base{"N_Base"}
	foo := &Foo{}
	base.Foo()
	foo.Foo()
	foo.Bar()

	//Go语言中任何对象实例都满足空接口interface{}

}
