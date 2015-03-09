package agent

import (
  "fmt"
  "net"
  "time"
  "encoding/json"
  "openvpn_stats/dto"
)

const RefreshInterval = 2000

func Run(filename, serverAddress, serverPort string) {
  fmt.Println("[AGENT] Agent started")
  for {
    data := loadConnections(filename)
    go sendToServer(data, serverAddress, serverPort)
    time.Sleep(RefreshInterval * time.Millisecond)
  }
}

func loadConnections(filename string) []dto.Client {
  lines := read(filename)
  data  := parseConnections(lines)
  return data
}

func sendToServer(data []dto.Client, serverAddress, serverPort string) {
  hostWithPort := fmt.Sprintf("%s:%s", serverAddress, serverPort)
  conn, err := net.Dial("tcp", hostWithPort)
  if err != nil {
    fmt.Println("[AGENT][ERROR] Unable to connect to server", hostWithPort)
  }
  message, err := json.Marshal(data)
  if err != nil {
    fmt.Println("[AGENT][ERROR] Unable to serizlize data to json")
  }
  fmt.Fprintf(conn, string(message))
  conn.Close()
}
