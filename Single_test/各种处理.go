package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func jsondef() {
	type User struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Age       int    `json:"age"`
	}
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		_ = json.NewDecoder(r.Body).Decode(&user)
		fmt.Fprintf(w, "%s %s is %d years old", user.Firstname, user.Lastname, user.Age)
	})
	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "Rico",
			Lastname:  "Lan",
			Age:       20,
		}
		json.NewEncoder(w).Encode(peter)
	})
	http.ListenAndServe(":8081", nil)
}

func jsondef2() {
	type Server struct {
		ServerName string
		ServerIp   string
	}
	type Serverslice struct {
		Servers []Server
	}

	//下面是执行部分
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	for _, server := range s {
		fmt.Println(server.ServerName, server.ServerIp)
	}
	fmt.Println(s)
}

func jsondef3() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	_ := json.Unmarshal(b, &f)
	f = map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
		"Parents": []interface{}{
			"Gomez",
			"Morticia",
		},
	}
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.type() {
		case string:
			fmt.Println(k,"是字符串",vv)
		case int :
			fmt.Println(k,"是int",vv)
		case float64:
			fmt.Println(k,"是浮点数",vv)
		case []interface{}:
			fmt.Println(k,"是数组")
			for i, u := range vv{
				fmt.Println(i,u)
			}
		default:
			fmt.Println(k,"一个未知类型")
		}
		
	}
}
func main() {
	jsondef2()
}
