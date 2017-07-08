package main

import (
  // "os"
  // "fmt"
  // "io/ioutil"
  "net/http"
  "github.com/gorilla/websocket"
  "log"
)

var clients = make(map[*websocket.Conn] bool) //To keep a record if the clients are connected or not
var broadcast = make(chan Message) //A channel to broadcast message to all connected clients

var upgrader = websocket.Upgrader{
	CheckOrigin: func(request *http.Request) bool {
		return true
	},
}
type Message struct {
  Email string  `json:"email"`
  Username string `json:"username"`
  Message string  `json:"message"`
}

func main(){
  fs := http.FileServer(http.Dir("/Self_Projects/Khufiya/html/"))
  http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/websocket",handleConnections)

  go handleMessages()
  log.Println("http server started on :8080")
  err := http.ListenAndServe(":8080",nil )
  if err != nil {
    log.Fatal("ListenAndServe : ",err)
  }
}

func handleConnections(writer http.ResponseWriter, req *http.Request) {
  ws, err := upgrader.Upgrade(writer, req, nil)
  if err != nil {
    log.Fatal(err)
  }
  // Make sure we close the connection when the function returns
  defer ws.Close()
  clients[ws] = true
  for {
    var msg Message
    err := ws.ReadJSON(&msg)
    if err != nil{
      log.Printf("error: %v", err)
      delete(clients, ws)
      break
    }
    broadcast <- msg
  }
}

func handleMessages(){
  for {
    msg := <-broadcast
    for client := range clients {
      err := client.WriteJSON(msg)
      if err != nil {
        log.Printf("error: %v",err)
        client.Close()
        delete(clients,client)
      }
    }
  }
}
