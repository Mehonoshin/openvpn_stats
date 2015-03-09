package server

import (
  "net"
  "fmt"
  "bufio"
  "encoding/json"
  "openvpn_stats/agent"
)

func Run(address, port string) {
  fmt.Println("[SERVER] Server started")
  hostWithPort := fmt.Sprintf("%s:%s", address, port)

  ln, err := net.Listen("tcp", hostWithPort)
  if err != nil {
    fmt.Println("[SERVER] Unable to bind server on", hostWithPort)
  }

  for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Println("[SERVER] Unable to handle connect")
    }
    go handleConnection(conn)
  }
}

func handleConnection(conn net.Conn) {
  fmt.Println("[SERVER] Handling connection")
  scanner := bufio.NewScanner(conn)
  scanner.Scan()
  message := scanner.Text()

  var clients []agent.Client
  err := json.Unmarshal([]byte(message), &clients)
  if err != nil {
    fmt.Println("Unable to decode message")
  }
  fmt.Println("Clients:", clients)
  conn.Close()
}
