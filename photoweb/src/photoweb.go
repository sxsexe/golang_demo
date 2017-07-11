/**
	一个小型的相片分享网站，功能如下：
	1：支持图片上传
	2：查看已经上传的图片
	3：查看已经上传的图片列表
	4：删除指定图片

**/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, "<form method=\"POST\" action=\"/upload\" "+
			" enctype=\"multipart/form-data\">"+
			"Choose an image to upload: <input name=\"image\" type=\"file\" />"+"<input type=\"submit\" value=\"Upload\" />"+
			"</form>")
	} else {
		fmt.Println("Unsupported method ", r.Method)
	}
}

func main() {

	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal("ListenAndService : ", err.Error())
	}

}
