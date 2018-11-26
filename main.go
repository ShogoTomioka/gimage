package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

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

type Picture struct {
	file    *os.File
	Rec     image.Rectangle
	ImgGray *image.Gray
}

func main() {

	filePath_A := "./pictures/picture_A.png"
	//filePAth_B := "./pictures/picture_B.png"

	//画像ファイルのオープン
	file, _ := os.Open(filePath_A)
	defer file.Close()

	//ファイルをデコードしてImageオブジェクトを作成
	imageObj, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	//	srcBounds := imageObj.Bounds()

	// 出力用イメージ
	grayImage := Binarization(imageObj)

	// 書き出し用ファイル準備
	outfile, _ := os.Create("out.png")
	defer outfile.Close()
	// 書き出し
	png.Encode(outfile, grayImage)
}
