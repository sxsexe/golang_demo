package main

import (
	"fmt"
	"math/rand"
	"time"
)

func produce(channel chan string) {
	for {
		value := fmt.Sprintf("%v", rand.Intn(100))

		channel <- value
		fmt.Println("Produce ", value)
		time.Sleep(time.Second * time.Duration(1))
	}
}

func custome(channel chan string) {
	for {
		value := <-channel
		fmt.Println("Custome ", value)
		time.Sleep(time.Second * time.Duration(5))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	channel := make(chan string, 5) // 定义带有5个缓冲区位置的信道
	go produce(channel)

	custome(channel)

}
