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
	"path"
	"runtime/debug"
	"strings"
	"time"
)

const (
	UPLOAD_DIR   = "./upload"
	TEMPLATE_DIR = "./views"
)

var gTemplates map[string]*template.Template = make(map[string]*template.Template)

//在main之前执行
func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	checkErr(err)
	t1 := time.Now()

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if exist := path.Ext(templateName); exist != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Load Template : ", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		tmpl := strings.TrimSuffix(templateName, ".html")
		gTemplates[tmpl] = t
	}
	t2 := time.Now()
	log.Println("Loading Templates took ", t2.Sub(t1))

}

func renderHtml(w http.ResponseWriter, tml string, locals map[string]interface{}) error {
	err := gTemplates[tml].Execute(w, locals)
	return err
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := renderHtml(w, "upload", nil)
		checkErr(err)
	} else if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		checkErr(err)

		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		checkErr(err)
		defer t.Close()

		_, err = io.Copy(t, f)
		checkErr(err)

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
	checkErr(err)

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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError) // 或者输出自定义的 50x 错误页面
				w.WriteHeader(http.StatusInternalServerError)

				log.Println("WARN : painc in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

func main() {

	http.HandleFunc("/upload", safeHandler(uploadHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/list", safeHandler(listHandler))
	http.HandleFunc("/", safeHandler(listHandler))

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal("ListenAndService : ", err.Error())
	}

}
