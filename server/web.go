package server

import (
  "io"
  "net/http"
)

const HttpPort = "8070"

func startHttpServer(get chan string) {
  http.HandleFunc("/stat.json", func(w http.ResponseWriter, r *http.Request){
    handler(w, r, get)
  })
  http.ListenAndServe(":" + HttpPort, nil)
}

func statJsonHandler(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "Response")
}

func handler(w http.ResponseWriter, r *http.Request, get chan string) {

}
