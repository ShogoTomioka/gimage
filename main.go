package main

import (
	"flag"
	"image/png"
	"os"

	gimage "github.com/ShogoTomioka/gimage/lib"
)

func main() {

	const (
		FilePathA1 = "./testdata/picture_A.png"
		FilePathA2 = "./testdata/picture_B.png"
		FilePathB1 = "./testdata/picture_C.png"
		FilePathB2 = "./testdata/picture_D.png"
	)

	var pathA string
	var pathB string
	flag.Parse()
	arg := flag.Args()[0]

	switch arg {
	case "A":
		pathA = FilePathA1
		pathB = FilePathA2
	case "B":
		pathA = FilePathB1
		pathB = FilePathB2
	}

	filter := &gimage.Filter{Threshold: 10, Division: 20}

	imageA, _ := gimage.NewColorImage(pathA)
	imageB, _ := gimage.NewColorImage(pathB)

	gray := &gimage.Gray{}

	gray.Graying(imageA, imageB)

	gray.GrayDiff()

	//DilationとErotionを"3回"繰り返し
	gray.Convert(3)

	filter.ScanImage(gray.Image)

	//元の画像に間違い部分をフィルターする
	filteredImg := filter.OverlaidFilter(imageA)

	filterFile, _ := os.Create("./testdata/filtered.png")
	defer filterFile.Close()

	if err := png.Encode(filterFile, filteredImg); err != nil {
		panic(err)
	}

	binFile, _ := os.Create("./testdata/binary.png")
	defer binFile.Close()

	if err := png.Encode(binFile, gray.Image); err != nil {
		panic(err)
	}

}
