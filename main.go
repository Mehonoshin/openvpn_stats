package main

import (
  "fmt"
  "os"
  "openvpn_stats/agent"
)

func main() {
  filename := os.Args[1]
  //serverAddress := os.Args[2]
  //serverPort := os.Args[3]
  channel := make(chan []agent.Client)

  go agent.LoadConnections(filename, channel)
  for data := range channel {
    fmt.Println(data)
  }
}
