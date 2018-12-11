package gimage

import (
	"image"
	"image/color"
)

//Gray は二値に関するメソッドを持った構造体
type Gray struct {
	ImageA *image.Gray
	ImageB *image.Gray
	Image  *image.Gray
	Rec    image.Rectangle
}

//Graying は画像イメージをグレー化して返却
func (g *Gray) Graying(imgObjectA image.Image, imgObjectB image.Image) {

	g.Rec = imgObjectA.Bounds()
	g.ImageA = image.NewGray(g.Rec)
	g.ImageB = image.NewGray(g.Rec)

	// グレー化したものSetして返却
	for v := g.Rec.Min.Y; v < g.Rec.Max.Y; v++ {
		for h := g.Rec.Min.X; h < g.Rec.Max.X; h++ {
			ca := color.GrayModel.Convert(imgObjectA.At(h, v))
			cb := color.GrayModel.Convert(imgObjectB.At(h, v))
			grayA, _ := ca.(color.Gray)
			grayB, _ := cb.(color.Gray)
			g.ImageA.Set(h, v, grayA)
			g.ImageB.Set(h, v, grayB)

		}
	}
}

//Convert は指定した回数分ErotionとDilationを行う
func (g *Gray) Convert(times int) {
	g.ErosionImage(times)
	g.DilationImage(times)
}

//ErosionImage は縮小処理(Erossion)をするための関数
func (g *Gray) ErosionImage(times int) {

	for i := 0; i < times; i++ {
		resImage := image.NewGray(g.Rec)
		gray := color.Gray{Y: 255}
		for v := g.Rec.Min.Y; v < g.Rec.Max.Y; v++ {
			for h := g.Rec.Min.X; h < g.Rec.Max.X; h++ {
				g1 := g.Image.GrayAt(h-1, v).Y == 255
				g2 := g.Image.GrayAt(h+1, v).Y == 255
				g3 := g.Image.GrayAt(h, v-1).Y == 255
				g4 := g.Image.GrayAt(h, v+1).Y == 255
				g5 := g.Image.GrayAt(h-1, v-1).Y == 255
				g6 := g.Image.GrayAt(h-1, v+1).Y == 255
				g7 := g.Image.GrayAt(h+1, v-1).Y == 255
				g8 := g.Image.GrayAt(h+1, v+1).Y == 255

				//周りのピクセルが全部白の場合のみ明るくする
				if g1 && g2 && g3 && g4 && g5 && g6 && g7 && g8 {
					resImage.SetGray(h, v, gray)
				}
			}
		}
		g.Image = resImage
	}

}

//DilationImage は拡張処理(Dilation)をするための関数
func (g *Gray) DilationImage(times int) {

	gray := color.Gray{Y: 255}
	for i := 0; i < times; i++ {
		resImage := image.NewGray(g.Rec)
		for v := g.Rec.Min.Y; v < g.Rec.Max.Y; v++ {
			for h := g.Rec.Min.X; h < g.Rec.Max.X; h++ {
				//その点が白なら周りも全部白にする
				if g.Image.GrayAt(h, v).Y == 255 {
					resImage.SetGray(h, v, gray)
					resImage.SetGray(h-1, v, gray)
					resImage.SetGray(h+1, v, gray)
					resImage.SetGray(h, v-1, gray)
					resImage.SetGray(h, v+1, gray)
					resImage.SetGray(h-1, v-1, gray)
					resImage.SetGray(h-1, v+1, gray)
					resImage.SetGray(h+1, v-1, gray)
					resImage.SetGray(h+1, v+1, gray)
				}
			}
		}
		g.Image = resImage
	}
}

//GrayDiff は二つのグレー画像から差分をとった二値画像を作成する
func (g *Gray) GrayDiff() {

	diffBinary := image.NewGray(g.Rec)
	gray := color.Gray{Y: 255}

	for v := g.Rec.Min.Y; v < g.Rec.Max.Y; v++ {
		for h := g.Rec.Min.X; h < g.Rec.Max.X; h++ {
			if g.ImageA.GrayAt(h, v) != g.ImageB.GrayAt(h, v) {
				diffBinary.SetGray(h, v, gray)
			}
		}
	}
	g.Image = diffBinary
}
