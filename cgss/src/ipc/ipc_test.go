package ipc

import (
	"fmt"
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{"ECHO:" + method, "ECHO:" + params}
}

func (sever *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIPCServer(&EchoServer{})

	client1 := NewIPCClient(server)
	client2 := NewIPCClient(server)

	resp1, _ := client1.Call("From Client1", "Params From Client1")
	resp2, _ := client2.Call("From Client2", "Params From Client2")

	fmt.Println("resp1 = ", resp1)
	fmt.Println("resp2 = ", resp2)

	client1.Close()
	client2.Close()
}
