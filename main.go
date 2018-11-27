package main

import (
	"flag"
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

func GrayDiff(g1 *image.Gray, g2 *image.Gray) *image.Gray {
	//本来は画像サイズが違うことを考慮しないといけないが今回は同じサイズの画像に限るので(~_~;)

	diffBinary := image.NewGray(g1.Rect)
	gray := color.Gray{Y: 255}

	for v := g1.Rect.Min.Y; v < g1.Rect.Max.Y; v++ {
		for h := g1.Rect.Min.X; h < g1.Rect.Max.X; h++ {
			if g1.At(h, v) == g2.At(h, v) {
				diffBinary.Set(h, v, gray)
			}
		}
	}
	return diffBinary
}

func main() {

	var (
		filePath_A string
		filePath_B string
		outPath    string
	)

	flag.StringVar(&filePath_A, "i", "pictures/picture_A.png", "input file name 1")
	flag.StringVar(&filePath_B, "f", "pictures/picture_B.png", "input file name 2")
	flag.StringVar(&outPath, "o", "outfile.png", "output file name")
	flag.Parse()

	imageA := GenerateImage(filePath_A)
	imageB := GenerateImage(filePath_B)

	grayImage := Binarization(imageA)
	grayImageB := Binarization(imageB)
	/*
		// 書き出し用ファイル準備
		outfile, _ := os.Create("out.png")
		defer outfile.Close()
		// 書き出し
		png.Encode(outfile, grayImage)

		outfileB, _ := os.Create("outB.png")
		defer outfileB.Close()
	*/
	// 差分を求めて書き出し
	diffImage := GrayDiff(grayImage, grayImageB)

	outfile, _ := os.Create(outPath)
	defer outfile.Close()
	png.Encode(outfile, diffImage)

}
