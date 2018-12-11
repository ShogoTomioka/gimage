package gimage

import (
	"image"
	"image/color"
	"os"
)

//NewColorImage は指定されたパスの画像ファイルを開いて、画像イメージを返却する
func NewColorImage(path string) (image.Image, error) {

	//画像ファイルのオープン
	file, _ := os.Open(path)
	defer file.Close()

	//ファイルをデコードしてImageオブジェクトを作成
	imageObj, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	return imageObj, err
}

//元画像をimgに描写して返却する。
func fillColor(img *image.NRGBA, srcImg image.Image, width int, height int) *image.NRGBA {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			R, G, B, _ := srcImg.At(x, y).RGBA()

			img.Set(x, y, color.NRGBA{
				R: uint8(R),
				G: uint8(G),
				B: uint8(B),
				A: 255,
			})
		}
	}
	return img
}
