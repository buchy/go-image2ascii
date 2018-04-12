package main

import (
	"os"
	"image"
	"fmt"
	"github.com/buchy/go-image2ascii"
	_ "image/jpeg"
)

func main() {
	file, err := os.Open("_example/test.jpg")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	image2ascii.SetConf([]image2ascii.LumAsciiMap{
		{0.0, 'a'},
		{0.1, 'b'},
		{0.2, 'c'},
		{0.3, 'd'},
		{0.4, 'e'},
		{0.5, 'f'},
		{0.6, 'g'},
		{0.7, 'h'},
		{0.8, 'i'},
		{0.9, 'j'},
		{1.0, 'k'},
	})
	fmt.Println(image2ascii.Convert(img))
}
