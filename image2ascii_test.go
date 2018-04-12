package image2ascii

import (
	"testing"
	"image"
	"image/color"
)

func BenchmarkConvert(b *testing.B) {
	Convert(image.NewGray(image.Rect(0, 0, 1500, 1500)))
}

func TestConvert(t *testing.T) {
	SetConf([]lumAsciiMap{
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

	img := image.NewGray(image.Rect(0, 0, 256, 1))
	for i := img.Rect.Min.X; i < img.Rect.Max.X; i++ {
		img.Set(i, 0, color.Gray{uint8(i)})
	}
	actual := Convert(img)
	expect := "aaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbccccccccccccccccccccccccccdddddddddddddddddddddddddeeeeeeeeeeeeeeeeeeeeeeeeeefffffffffffffffffffffffffgggggggggggggggggggggggggghhhhhhhhhhhhhhhhhhhhhhhhhiiiiiiiiiiiiiiiiiiiiiiiiiijjjjjjjjjjjjjjjjjjjjjjjjjk\n"

	if actual != expect {
		t.Errorf("actual:%v, expect:%v", actual, expect)
	}
}
