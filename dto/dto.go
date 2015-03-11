package dto

import (
  "fmt"
)

type Client struct {
  CommonName     string
  IpAddress      string
  BytesSent      string
  BytesReceived  string
  ConnectedSince string
  Hostname       string
}

func (c Client) String() string {
  return fmt.Sprintf("%v, %v, %v, %v, %v", c.CommonName, c.IpAddress, c.BytesSent, c.BytesReceived, c.ConnectedSince)
}
