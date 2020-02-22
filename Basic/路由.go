package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func creaat() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Println(vars["title"], vars["page"]) //解析地址中的内容

	})
	err := http.ListenAndServe(":80", r)
	if err != nil {
		println(err)
	}
}

func test() {
	creaat()

}
