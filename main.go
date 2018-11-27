package main

import (
	"flag"
	"fmt"
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

func GenerateFileter(g *image.Gray) *image.RGBA {

	fileter := image.NewRGBA(g.Rect)
	_ = ScanImage(g)

	return fileter
}

//精査された二値データから明るいか(True)、暗いか(True)の情報が入った配列を作成
func ScanImage(g *image.Gray) [][]bool {
	var lists [][]bool
	var t bool
	var point image.Point
	// 全体を100分割する
	width := (g.Rect.Max.X - g.Rect.Min.X) / 10
	height := (g.Rect.Max.Y - g.Rect.Min.Y) / 10

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

	red := color.RGBA{255, 0, 0, 0}
	//rectの範囲に枠線を書く
	return img
}

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

	for i := 0; i < 5; i++ {
		diffImage = ErosionImage(diffImage)
	}
	for i := 0; i < 5; i++ {
		diffImage = DilationImage(diffImage)
	}

	_ = GenerateFileter(diffImage)
	outfile, _ := os.Create(OUT_PATH)
	defer outfile.Close()
	png.Encode(outfile, diffImage)

}
