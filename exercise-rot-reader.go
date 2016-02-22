package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct{
  r io.Reader
}

func (rr rot13Reader) Read(b []byte) (int, error) {
  c, err := rr.r.Read(b)

  for i, ch := range b {
    switch {
    case ch >= 65 && ch <= 90:
      b[i] = ((b[i] - 65) + 13) % 26 + 65
    case ch >= 97 && ch <= 122:
      b[i] = ((b[i] - 97) + 13) % 26 + 97
    }
  }

  return c, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}

/*
> go run exercise-rot-reader.go
You cracked the code!⏎
*/
