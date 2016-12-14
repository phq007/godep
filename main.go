package main

import (
	"image"
	"image/png"
	"log"
	"os"

	"code.google.com/p/graphics-go/graphics"
)

// Load Image decodes an image from a file of image.
func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

//Save Image to path from a processed image
func SaveImage(path string, img image.Image) (err error) {
	imgfile, err := os.Create(path)
	if err != nil {
		return
	}
	defer imgfile.Close()
	err = png.Encode(imgfile, img)
	return
}

func PngScale(imgsrc string, imgdst string, dx int, dy int) (err error) {
	src, err := LoadImage(imgsrc)
	if err != nil {
		return
	}
	dst := image.NewRGBA(image.Rect(0, 0, dx, dy))
	err = graphics.Scale(dst, src)
	if err != nil {
		return
	}
	SaveImage(imgdst, dst)
	return
}

func main() {
	log.Println("test")
}
