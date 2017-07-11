package icmptest

import (
	//	"bytes"
	"fmt"
	"net"
	"os"
)

type IcmpTestDemo struct {
}

func (demo *IcmpTestDemo) DoTest(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage : ", os.Args[0], " host")
	}
	service := args[1]

	conn, err := net.Dial("ip4:icmp", service)
	checkError(err)

	var msg [512]byte
	msg[0] = 8  // echo
	msg[1] = 0  // code 0
	msg[2] = 0  // checksum
	msg[3] = 0  // checksum
	msg[4] = 0  // identifier[0]
	msg[5] = 13 //identifier[1]
	msg[6] = 0  // sequence[0]
	msg[7] = 37 // sequence[1]
	len := 8

	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	fmt.Println("Write msg ", msg[0:len])
	_, err = conn.Write(msg[0:len])
	checkError(err)

	fmt.Println("Reading Response")
	_, err = conn.Read(msg[0:])
	checkError(err)

	fmt.Println("Got Response")
	if msg[5] == 13 {
		fmt.Println("Identifier matches")
	}
	if msg[7] == 37 {
		fmt.Println("Sequence matches")
	}
	os.Exit(0)

}

func checkSum(msg []byte) uint16 {
	sum := 0
	// assume even for now
	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
