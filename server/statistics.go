package server

import (
  "fmt"
  "strconv"
  "openvpn_stats/dto"
)

type Statistics struct {
  TotalBytesIn     int64
  TotalBytesOut    int64
  ClientsConnected int64
}

func (s Statistics) String() string {
  return fmt.Sprintf("%v, %v, %v", s.TotalBytesIn, s.TotalBytesOut, s.ClientsConnected)
}

type Server struct {
}

type Client struct {
}

func newStatistics(add chan []dto.Client, get chan string) {
  totalStatistics := Statistics{0, 0, 0}

  for {
    select {
    case clients := <-add:
      totalStatistics = countStatistics(totalStatistics, clients)
      fmt.Println("Current stats is:", totalStatistics)
    }
  }
}

func countStatistics(statistics Statistics, clients []dto.Client) Statistics {
  var totalIn, totalOut, totalClients int64
  totalIn, totalOut, totalClients = statistics.TotalBytesIn, statistics.TotalBytesOut, statistics.ClientsConnected

  for i := 0; i < len(clients); i++ {
    bytesIn, _ := strconv.ParseInt(clients[i].BytesReceived, 10, 0)
    totalIn = totalIn + bytesIn
    bytesOut, _ := strconv.ParseInt(clients[i].BytesSent, 10, 0)
    totalOut = totalOut + bytesOut
    totalClients++
  }
  return Statistics{totalIn, totalOut, totalClients}
}
