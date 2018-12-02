package main

import (
	"image/png"
	"os"

	"github.com/ShogoTomioka/gimage"
)

func main() {

	const (
		FILE_PATH_A = "./pictures/picture_A.png"
		FILE_PATH_B = "./pictures/picture_B.png"
	)

	filter := gimage.Filter{Threshold: 10, Division: 10}

	//比較する画像データを読み込む
	imageA := gimage.NewColorImage(FILE_PATH_A)
	imageB := gimage.NewColorImage(FILE_PATH_B)

	//二つの比較する画像をそれぞれグレースケール化する
	grayImage := gimage.Graying(imageA)
	grayImageB := gimage.Graying(imageB)

	//グレー化した画像を比較し、二値画像を作成する
	diffImage := gimage.GrayDiff(grayImage, grayImageB)

	file1, _ := os.Create("test.png")
	defer file1.Close()

	if err := png.Encode(file1, grayImageB); err != nil {
		panic(err)
	}
	//DilationとErotionを繰り返し、画像を鳴らす
	for i := 0; i < 3; i++ {
		diffImage = gimage.ErosionImage(diffImage)
	}
	for i := 0; i < 3; i++ {
		diffImage = gimage.DilationImage(diffImage)
	}

	//二値画像から明るい部分がTrue、暗い部分がFalseになった二次元配列を獲得
	filter.ScanImage(diffImage)

	file, _ := os.Create("sample.png")
	defer file.Close()

	if err := png.Encode(file, diffImage); err != nil {
		panic(err)
	}

	filteredImg := filter.OverlaidFilter(imageA)

	filterFile, _ := os.Create("filtered.png")
	defer filterFile.Close()

	if err := png.Encode(filterFile, filteredImg); err != nil {
		panic(err)
	}

}
