package agent

import (
  "fmt"
  "time"
)

const RefreshInterval = 1000

func Run(filename string) {
  channel := make(chan []Client)
  go loadConnections(filename, channel)
  for data := range channel {
    fmt.Println(data)
  }
}

func loadConnections(filename string, ch chan []Client) {
  for {
    lines := read(filename)
    data  := parseConnections(lines)
    ch <- data
    time.Sleep(RefreshInterval * time.Millisecond)
  }
}
