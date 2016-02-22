package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	r := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		r = append(r, make([]uint8, dx))
		for x := 0; x < dx; x++ {
			r[y] = append(r[y], uint8(x*y))
		}
	}
	return r
}

func main() {
	pic.Show(Pic)
}
