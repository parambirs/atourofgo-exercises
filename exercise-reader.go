package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (int, error) {
  for i := range b {
    b[i] = 'A'
  }
  return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}

/*
> go run exercise-reader.go
OK!
*/