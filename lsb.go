package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"image"
	"image/png"
	"strings"
)

func main() {
	var err error

	fmt.Println("Absolute image path: ")
	reader := bufio.NewReader()
	path, err := reader.ReadString()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if strings.Contains(path, ".png") {
		image.RegisterFormat("png", "png", png.Decode(), png.DecodeConfig())
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()
}

type Pixel struct {
	r	u8
	g	u8
	b	u8
	a	u8
}

func imageToRGBA(img image.Image) []Pixel {
	size := img.Bounds()
	raw := make([]Pixel, (size.Max.X - size.Min.X) * (size.Max.Y - size.Min.Y))
	idx := 0

	for y := size.Min.Y; y < size.Max.Y; y++ {
		for x := size.Min.X; x < size.Max.X; x++ {
			r, g, b, a = img.At(x, y).RGBA()
			raw[i] = Pixel{r: r, g: g, b: b, a: a}
		}
	}
}
