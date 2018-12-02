package main

import (
<<<<<<< HEAD
=======
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
>>>>>>> master
	"image/png"
	"os"

<<<<<<< HEAD
	"github.com/ShogoTomioka/go-images/lib"
)

func main() {
=======
const (
	DIVISION  = 10
	THREAHOLD = 10
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

//GenerateFilter は二値画像の明るい部分に枠線を描く
func GenerateFilter(g *image.Gray) *image.RGBA {

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

	fmt.Println(rec)
	for i := 0; i < DIVISION; i++ {
		for t := 0; t < DIVISION; t++ {
			if scanList[t][i] == true {
				p := image.Point{X: width * t, Y: height * i}
				redRec := rec.Add(p)
				fmt.Println(redRec)
				//scanListでTrueになっている部分のRectangleに枠線を描写
				//	filter.Set(i, t, red)
				filter = DrawBound(filter, redRec)
			}
		}
	}
	return filter
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

//WatchArea は、指定された範囲内の明るさが大きければTrueを、そうでなければFalseを返す
func WatchArea(g *image.Gray, width int, height int, p image.Point) bool {
>>>>>>> master

	const (
		FILE_PATH_A = "./pictures/picture_A.png"
		FILE_PATH_B = "./pictures/picture_B.png"
	)

	//比較する画像データを読み込む
	imageA := lib.GenerateImage(FILE_PATH_A)
	imageB := lib.GenerateImage(FILE_PATH_B)

<<<<<<< HEAD
	//二つの比較する画像をそれぞれグレースケール化する
	grayImage := lib.NoBinarization(imageA)
	grayImageB := lib.NoBinarization(imageB)

	//グレー化した画像を比較し、二値画像を作成する
	diffImage := lib.GrayDiff(grayImage, grayImageB)

	//DilationとErotionを繰り返し、画像を鳴らす
	for i := 0; i < 3; i++ {
		diffImage = lib.ErosionImage(diffImage)
	}
	for i := 0; i < 3; i++ {
		diffImage = lib.DilationImage(diffImage)
	}

	//二値画像から明るい部分がTrue、暗い部分がFalseになった二次元配列を獲得
	lists := lib.ScanImage(diffImage)

	file, _ := os.Create("sample.png")
	defer file.Close()

	if err := png.Encode(file, diffImage); err != nil {
		panic(err)
=======
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

>>>>>>> master
	}

	filteredImg := lib.OverlaidFilter(imageA, lists)

	filterFile, _ := os.Create("filtered.png")
	defer filterFile.Close()

<<<<<<< HEAD
	if err := png.Encode(filterFile, filteredImg); err != nil {
		panic(err)
	}

=======
	imageA := GenerateImage(FILE_PATH_A)
	imageB := GenerateImage(FILE_PATH_B)

	grayImage := NoBinarization(imageA)
	grayImageB := NoBinarization(imageB)

	// 差分を求めて書き出し
	diffImage := GrayDiff(grayImage, grayImageB)

	for i := 0; i < 5; i++ {
		diffImage = ErosionImage(diffImage)
	}
	for i := 0; i < 5; i++ {
		diffImage = DilationImage(diffImage)
	}

	filter := GenerateFilter(diffImage)
	DrawBound(filter, filter.Bounds())

	//出力用のイメージを用意
	outRect := image.Rectangle{image.Pt(0, 0), imageA.Bounds().Size()}
	out := image.NewRGBA(outRect)

	//フィルターを元画像に対して上書きする
	RectB := image.Rectangle{image.Pt(0, 0), imageA.Bounds().Size()}
	draw.Draw(out, RectB, imageA, image.Pt(0, 0), draw.Over)

	outfile, _ := os.Create(OUT_PATH)
	defer outfile.Close()
	//	png.Encode(outfile, diffImage)
	png.Encode(outfile, diffImage)

	// 出力用ファイル作成(エラー処理は略)
	file, _ := os.Create("sample.jpeg")
	defer file.Close()

	// JPEGで出力(100%品質)
	if err := jpeg.Encode(file, filter, &jpeg.Options{100}); err != nil {
		panic(err)
	}
>>>>>>> master
}
