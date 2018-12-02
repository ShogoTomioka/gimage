package gimage

import (
	"image"
	"image/color"
)

type Filter struct {
	Threshold int
	Division  int
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
	if count > (width*height)/f.Threshold {
		return true
	} else {
		return false
	}
}

func (f Filter) OverlaidFilter(srcImg image.Image, lists [][]bool) *image.NRGBA {

	//元画像からレクタングルを取得
	rec := srcImg.Bounds()
	width := rec.Max.X
	height := rec.Max.Y

	//各ボックスの辺の長さをDIVISIONから求める
	box_width := width / f.Division
	box_height := height / f.Division

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	img = fillColor(img, srcImg, width, height)
	for h := 0; h < f.Division; h++ {
		for w := 0; w < f.Division; w++ {
			if lists[h][w] == true {
				x := box_width * w
				y := box_height * h
				for i := y; i < y+box_height; i++ {
					for t := x; t < x+box_width; t++ {
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

//精査された二値データから明るいか(True)、暗いか(True)の情報が入った配列を作成
func (f Filter) ScanImage(g *image.Gray) [][]bool {
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
	return lists
}

//GenerateFilter は二値画像の明るい部分に枠線を描く
func (f Filter) GenerateFilter(g *image.Gray) (*image.RGBA, [][]bool) {

	//間違っている部分を示すためのフィルターImage
	filter := image.NewRGBA(g.Rect)
	scanList := f.ScanImage(g)

	size := g.Rect.Size()
	width := size.X / f.Division
	height := size.Y / f.Division

	//フィルターの枠線のRectangleを作成、あとでFor分でこのRectangleをずらしていく
	min := image.Point{X: 0, Y: 0}
	point := image.Point{X: width, Y: height}
	rec := image.Rectangle{Min: min, Max: point}

	for i := 0; i < f.Division; i++ {
		for t := 0; t < f.Division; t++ {
			if scanList[t][i] == true {
				p := image.Point{X: width * t, Y: height * i}
				redRec := rec.Add(p)

				filter = DrawBound(filter, redRec)
			}
		}
	}
	return filter, scanList
}
