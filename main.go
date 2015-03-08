package main

import (
  "os"
  "fmt"
  "openvpn_stats/args"
  "openvpn_stats/agent"
)

func main() {
  args := args.Parse(os.Args)

  if args.Mode == "agent" || args.Mode == "mixed" {
    agent.Run(args.Source)
  } else {
    fmt.Println("Agent not running")
  }
}
