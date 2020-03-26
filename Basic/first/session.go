package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

//这里是官方的简单实例
var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	//如果用户登录过了
	if auth, ok := session.Value["autenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	fmt.Println(w, "这是假的")
}
func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	//把user设置成有权限的
	session.Values["authenticated"] = true
	session.Save(r, w)
}
func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func test() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/loginout", loginout)

	http.ListenAndServe(":8090", nil)
}
