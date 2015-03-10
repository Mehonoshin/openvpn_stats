package server

import (
  "net"
  "fmt"
  "bufio"
  "encoding/json"
  "openvpn_stats/dto"
)

func Run(address, port string) {
  fmt.Println("[SERVER] Server started")

  add := make(chan []dto.Client)
  get := make(chan string)

  go newStatistics(add, get)
  go startHttpServer(get)
  startTcpServer(address, port, add)
}

func startTcpServer(address, port string, channel chan []dto.Client) {
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
    go handleConnection(conn, channel)
  }
}

func handleConnection(conn net.Conn, channel chan []dto.Client) {
  fmt.Println("[SERVER] Handling connection")
  scanner := bufio.NewScanner(conn)
  scanner.Scan()
  message := scanner.Text()

  var clients []dto.Client
  err := json.Unmarshal([]byte(message), &clients)
  if err != nil {
    fmt.Println("Unable to decode message")
  }
  channel <- clients
  conn.Close()
}
