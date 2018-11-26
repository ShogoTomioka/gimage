package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"reflect"
)

type picture struct {
	File *image.NRGBA
	start (int,int)
}

//画像データをグレー化させる


//

func main() {
	// 画像ファイルを開く(書き込み元)
	src, _ := os.Open("./pictures/picture_A.png")
	defer src.Close()

	// デコードしてイメージオブジェクトを準備
	srcImg, _, err := image.Decode(src)
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(srcImg))
	srcBounds := srcImg.Bounds()

	// 出力用イメージ
	dest := image.NewGray(srcBounds)

	// グレー化
	for v := srcBounds.Min.Y; v < srcBounds.Max.Y; v++ {
		for h := srcBounds.Min.X; h < srcBounds.Max.X; h++ {
			c := color.GrayModel.Convert(srcImg.At(h, v))
			gray, _ := c.(color.Gray)
			dest.Set(h, v, gray)
		}
	}

	// 書き出し用ファイル準備
	outfile, _ := os.Create("out.png")
	defer outfile.Close()
	// 書き出し
	png.Encode(outfile, dest)
}
