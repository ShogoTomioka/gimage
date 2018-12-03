package main

import (
	"image/png"
	"os"

	"github.com/ShogoTomioka/gimage"
)

func main() {

	const (
		// 比較する二つの画像へのパス
		FilePathA = "./pictures/picture_A.png"
		FilePathB = "./pictures/picture_B.png"
	)
	// Thresholdは、フィルターをかける時の閾値、大きいと小さい差分でも検出し、小さいと大きいものしか検出しない
	// Divisionは画像を分割する単位、大きいと粗く、小さいと細かく処理を行う
	filter := &gimage.Filter{Threshold: 10, Division: 20}

	//比較する画像データを読み込む
	imageA, _ := gimage.NewColorImage(FilePathA)
	imageB, _ := gimage.NewColorImage(FilePathB)

	gray := &gimage.Gray{}
	//二つの比較する画像をそれぞれグレースケール化する
	gray.Graying(imageA, imageB)

	//グレー化した画像を比較し、二値画像を作成する
	gray.GrayDiff()

	//DilationとErotionを"3回"繰り返し、画像を鳴らす
	gray.Convert(3)

	//二値画像から明るい部分がTrue、暗い部分がFalseになった二次元配列を獲得
	filter.ScanImage(gray.Image)

	filteredImg := filter.OverlaidFilter(imageA)

	filterFile, _ := os.Create("filtered.png")
	defer filterFile.Close()

	if err := png.Encode(filterFile, filteredImg); err != nil {
		panic(err)
	}

}
