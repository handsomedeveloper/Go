package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	error := http.ListenAndServe(":9090", nil)
	if error != nil {
		log.Fatal("ListenAndServe: ", error)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Printf("%+v", r)
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value: ", v)
		fmt.Println(strings.Join(v, ","))
	}
	fmt.Fprintf(w, "Hello %s!", "world")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
		fmt.Fprintf(w, r.Form.Get("username"))
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println("Upload Method: ", method)
	if method == "GET" {
		timestamp := time.Now().Unix()
		hash := md5.New()
		io.WriteString(hash, strconv.FormatInt(timestamp, 10))
		token := fmt.Sprintf("%x", hash.Sum(nil))
		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadFile")
		defer file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v\n", file)
		fmt.Printf("%+v\n", handler)

		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		defer f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		io.Copy(f, file)
	}
}
