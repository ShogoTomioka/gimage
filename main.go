package main

import (
	"image/png"
	"os"

	"github.com/ShogoTomioka/go-images/lib"
)

func main() {

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
	}

	filteredImg := lib.OverlaidFilter(imageA, lists)

	filterFile, _ := os.Create("filtered.png")
	defer filterFile.Close()

	if err := png.Encode(filterFile, filteredImg); err != nil {
		panic(err)
	}

}
