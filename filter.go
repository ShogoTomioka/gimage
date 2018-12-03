package gimage

import (
	"image"
	"image/color"
)

//Filter 二値画像に対するフィルタリング処理に関する構造体
type Filter struct {
	Image     *image.NRGBA
	Threshold int
	Division  int
	Lists     [][]bool
}

//WatchArea は、指定された範囲内の明るさが大きければTrueを、そうでなければFalseを返す
func (f Filter) WatchArea(g *image.Gray, width int, height int, p image.Point) bool {

	x := p.X
	y := p.Y

	var count = 0
	gray := color.Gray{Y: 255}

	for v := 0; v < width; v++ {
		for h := 0; h < height; h++ {
			if g.GrayAt(x+h, y+v) == gray {
				count++
			}
		}
	}
	var result bool
	if count > (width*height)/f.Threshold {
		result = true
	} else {
		result = false
	}
	return result
}

// OverlaidFilter は間違いのあるところを赤っぽくする
func (f Filter) OverlaidFilter(srcImg image.Image) *image.NRGBA {

	//元画像からレクタングルを取得
	rec := srcImg.Bounds()
	width := rec.Max.X
	height := rec.Max.Y

	//各ボックスの辺の長さをDIVISIONから求める
	boxWidth := width / f.Division
	boxHeight := height / f.Division

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	img = fillColor(img, srcImg, width, height)
	for h := 0; h < f.Division; h++ {
		for w := 0; w < f.Division; w++ {
			if f.Lists[h][w] == true {
				x := boxWidth * w
				y := boxHeight * h
				for i := y; i < y+boxHeight; i++ {
					for t := x; t < x+boxWidth; t++ {
						_, G, B, _ := srcImg.At(t, i).RGBA()
						img.Set(t, i, color.RGBA{
							R: uint8(255),
							G: uint8(G),
							B: uint8(B),
							A: 255,
						})
					}
				}
			}
		}
	}
	return img
}

//ScanImage は精査された二値データから明るいか(True)、暗いか(True)の情報が入った配列を作成
func (f *Filter) ScanImage(g *image.Gray) {
	var lists [][]bool
	var t bool
	var point image.Point

	// 全体を100分割する
	size := g.Rect.Size()
	width := size.X / f.Division
	height := size.Y / f.Division

	for v := 0; v < f.Division; v++ {
		var list []bool
		for h := 0; h < f.Division; h++ {
			point = image.Point{X: width * h, Y: height * v}
			t = f.WatchArea(g, width, height, point)
			list = append(list, t)
		}
		lists = append(lists, list)
	}
	f.Lists = lists
}
