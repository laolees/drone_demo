package main

import (
	"encoding/json"
	//"fmt"
	"io"
	//"io/ioutil"
	"log"
	"net/http"
	"path"
	"text/template"
)

//func SendHttpRequestDemo() {
//	fmt.Println("vim-go")
//
//	// 自定义请求
//	req, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
//	if err != nil {
//		panic(err)
//	}
//	req.Header.Add("User-Agent", " Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
//
//	response, err := http.DefaultClient.Do(req)
//	//	response, err := http.Get("http://www.baidu.com")
//	//	defer response.Body.Close()
//
//	if err != nil {
//		panic(err)
//	}
//
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(body))
//}

func main() {
	//	SendHttpRequestDemo()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if _, err := io.WriteString(writer, "Hello Golang\n"); err != nil {
			panic(err)
		}
	})

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	http.HandleFunc("/json", func(writer http.ResponseWriter, request *http.Request) {
		user := User{
			Name: "laolees",
			Age:  20,
		}
		userJson, err := json.Marshal(user)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		if _, err := writer.Write(userJson); err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/image", func(writer http.ResponseWriter, request *http.Request) {
		image := path.Join("images", "golang.jpg")
		http.ServeFile(writer, request, image)
	})

	http.HandleFunc("/html", func(writer http.ResponseWriter, request *http.Request) {
		user := User{
			Name: "cnych",
			Age:  20,
		}
		htmlFile := path.Join("templates", "index.html")
		tmpl, err := template.ParseFiles(htmlFile)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(writer, user); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
