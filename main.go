package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"os"
)

func GenerateImage(filename string) image.Image {

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

// 二値化した画像のデータを返す関数
func Binarization(imgObject image.Image) *image.Gray {

	// imageデータをグレースケール化したものを作成
	rec := imgObject.Bounds()

	binary := image.NewGray(rec)

	// グレーイメージに対して二値化処理
	for v := rec.Min.Y; v < rec.Max.Y; v++ {
		for h := rec.Min.X; h < rec.Max.X; h++ {
			c := color.GrayModel.Convert(imgObject.At(h, v))
			gray, _ := c.(color.Gray)
			// しきい値で二値化
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

// 二値化ではなくグレースケールで比較
func NoBinarization(imgObject image.Image) *image.Gray {

	// imageデータをグレースケール化したものを作成
	rec := imgObject.Bounds()

	binary := image.NewGray(rec)

	// グレーイメージに対して二値化処理
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
	//本来は画像サイズが違うことを考慮しないといけないが今回は同じサイズの画像に限るので(~_~;)

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

func GenerateFileter(g *image.Gray) *image.RGBA {

	fileter := image.NewRGBA(g.Rect)

	// 全体を100分割する
	//width := (g.Rect.Max.X - g.Rect.Min.X) / 10
	//height := (g.Rect.Max.Y - g.Rect.Min.Y) / 10

	var lists [][]bool

	for v := 0; v < 10; v++ {
		var list []bool

		for h := 0; h < 10; h++ {
			var t = true

			rec := image.Rect{}

			list = append(list, t)
		}
		lists = append(lists, list)
	}
	return fileter
}

func DetectDiff() {}

// 枠線を描く
func drawBounds(img *image.RGBA, col color.Color, rect image.Rectangle) {
	// 矩形を取得
	//rect := img.Rect

	for h := 0; h < rect.Max.X; h++ {
		img.Set(h, 0, col)
		img.Set(h, rect.Max.Y-1, col)
	}
	for v := 0; v < rect.Max.Y; v++ {
		img.Set(0, v, col)
		img.Set(rect.Max.X-1, v, col)
	}
}

func main() {

	var (
		FILE_PATH_A string
		FILE_PATH_B string
		OUT_PATH    string
	)

	flag.StringVar(&FILE_PATH_A, "i", "pictures/picture_A.png", "input file name 1")
	flag.StringVar(&FILE_PATH_B, "f", "pictures/picture_B.png", "input file name 2")
	flag.StringVar(&OUT_PATH, "o", "outfile.png", "output file name")

	flag.Parse()

	imageA := GenerateImage(FILE_PATH_A)
	imageB := GenerateImage(FILE_PATH_B)

	grayImage := NoBinarization(imageA)
	grayImageB := NoBinarization(imageB)

	// 差分を求めて書き出し
	diffImage := GrayDiff(grayImage, grayImageB)

	diffImage = ErosionImage(diffImage)

	diffImage = DilationImage(diffImage)

	for i := 0; i < 7; i++ {
		diffImage = ErosionImage(diffImage)
	}
	for i := 0; i < 7; i++ {
		diffImage = DilationImage(diffImage)
	}

	_ = GenerateFileter(diffImage)
	outfile, _ := os.Create(OUT_PATH)
	defer outfile.Close()
	png.Encode(outfile, diffImage)

}
