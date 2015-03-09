package agent

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "openvpn_stats/dto"
)

func read(filename string) []string {
  fmt.Println("[AGENT] Reading file", filename)

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

func parseConnections(lines []string) []dto.Client {
  var clients []dto.Client
  clients = make([]dto.Client, 0)

  for i := 0; i < len(lines); i++ {
    tokens := strings.Split(lines[i], ",")
    if (len(tokens) == 5) && (tokens[0] != "Common Name") {
      client := dto.Client{tokens[0], tokens[1], tokens[2], tokens[3], tokens[4]}
      clients = append(clients, client)
    }
  }
  fmt.Println("[AGENT] Parsed connections")
  return clients
}
