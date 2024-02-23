package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
)

type tile struct {
	X    int         `json:"x"`
	Y    int         `json:"y"`
	Name string      `json:"name"`
	Img  image.Image `json:"-"`
}

var tiles []*tile

const perRow = 8

var outFile = flag.String("out", "output.png", "png file to write the sprite sheet")

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println("Usage: spritemapper --out /path/to/output.png image1.png image2.png")
		os.Exit(1)
	}

	for _, name := range flag.Args() {
		img, err := readPng(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		tiles = append(tiles, img)
	}

	w, h := spriteBounds()

	rows := len(tiles) / perRow

	canvas := image.NewRGBA(image.Rect(0, 0, perRow*w, h*rows))

	x, y := 0, 0
	for _, tile := range tiles {
		dp := image.Point{x * w, y * h}
		fb := tile.Img.Bounds()

		r := image.Rectangle{dp.Add(fb.Min), dp.Add(fb.Max)}
		tile.X = x
		tile.Y = y

		draw.Draw(canvas, r, tile.Img, fb.Min, draw.Over)

		x += 1
		if x >= perRow {
			y += 1
			x = 0
		}
	}
	out, err := os.Create(*outFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := png.Encode(out, canvas); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tileJson, err := json.Marshal(tiles)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(tileJson))
}

func spriteBounds() (w, h int) {
	bounds := tiles[0].Img.Bounds()

	w, h = bounds.Size().X, bounds.Size().Y

	for _, f := range tiles {
		bounds = f.Img.Bounds()
		if bounds.Size().X > w {
			w = bounds.Size().X
		}
		if bounds.Size().Y > h {
			h = bounds.Size().Y
		}
	}

	return
}

func readPng(name string) (*tile, error) {
	fileName := filepath.Base(name)

	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		return nil, err
	}
	return &tile{Img: img, Name: fileName}, nil
}
