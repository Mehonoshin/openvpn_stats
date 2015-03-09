package server

import (
  "fmt"
  "openvpn_stats/dto"
)

type Statistics struct {
}

type Server struct {
}

type Client struct {
}

func NewStatistics(channel chan []dto.Client) {
  for clients := range channel {
    fmt.Println("[Stats received]", clients)
  }
}
