package main

import (
	_ "embed"
	"flag"
	"lzrimg/imgproc"
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

var srcImageFilename string
var dstImageFilename string

func main() {

	flag.StringVar(&srcImageFilename, "i", "", "Source Image")
	flag.StringVar(&dstImageFilename, "o", "", "Destination Image")
	flag.Parse()

	if srcImageFilename == "" || dstImageFilename == "" {
		flag.PrintDefaults()
		return
	}

	// load the image
	srcImage := imgproc.LoadImageFromFile(srcImageFilename)

}
