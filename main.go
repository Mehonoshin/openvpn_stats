package main

import (
  "fmt"
  "os"
  "time"
  "openvpn_stats/reader"
)

const RefreshInterval = 1000

func loadConnections(filename string, ch chan []reader.Client) {
  for {
    lines := reader.Read(filename)
    data  := reader.ParseConnections(lines)
    ch <- data
    time.Sleep(RefreshInterval * time.Millisecond)
    fmt.Println("All data send for now")
  }
}

func main() {
  filename := os.Args[1]
  channel := make(chan []reader.Client)

  go loadConnections(filename, channel)
  for data := range channel {
    fmt.Println(data)
  }
}
