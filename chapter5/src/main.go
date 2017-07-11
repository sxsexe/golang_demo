package main

import (
	"fmt"
	//	"icmptest"
	"os"
	"tcptest"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host")
		os.Exit(1)
	}

	//ICMP
	//	icmpDemo := &icmptest.IcmpTestDemo{}
	//	icmpDemo.DoTest(os.Args)

	//TCP
	tcpDemo := &tcptest.TcpTestDemo{}
	tcpDemo.DoTest(os.Args)

}
