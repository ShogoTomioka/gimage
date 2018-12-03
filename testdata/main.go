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

	filter := &gimage.Filter{Threshold: 10, Division: 20}

	//比較する画像データを読み込む
	imageA := gimage.NewColorImage(FILE_PATH_A)
	imageB := gimage.NewColorImage(FILE_PATH_B)

	gray := &gimage.Gray{}
	//二つの比較する画像をそれぞれグレースケール化する
	gray.Graying(imageA, imageB)

	//グレー化した画像を比較し、二値画像を作成する
	gray.GrayDiff()

	//DilationとErotionを"3回"繰り返し、画像を鳴らす
	gray.ErosionImage(3)
	gray.DilationImage(3)
	fi, _ := os.Create("test.png")
	defer fi.Close()

	if err := png.Encode(fi, gray.Image); err != nil {
		panic(err)
	}
	//二値画像から明るい部分がTrue、暗い部分がFalseになった二次元配列を獲得
	filter.ScanImage(gray.Image)

	filteredImg := filter.OverlaidFilter(imageA)

	filterFile, _ := os.Create("filtered.png")
	defer filterFile.Close()

	if err := png.Encode(filterFile, filteredImg); err != nil {
		panic(err)
	}

}
