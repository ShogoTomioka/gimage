package gimage

import (
	"image"
	"image/color"
)

type binaryImage struct {
}

//画像データを二値化して返却する
func (b binaryImage) Binarization(imgObject image.Image) *image.Gray {

	rec := imgObject.Bounds()
	binary := image.NewGray(rec)

	// グレーイメージに対して二値化処理
	for v := rec.Min.Y; v < rec.Max.Y; v++ {
		for h := rec.Min.X; h < rec.Max.X; h++ {
			c := color.GrayModel.Convert(imgObject.At(h, v))
			gray, _ := c.(color.Gray)
			// しきい値(128)で二値化
			if gray.Y > 128 {
				gray.Y = 255
			} else {
				gray.Y = 0
			}
			binary.Set(h, v, gray)
		}
	}
	return binary
}
