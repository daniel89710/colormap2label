package main

import (
	"image"
	"image/png"
	"os"
	//	"image/draw"
	"fmt"
	"image/color"
)

func read_image(filename string) image.Image {
	//     fmt.Println("###Read ", filename, "...")
	src, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer src.Close()
	img, _, err := image.Decode(src)
	return img
}

func write_image(filename string, img image.Image) {
	fmt.Println("###Write ", filename, "...")
	outfile, _ := os.Create(filename)
	defer outfile.Close()
	png.Encode(outfile, img)
}

func color2gray(img image.Image, colormap map[string]Colormap, filename string) *image.RGBA {
	rect := image.Rectangle{image.Pt(0, 0), img.Bounds().Size()}
	dst := image.NewRGBA(rect)
	width := rect.Max.X
	height := rect.Max.Y
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r_, g_, b_, _ := img.At(x, y).RGBA()
			r := uint8(r_ & 0xff)
			g := uint8(g_ & 0xff)
			b := uint8(b_ & 0xff)
			id := search_id_from_rgb(colormap, r, g, b)
			if id > 0 {
				dst.Set(x, y, color.RGBA{
					uint8(id), uint8(id), uint8(id), 255})
			}
			if id == -1 {
				//				fmt.Println("Unknown colormap", r, g, b, x, y, filename)
			}
		}
	}
	return dst
}
