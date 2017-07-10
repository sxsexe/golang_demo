package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct { // 继承自Server
	Server
}

func NewIPCServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0)

	go func(c chan string) {
		for {
			fmt.Println("Connect : waiting for client")
			request := <-c //此处c即外边的session
			fmt.Println("Connect : client comming, req = ", request)
			if request == "CLOSE" {
				break
			}

			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}

			// 这里的server是IPCServer，通过Go语言的组合“继承”方式，实际上是调用
			// ipcServer.Server.Handle
			fmt.Println("IpcServer Handle method = ", req.Method)
			resp := server.Handle(req.Method, req.Params)

			b, err := json.Marshal(resp)

			fmt.Println("IpcServer resp = ", string(b))

			c <- string(b)

		}
	}(session)

	fmt.Println("A new session has been created successfully.")
	return session

}
