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

	fmt.Println(image2ascii.Convert(img))
}
