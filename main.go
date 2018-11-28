package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

const (
	DIVISION = 10
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

func GenerateFileter(g *image.Gray) *image.RGBA {

	//間違っている部分を示すためのフィルターImage
	fileter := image.NewRGBA(g.Rect)
	scanList := ScanImage(g)

	size := g.Rect.Size()
	width := size.X / DIVISION
	height := size.Y / DIVISION

	//フィルターの枠線のRectangleを作成、あとでFor分でこのRectangleをずらしていく
	min := image.Point{X: 0, Y: 0}
	point := image.Point{X: width, Y: height}
	rec := image.Rectangle{Min: min, Max: point}

	for i := 0; i > DIVISION; i++ {
		for t := 0; t < DIVISION; t++ {
			if scanList[t][i] == true {

				//scanListでTrueになっている部分のRectangleに枠線を描写
				fileter = DrawBound(image)
			}
		}
	}

	return fileter
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

	for v := 0; v < 10; v++ {
		var list []bool
		for h := 0; h < 10; h++ {
			point = image.Point{X: width * h, Y: height * v}
			t = WatchArea(g, width, height, point)
			list = append(list, t)
		}
		lists = append(lists, list)
	}
	fmt.Println(lists)
	return lists
}

//指定された範囲内の明るさが大きければTrueを、そうでなければFalseを返す
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
	if count > (width*height)/10 {
		return true
	} else {
		return false
	}
}

func DrawBound(img *image.RGBA, rect image.Rectangle) *image.RGBA {

	//間違っている部分を囲う枠線の色
	red := color.RGBA{255, 0, 0, 0}
	//rectの範囲に枠線を書く
	// 上下の枠
	for h := 0; h < rect.Max.X; h++ {
		img.Set(h, 0, red)
		img.Set(h, rect.Max.Y-1, red)
	}
	// 左右の枠
	for v := 0; v < rect.Max.Y; v++ {
		img.Set(0, v, red)
		img.Set(rect.Max.X-1, v, red)
	}
	return img

}

// 枠線を描く
func drawBounds(img *image.RGBA, col color.Color, rect image.Rectangle) *image.RGBA {

	for h := 0; h < rect.Max.X; h++ {
		img.Set(h, 0, col)
		img.Set(h, rect.Max.Y-1, col)
	}
	for v := 0; v < rect.Max.Y; v++ {
		img.Set(0, v, col)
		img.Set(rect.Max.X-1, v, col)
	}
	return img
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

	for i := 0; i < 5; i++ {
		diffImage = ErosionImage(diffImage)
	}
	for i := 0; i < 5; i++ {
		diffImage = DilationImage(diffImage)
	}

	_ = GenerateFileter(diffImage)

	//出力用のイメージを用意
	outRect := image.Rectangle{image.Pt(0, 0), imageA.Bounds().Size()}
	out := image.NewRGBA(outRect)

	//元画像に対して作成したフィルターを上書き
	RectA := image.Rectangle{image.Pt(0, 0), imageA.Bounds().Size()}
	draw.Draw(out, RectA, imageA, image.Pt(0, 0), draw.Src)

	//フィルターを元画像に対して上書きする
	RectB := image.Rectangle{image.Pt(0, 0), imageB.Bounds().Size()}
	draw.Draw(out, RectB, imageB, image.Pt(0, 0), draw.Over)

	outfile, _ := os.Create(OUT_PATH)
	defer outfile.Close()
	//png.Encode(outfile, diffImage)
	png.Encode(outfile, out)

}
