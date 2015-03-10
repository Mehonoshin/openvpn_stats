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

func handler(w http.ResponseWriter, r *http.Request, get chan string) {
  get<- "get stats"
  response := <-get
  io.WriteString(w, response)
}
