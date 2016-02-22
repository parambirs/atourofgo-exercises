package main

import (
  "fmt"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
  return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
  addrs := map[string]IPAddr{
    "loopback": {127, 0, 0, 1},
    "googleDNS": {8, 8, 8, 8},
  }

  for n, a := range addrs {
    fmt.Printf("%v: %v\n", n, a)
  }
}

/*
> go run exercise-stringer.go
loopback: 127.0.0.1
googleDNS: 8.8.8.8
*/