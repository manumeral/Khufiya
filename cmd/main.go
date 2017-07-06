package main

import (
  "os"
  "fmt"
  "io/ioutil"
  "net/http"
)

func main(){
  webApp,err := os.Open("html/Client.html")
  if err != nil {
    fmt.Println(err)
  }
  content,err := ioutil.ReadAll(webApp)
  if err != nil {
    fmt.Println(content,err)
  }
  http.HandleFunc("/websocket",func(writer http.ResponseWriter, req *http.Request) {
    writer.Write([]byte(content))
  })
  err = http.ListenAndServe(":8000",nil )
}
