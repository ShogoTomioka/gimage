/*
　画像データを二値化する、もしくは二値化したデータに対する処理をまとめる
　二値化する際は、Grayスケールの値が128より大きいものを255(白)、128以下を0(黒)にしている

*/

package main

import (
	"image"
	"image/color"
)

// 二値化した画像のデータを返す関数
func Binarization(imgObject image.Image) *image.Gray {

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

func NoBinarization(imgObject image.Image) *image.Gray {

	rec := imgObject.Bounds()
	binary := image.NewGray(rec)

	// グレー化したものSetして返却
	for v := rec.Min.Y; v < rec.Max.Y; v++ {
		for h := rec.Min.X; h < rec.Max.X; h++ {
			c := color.GrayModel.Convert(imgObject.At(h, v))
			gray, _ := c.(color.Gray)
			binary.Set(h, v, gray)
		}
	}
	return binary
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

/*Erosion(縮小)処理をするための関数*/
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
