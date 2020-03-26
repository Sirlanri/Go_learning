package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func sayhelloName(w http.Response, r *http.Request) {
	r.ParseForm() //解析url参数，解析post的主体
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))

	}
	fmt.Fprintf(w, "hello, practice!") //写入到w，输出到客户端
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		//如果的登录请求，就执行逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("methods:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handle, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		fmt.Fprintln(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

}

//客户端上传
func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	//下面是一个关键（但是我现在也不明白鸭
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("向buffer写入过程中出错！")
		return err
	}
	//打开文档句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("打开文件出错")
		return err
	}
	defer fh.Close()

	//io copy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func main_one() {
	//特意来练练主函数（背下来就行
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置login路由
	http.HandleFunc("/upload", upload)       //上传文件的路由
	err := http.ListenAndServe(":9090", nil) //监听端口

	//下面是客户端上传文件的主函数
	target_url := "http://localhost:9090/upload"
	filename := "./astaxie.pdf"
	postFile(filename, target_url)
}
