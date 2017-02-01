package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	r := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		r[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			r[y][x] = uint8(x * y)
		}
	}
	return r
}

func main() {
	pic.Show(Pic)
}
