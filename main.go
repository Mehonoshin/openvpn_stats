package main

import (
  "os"
  "fmt"
  "openvpn_stats/args"
  "openvpn_stats/agent"
  "openvpn_stats/server"
)

func main() {
  args := args.Parse(os.Args)

  switch args.Mode {
  case "agent":
    agent.Run(args.Source)
  case "server":
    server.Run(args.BindAddress, args.BindPort)
  case "mixed":
    server.Run(args.BindAddress, args.BindPort)
    agent.Run(args.Source)
  default:
    fmt.Println("No agent, no server running")
  }
}
