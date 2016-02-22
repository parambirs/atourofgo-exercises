package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := 1.0
  for {
    newz := z - (z * z - x) / (2 * z)
    if math.Abs(newz - z) < 1e-10 {
      break
    } else {
      z = newz
    }
  }
  return z
}

func main() {
  fmt.Println(Sqrt(4))
}

/*
> go run exercise-loops-and-functions.go
2.000000000000002
*/