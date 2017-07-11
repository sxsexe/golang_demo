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
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	UPLOAD_DIR = "./upload"
)

func renderHtml(w http.ResponseWriter, tml string, locals map[string]interface{}) error {
	t, err := template.ParseFiles(tml + ".html")
	if err != nil {
		return err
	}

	err = t.Execute(w, locals)
	return err
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := renderHtml(w, "upload", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close()

		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/vie?/id="+filename, http.StatusFound)

	} else {
		fmt.Println("Unsupported method ", r.Method)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	fmt.Println("view ", imagePath)
	if exist := isFileExists(imagePath); !exist {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "image")
	//将该路径下的文件从磁盘中 读取并作为服务端的返回信息输出给客户端
	http.ServeFile(w, r, imagePath)
}

func isFileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	fmt.Println("list ", images)
	locals["images"] = images
	err = renderHtml(w, "list", locals)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func main() {

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/", listHandler)

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal("ListenAndService : ", err.Error())
	}

}
