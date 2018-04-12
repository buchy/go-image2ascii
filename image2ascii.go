package image2ascii

import (
	"image"
)

type LumAsciiMap struct {
	Luminance float64
	Ascii     byte
}

var conf = []LumAsciiMap{
	{0.0, '@'},
	{0.1, 'X'},
	{0.2, '%'},
	{0.3, '*'},
	{0.4, 'o'},
	{0.5, '='},
	{0.6, '<'},
	{0.65, '>'},
	{0.7, '+'},
	{0.75, '|'},
	{0.8, '-'},
	{0.85, '~'},
	{0.9, ','},
	{0.95, '.'},
	{1.0, ' '},
}

func SetConf(userConf []LumAsciiMap) {
	conf = userConf
}

func Convert(img image.Image) string {
	bounds := img.Bounds()
	var ascii []byte
	gray := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray.Set(x, y, img.At(x, y))
			ascii = append(ascii, pix2ascii(gray.GrayAt(x, y).Y))
		}
		ascii = append(ascii, '\n')
	}
	return string(ascii)
}

func pix2ascii(pix uint8) byte {
	luminance := float64(pix) / 255
	var ascii byte
	for _, v := range conf {
		if v.Luminance <= luminance {
			ascii = v.Ascii
		}
	}
	return ascii
}
