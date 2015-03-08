package agent

import (
  "fmt"
  "time"
)

const RefreshInterval = 1000

func LoadConnections(filename string, ch chan []Client) {
  for {
    lines := read(filename)
    data  := parseConnections(lines)
    ch <- data
    time.Sleep(RefreshInterval * time.Millisecond)
    fmt.Println("Sending data to server")
  }
}
