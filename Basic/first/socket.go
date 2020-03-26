package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func main()  {
	http.HandleFunc("/echo", func(w http.ResponseWriter,r *http.Request){
		conn, _:=upgrader.Upgrade(w http.ResponseWriter, r *http.Request, nil)

		for{
			//不停的循环呀循环鸭
			msgType,msg,err:=conn.ReadMessage()
			if err != nil {
				return
			}
			fmt.Printf("%s sent: %s\n",conn.RemoteAddr(),string(msg))
			if err=conn.WriteMessage(msgType,msg);err!=nil{
				return
			}
			
		}
	})

	http.HandleFunc("/",func (w http.ResponseWriter,r *http.Request){
		http.ServerFile(w,r,"websocket.html")
	})
	
	http.ListenAndServe(":8080",nil)
	
}
