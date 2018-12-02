package gimage

import (
	"image"
	"image/color"
)

//WatchArea は、指定された範囲内の明るさが大きければTrueを、そうでなければFalseを返す
func WatchArea(g *image.Gray, width int, height int, p image.Point) bool {

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
	if count > (width*height)/THREAHOLD {
		return true
	} else {
		return false
	}
}

//精査された二値データから明るいか(True)、暗いか(True)の情報が入った配列を作成
func ScanImage(g *image.Gray) [][]bool {
	var lists [][]bool
	var t bool
	var point image.Point

	// 全体を100分割する
	size := g.Rect.Size()
	width := size.X / DIVISION
	height := size.Y / DIVISION

	for v := 0; v < DIVISION; v++ {
		var list []bool
		for h := 0; h < DIVISION; h++ {
			point = image.Point{X: width * h, Y: height * v}
			t = WatchArea(g, width, height, point)
			list = append(list, t)
		}
		lists = append(lists, list)
	}
	return lists
}

//Erosion(縮小)処理をするための関数
func ErosionImage(g *image.Gray) *image.Gray {
	ResultBinary := image.NewGray(g.Rect)
	gray := color.Gray{Y: 255}

	for v := g.Rect.Min.Y; v < g.Rect.Max.Y; v++ {
		for h := g.Rect.Min.X; h < g.Rect.Max.X; h++ {
			g1 := g.GrayAt(h-1, v).Y == 255
			g2 := g.GrayAt(h+1, v).Y == 255
			g3 := g.GrayAt(h, v-1).Y == 255
			g4 := g.GrayAt(h, v+1).Y == 255
			g5 := g.GrayAt(h-1, v-1).Y == 255
			g6 := g.GrayAt(h-1, v+1).Y == 255
			g7 := g.GrayAt(h+1, v-1).Y == 255
			g8 := g.GrayAt(h+1, v+1).Y == 255

			if g1 && g2 && g3 && g4 && g5 && g6 && g7 && g8 {
				ResultBinary.SetGray(h, v, gray)
			}
		}
	}
	return ResultBinary
}

//Dilation(拡張)処理をするための関数
func DilationImage(g *image.Gray) *image.Gray {
	ResultBinary := image.NewGray(g.Rect)
	gray := color.Gray{Y: 255}

	for v := g.Rect.Min.Y; v < g.Rect.Max.Y; v++ {
		for h := g.Rect.Min.X; h < g.Rect.Max.X; h++ {
			if g.GrayAt(h, v).Y == 255 {
				ResultBinary.SetGray(h-1, v, gray)
				ResultBinary.SetGray(h+1, v, gray)
				ResultBinary.SetGray(h, v-1, gray)
				ResultBinary.SetGray(h, v+1, gray)
				ResultBinary.SetGray(h-1, v-1, gray)
				ResultBinary.SetGray(h-1, v+1, gray)
				ResultBinary.SetGray(h+1, v-1, gray)
				ResultBinary.SetGray(h+1, v+1, gray)
			}

		}
	}
	return ResultBinary
}

func GrayDiff(g1 *image.Gray, g2 *image.Gray) *image.Gray {

	diffBinary := image.NewGray(g1.Rect)
	gray := color.Gray{Y: 255}

	for v := g1.Rect.Min.Y; v < g1.Rect.Max.Y; v++ {
		for h := g1.Rect.Min.X; h < g1.Rect.Max.X; h++ {
			if g1.GrayAt(h, v) != g2.GrayAt(h, v) {
				diffBinary.SetGray(h, v, gray)
			}
		}
	}
	return diffBinary
}
