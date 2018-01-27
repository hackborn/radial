package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	width := 128
	height := width
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	center := Ptf{(float64(width) / 2.0) - 0.5, (float64(height) / 2.0) - 0.5}
	maxD := center.DistanceTo(Ptf{center.X, 0.0})

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pt := Ptf{float64(x) - 0.5, float64(y) - 0.5}
			d := center.DistanceTo(pt)
			unit := 1.0 - (d / maxD)

			// Color, with a little shaping
			shaped := math.Pow(unit, 0.25)
			b := toByte(shaped)

			// Alpha
			var a uint8
			if d < maxD {
				// Antialias the edge pixels.
				a = toByte(math.Abs(d - maxD))
			}

			img.Set(x, y, color.NRGBA{
				R: b,
				G: b,
				B: b,
				A: a,
			})
		}
	}

	err := saveImage("image.png", img)
	if err != nil {
		log.Fatal(err)
	}
}

// func toByte() converts a unit float (0-1) to a byte (0-255).
func toByte(v float64) uint8 {
	v = v * 255
	if v >= 255 {
		return uint8(255)
	} else if v > 0 {
		return uint8(v)
	}
	return uint8(0)
}

// func saveImage() saves thje supplied image to the filename.
func saveImage(filename string, img *image.NRGBA) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}
