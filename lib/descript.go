package lib

import (
	"image"
	"image/color"
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

//画像データを二値化して返却する
func Binarization(imgObject image.Image) *image.Gray {

	rec := imgObject.Bounds()
	binary := image.NewGray(rec)

	// グレーイメージに対して二値化処理
	for v := rec.Min.Y; v < rec.Max.Y; v++ {
		for h := rec.Min.X; h < rec.Max.X; h++ {
			c := color.GrayModel.Convert(imgObject.At(h, v))
			gray, _ := c.(color.Gray)
			// しきい値(128)で二値化
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

//画像イメージをグレー化して返却
func NoBinarization(imgObject image.Image) *image.Gray {

	rec := imgObject.Bounds()
	binary := image.NewGray(rec)

	// グレー化したものSetして返却
	for v := rec.Min.Y; v < rec.Max.Y; v++ {
		for h := rec.Min.X; h < rec.Max.X; h++ {
			c := color.GrayModel.Convert(imgObject.At(h, v))
			gray, _ := c.(color.Gray)
			binary.Set(h, v, gray)
		}
	}
	return binary
}

//元画像をimgに描写して返却する。
//要するに元画像を書くだけ
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
func OverlaidFilter(srcImg image.Image, lists [][]bool) *image.NRGBA {

	//元画像からレクタングルを取得
	rec := srcImg.Bounds()
	width := rec.Max.X
	height := rec.Max.Y

	//各ボックスの辺の長さをDIVISIONから求める
	box_width := width / DIVISION
	box_height := height / DIVISION

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	img = fillColor(img, srcImg, width, height)
	for h := 0; h < DIVISION; h++ {
		for w := 0; w < DIVISION; w++ {
			if lists[h][w] == true {
				x := box_width * w
				y := box_height * h
				for i := y; i < y+box_height; i++ {
					for t := x; t < x+box_width; t++ {
						_, G, B, _ := srcImg.At(t, i).RGBA()
						img.Set(t, i, color.RGBA{
							R: uint8(255),
							G: uint8(G),
							B: uint8(B),
							A: 255,
						})
					}
				}
			}
		}
	}
	return img
}
