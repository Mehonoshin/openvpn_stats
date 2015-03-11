package server

import (
  "fmt"
  "strconv"
  "time"
  "encoding/json"
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
  Clients map[string]*Client
}

type Client struct {
  Events []Event
}

type Event struct {
  EventTimestamp int64
  BytesSent      int64
  BytesReceived  int64
}

func newStatistics(add chan []dto.Client, get chan string) {
  totalStatistics := Statistics{0, 0, 0}
  var servers map[string]Server
  servers = make(map[string]Server)

  for {
    select {
    case clients := <-add:
      servers = addEvent(servers, clients)
      totalStatistics = refreshStatistics(totalStatistics, servers)
      fmt.Println("Servers state now" )
      printServersState(servers)
    case <- get:
      get<- statsToJson(totalStatistics)
    }
  }
}

func addEvent(servers map[string]Server, newClients []dto.Client) map[string]Server {
  for _, newClient := range newClients {
    if _, ok := servers[newClient.Hostname]; !ok {
      servers[newClient.Hostname] = Server{make(map[string]*Client)}
    }
    server := servers[newClient.Hostname]

    if _, ok := server.Clients[newClient.CommonName]; !ok {
      server.Clients[newClient.CommonName] = &Client{}
    }

    bytesSent,     _ := strconv.ParseInt(newClient.BytesSent, 10, 0)
    bytesReceived, _ := strconv.ParseInt(newClient.BytesReceived, 10, 0)

    events := server.Clients[newClient.CommonName].Events
    events = addEventToList(events, Event{time.Now().Unix(), bytesSent, bytesReceived})

    server.Clients[newClient.CommonName].Events = events
    servers[newClient.Hostname] = server
  }
  return servers
}

func addEventToList(events []Event, event Event) []Event {
  if len(events) < cap(events)  {
    //shift and append
  }
  events = append(events, event)
  return events
}

func printServersState(servers map[string]Server) {
  for hostname, server := range servers {
    fmt.Println("Server", hostname)
    for commonName, client := range server.Clients {
      fmt.Println("Client", commonName)
      for _, events := range client.Events {
        fmt.Println("Events", events)
      }
    }
  }
}

func refreshStatistics(statistics Statistics, servers map[string]Server) Statistics {
  var totalIn, totalOut, totalClients int64
  //totalIn, totalOut, totalClients = statistics.TotalBytesIn, statistics.TotalBytesOut, statistics.ClientsConnected

  //for i := 0; i < len(clients); i++ {
    //bytesIn, _ := strconv.ParseInt(clients[i].BytesReceived, 10, 0)
    //totalIn = totalIn + bytesIn
    //bytesOut, _ := strconv.ParseInt(clients[i].BytesSent, 10, 0)
    //totalOut = totalOut + bytesOut
    //totalClients++
  //}
  return Statistics{totalIn, totalOut, totalClients}
}

func statsToJson(stats Statistics) string {
  message, _ := json.Marshal(stats)
  return string(message)
}
