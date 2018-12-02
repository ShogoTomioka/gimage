package gimage

import (
	"image"
	"image/color"
	"os"
)

func NewColorImage(filename string) image.Image {

	//画像ファイルのオープン
	file, _ := os.Open(filename)
	defer file.Close()

	//ファイルをデコードしてImageオブジェクトを作成
	imageObj, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	return imageObj
}

//元画像をimgに描写して返却する。
func fillColor(img *image.NRGBA, srcImg image.Image, width int, height int) *image.NRGBA {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			R, G, B, _ := srcImg.At(x, y).RGBA()

			img.Set(x, y, color.NRGBA{
				R: uint8(R),
				G: uint8(G),
				B: uint8(B),
				A: 255,
			})
		}
	}
	return img
}

//DrawBound は指定されたRectangleの範囲に赤い枠線を描く
func DrawBound(img *image.RGBA, rect image.Rectangle) *image.RGBA {
	//間違っている部分を囲う枠線の色
	red := color.RGBA{255, 0, 0, 0}
	//rectの範囲に枠線を書く
	// 上下の枠
	for h := rect.Min.X; h < rect.Max.X; h++ {
		img.Set(h, rect.Min.Y, red)
		img.Set(h, rect.Max.Y-1, red)
	}
	// 左右の枠
	for v := rect.Min.Y; v < rect.Max.Y; v++ {
		img.Set(rect.Min.X, v, red)
		img.Set(rect.Max.X-1, v, red)

	}
	return img
}
