package main

import (
	"fmt"
	"image"
	"os"
	"io"
	"image/png"
	"log"
)

func main() {
	fmt.Print("Hello QR Code!")
	file, _ := os.Create("qrcode.png")
	defer file.Close()
	err := GenerateQRCode(file, "555-432", Version(1))
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateQRCode(w io.Writer, code string, version Version) error {
	size := version.PatternSize()
	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	return png.Encode(w, img)
}

func (v Version) PatternSize() int {
	return 4*int(v) + 17
}

type Version int8
