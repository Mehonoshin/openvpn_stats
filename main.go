package main

import (
  "os"
  "openvpn_stats/agent"
)

func main() {
  filename := os.Args[1]
  //serverAddress := os.Args[2]
  //serverPort := os.Args[3]
  agent.Run(filename)
}
