package main

import (
  "golang.org/x/tour/pic"
  "image"
  "image/color"
)

type Image struct{
  X, Y int
}

func (i Image) ColorModel() color.Model {
  return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
  return image.Rect(0, 0, i.X, i.Y)
}

func (i Image) At(x, y int) color.Color {
  return color.RGBA{uint8(x^y), uint8(x^y), 255, 255}
}

func main() {
	m := Image{256, 256}
  pic.ShowImage(m)
}
