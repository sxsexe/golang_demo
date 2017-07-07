package main

import "fmt"
import "time"
import "runtime"

func Add(x, y int) int {
	z := x + y
	fmt.Println("z = ", z)
	return z
}

func Counter(ch chan int, i int) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Counting i ", i)
	ch <- i

}

const NCPU = 8

type Vector []float64

func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	fmt.Printf("DoSome i=%d, n=%d\n", i, n)
	var sum float64 = 0
	for ; i < n; i++ {
		sum += v[i]
	}
	time.Sleep(3000 * time.Millisecond)
	c <- int(sum)
}

func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU)
	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}

	var sum int = 0
	for i := 0; i < NCPU; i++ {
		fmt.Printf("i=%d is blocking \n", i)
		sum += <-c
	}

	fmt.Printf("sum=%d \n", sum)

}

func main() {

	//如下写法不会有任何打印，因为10个goroutine还没执行完main就结束了
	//	for i := 0; i < 10; i++ {
	//		go Add(i, i*2)
	//	}

	length := 3
	chs := make([]chan int, length)
	for i := 0; i < length; i++ {
		chs[i] = make(chan int)
		if i == 2 {
			go Counter(chs[i], i)
		}

	}

	fmt.Println("main before ch i ")
	var j = 9
	select {
	case <-chs[0]:
		fmt.Printf("main after ch i=%d, j=%d \n", 0, j)
	case <-chs[1]:
		fmt.Printf("main after ch i=%d, j=%d \n", 1, j)
	case <-chs[2]:
		fmt.Printf("main after ch i=%d, j=%d \n", 2, j)
	}

	//	j = <-chs[0] // 这里会导致程序报错，除了main goroutine以为没有任何等待的goroutine
	//	fmt.Println("main over")

	runtime.GOMAXPROCS(8)

	aLength := 1000
	var v Vector = make(Vector, aLength)
	for i := 0; i < len(v); i++ {
		v[i] = float64(i)
	}

	t1 := time.Now()
	//按CPU个数(16)分配任务，最后汇总
	v.DoAll(v)
	t2 := time.Now()
	fmt.Println("took ", t2.Sub(t1))

	fmt.Println("NumCPU = ", runtime.NumCPU())

}
