package reader

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type Client struct {
  CommonName     string
  IpAddress      string
  BytesSent      string
  BytesReceived  string
  ConnectedSince string
}

func (c Client) String() string {
  return fmt.Sprintf("%v, %v, %v, %v, %v", c.CommonName, c.IpAddress, c.BytesSent, c.BytesReceived, c.ConnectedSince)
}

func Read(filename string) []string {
  fmt.Println("Reading file", filename)

  file, err := os.Open(filename)
  if err != nil {
    return nil
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    lines = append(lines, line)
  }

  return lines
}

func ParseConnections(lines []string) []Client {
  var clients []Client
  clients = make([]Client, 0)

  for i := 0; i < len(lines); i++ {
    tokens := strings.Split(lines[i], ",")
    if (len(tokens) == 5) && (tokens[0] != "Common Name") {
      client := Client{tokens[0], tokens[1], tokens[2], tokens[3], tokens[4]}
      clients = append(clients, client)
    }
  }

  return clients
}
