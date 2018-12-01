package lib

import (
	"image"
	"image/color"
)

const (
	DIVISION  = 20
	THREAHOLD = 15
)

//GenerateFilter は二値画像の明るい部分に枠線を描く
func GenerateFilter(g *image.Gray) (*image.RGBA, [][]bool) {

	//間違っている部分を示すためのフィルターImage
	filter := image.NewRGBA(g.Rect)
	scanList := ScanImage(g)

	size := g.Rect.Size()
	width := size.X / DIVISION
	height := size.Y / DIVISION

	//フィルターの枠線のRectangleを作成、あとでFor分でこのRectangleをずらしていく
	min := image.Point{X: 0, Y: 0}
	point := image.Point{X: width, Y: height}
	rec := image.Rectangle{Min: min, Max: point}

	for i := 0; i < DIVISION; i++ {
		for t := 0; t < DIVISION; t++ {
			if scanList[t][i] == true {
				p := image.Point{X: width * t, Y: height * i}
				redRec := rec.Add(p)
				//scanListでTrueになっている部分のRectangleに枠線を描写
				//	filter.Set(i, t, red)
				filter = DrawBound(filter, redRec)
			}
		}
	}
	return filter, scanList
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
