package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/ShogoTomioka/go-images/lib"
)

func main() {

	height := size.Y / DIVISION

	const (
		FILE_PATH_A = "./pictures/picture_A.png"
		FILE_PATH_B = "./pictures/picture_B.png"
	)

	//比較する画像データを読み込む
	imageA := lib.GenerateImage(FILE_PATH_A)
	imageB := lib.GenerateImage(FILE_PATH_B)

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

	filteredImg := lib.OverlaidFilter(imageA, lists)

	filterFile, _ := os.Create("filtered.png")
	defer filterFile.Close()

	if err := png.Encode(filterFile, filteredImg); err != nil {
		panic(err)
	}

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
}
