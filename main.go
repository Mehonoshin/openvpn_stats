package main

import (
  "os"
  "fmt"
  "time"
  "openvpn_stats/args"
  "openvpn_stats/agent"
  "openvpn_stats/server"
)

// TODO: add some logger

func main() {
  args := args.Parse(os.Args)
  fmt.Println("[MAIN] App starting")

  switch args.Mode {
  case "agent":
    go agent.Run(args.Source, args.ServerAddress, args.ServerPort)
  case "server":
    go server.Run(args.BindAddress, args.BindPort)
  case "mixed":
    go server.Run(args.BindAddress, args.BindPort)
    go agent.Run(args.Source, args.ServerAddress, args.ServerPort)
  default:
    fmt.Println("[MAIN] No agent, no server running")
  }

  for {
    time.Sleep(100 * time.Millisecond)
  }
}
