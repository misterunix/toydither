package main

import (
	_ "embed"
	"flag"
	"fmt"
	"image"
	"lzrimg/internal/imgproc"
)

//go:embed embeds/w3.css
var w3embed []byte

//go:embed embeds/slide.css
var slideembed []byte

//go:embed embeds/dropdown.css
var dropdownembed []byte

//go:embed embeds/index.html
var indexembed []byte

//go:embed embeds/htmx.min.js
var htmxembed []byte

//go:embed embeds/formupload.html
var formuploadembed []byte

//go:embed embeds/lzr.css
var lzrembed []byte

var srcImageFilename string
var dstImageFilename string

var srcImage *image.RGBA

func main() {
	var err error
	flag.StringVar(&srcImageFilename, "i", "", "Source Image")
	flag.StringVar(&dstImageFilename, "o", "", "Destination Image")
	flag.Parse()

	if srcImageFilename == "" || dstImageFilename == "" {
		flag.PrintDefaults()
		return
	}

	// load the image
	srcImage, err = imgproc.LoadImageFromFile(srcImageFilename)
	if err != nil {
		fmt.Println("Error loading source image:", err)
	}
	server()
}
