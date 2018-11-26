package main

import (
	"flag"
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

func GrayDiff(g1 *image.Gray, g2 *image.Gray) {
	//本来は画像サイズが違うことを考慮しないといけない

}

func main() {

	var (
		filePath_A string
		filePath_B string
		outPath    string
	)
	flag.StringVar(&filePath_A, "i", "", "input file name 1")
	flag.StringVar(&filePath_B, "f", "", "input file name 2")
	flag.StringVar(&outPath, "o", "outfile.png", "output file name")
	flag.Parse()

	//画像ファイルのオープン
	file_A, _ := os.Open(filePath_A)
	defer file_A.Close()
	file_B, _ := os.Open(filePath_B)
	defer file_B.Close()

	//ファイルをデコードしてImageオブジェクトを作成
	imageObj, _, err := image.Decode(file_A)
	if err != nil {
		panic(err)
	}

	imageObjB, _, err := image.Decode(file_B)
	if err != nil {
		panic(err)
	}

	//ここで画像のサイズが違う可能性を考慮する

	grayImage := Binarization(imageObj)
	grayImageB := Binarization(imageObjB)

	// 書き出し用ファイル準備
	outfile, _ := os.Create("out.png")
	defer outfile.Close()
	// 書き出し
	png.Encode(outfile, grayImage)

	outfileB, _ := os.Create("outB.png")
	defer outfileB.Close()
	// 書き出し
	png.Encode(outfileB, grayImageB)

	GrayDiff(grayImage, grayImageB)

}
