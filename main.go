package main

import (
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

var executableName = filepath.Base(os.Args[0])

func main() {
	// print duration
	defer func(t time.Time) {
		fmt.Println("\nDuration:", time.Since(t).String())
	}(time.Now())

	app := cli.NewApp()
	app.Name = executableName
	app.Usage = "An image to ascii converter"
	app.UsageText = executableName + " [path-to-image] [character-width]"
	app.HideHelp = true
	app.Action = func(ctx *cli.Context) error {
		if ctx.NArg() < 1 {
			return errors.New("no image path provided")
		}

		if ctx.NArg() < 2 {
			return errors.New("no width is provided")
		}

		// load image
		imgPath := ctx.Args().First()
		imgFile, err := os.Open(imgPath)
		if err != nil {
			return err
		}
		defer imgFile.Close()

		// decode image
		img, _, err := image.Decode(imgFile)
		if err != nil {
			return err
		}

		// get width from args
		width, _ := strconv.Atoi(ctx.Args().Get(1))
		if width < 1 {
			width = img.Bounds().Dx()
		}

		// calculate height while maintaining aspect ratio
		height := width * img.Bounds().Dy() / img.Bounds().Dx()

		// create new gray image
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
			return err
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

		fmt.Println("File conversion completed\n\nOutput file:", fileName)

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("ERROR:\n", err.Error())
		fmt.Println("\nUSAGE:\n", app.UsageText)
	}
}
