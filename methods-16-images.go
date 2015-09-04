package main

import (
	//"fmt"
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width, height int
	image         [][]uint8
}

func (self *Image) Initialize() {
	self.image = make([][]uint8, self.width)

	for x := range self.image {
		self.image[x] = make([]uint8, self.height)

		for y := range self.image[x] {
			self.image[x][y] = 0
		}
	}
}

func (self *Image) Draw() {
	for x := range self.image {
		self.image[x] = make([]uint8, self.height)
		for y := range self.image[x] {
			self.image[x][y] = uint8((x + y)/3)
		}
	}
}

func (self *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (self *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, self.width, self.height)
}

func (self *Image) At(x, y int) color.Color {

	c := self.image[x][y]
	return color.RGBA{uint8(c), uint8(c), 255, 255}
}

func main() {
	m := Image{width: 200, height: 200}
	m.Initialize()
	m.Draw()
	pic.ShowImage(&m)
	//fmt.Println(m.image)
}
