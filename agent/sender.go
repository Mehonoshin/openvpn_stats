package agent

import (
  "fmt"
  "net"
  "time"
  "bufio"
)

const RefreshInterval = 2000

func Run(filename, serverAddress, serverPort string) {
  for {
    fmt.Println("[AGENT] Agent started")
    data := loadConnections(filename)
    fmt.Println("[AGENT] Sending data to server")
    go sendToServer(data, serverAddress, serverPort)
    time.Sleep(RefreshInterval * time.Millisecond)
  }
}

func loadConnections(filename string) []Client {
  lines := read(filename)
  data  := parseConnections(lines)
  return data
}

func sendToServer(data []Client, serverAddress, serverPort string) {
  fmt.Println("[AGENT][SENDING PAYLOAD]", data)

  hostWithPort := fmt.Sprintf("%s:%s", serverAddress, serverPort)
  conn, err := net.Dial("tcp", hostWithPort)
  if err != nil {
    fmt.Println("[AGENT][ERROR] Unable to connect to server", hostWithPort)
  }
  fmt.Fprintf(conn, "some data")
  status, err := bufio.NewReader(conn).ReadString('\n')
  fmt.Println("[AGENT][STATUS]", status)
}
