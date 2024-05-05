package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// print duration
	defer func(t time.Time) {
		println("Duration: ", time.Since(t).String())
	}(time.Now())

	// get image from arg
	if len(os.Args) < 2 {
		panic("error: no image path provided!")
	}
	imgPath := os.Args[1]

	// get scale from arg
	if len(os.Args) < 3 {
		panic("error: no width is provided!")
	}
	width, _ := strconv.Atoi(os.Args[2])
	if len(os.Args) < 4 {
		panic("error: no height is provided!")
	}
	height, _ := strconv.Atoi(os.Args[3])

	// load image from path
	imgFile, err := os.Open(imgPath)
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()

	// decode image
	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	// get gray image with new size
	newImage := image.NewGray(image.Rect(0, 0, width, height))

	scale_x := float64(width) / float64(img.Bounds().Dx())
	scale_y := float64(height) / float64(img.Bounds().Dy())

	for y := 0; y < height; y++ {
		srcY := int(math.Round(float64(y) / scale_y))
		for x := 0; x < width; x++ {
			srcX := int(math.Round(float64(x) / scale_x))
			newImage.Set(x, y, img.At(srcX, srcY))
		}
	}

	// save image to file
	fileName := strings.ReplaceAll(os.Args[1], ".", "_") + "_to_ascii.txt"
	newImageFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer newImageFile.Close()

	// map string to grayscale pixel value
	ascii, out := []byte{' ', '.', '-', '~', '+', '=', '%', '$', '#', '@'}, []byte{}

	bounds := newImage.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			y := float64(newImage.GrayAt(x, y).Y)
			val := int(y / 255 * float64(len(ascii)-1))
			out = append(out, ascii[val])
		}
		out = append(out, '\n')
	}
	newImageFile.Write(out)

	print("File conversion completed\n")
	print("Output file: ", fileName, "\n")
}
