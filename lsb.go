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

func getValues(file io.Reader)
