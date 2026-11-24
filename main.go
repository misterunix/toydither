package main

import (
	"embed"
	"encoding/base64"
	"fmt"
	"html/template"
	"image/png"
	"io"
	"log"
	"lzrimg/imgproc"
	"net/http"
	"strconv"

	"bytes"
	"image"
)

//go:embed templates/index.html
var content embed.FS

var (
	currentImage    *image.RGBA
	processedImage  *image.RGBA
	currentFilename string
)

type PageData struct {
	ImageLoaded bool
	Filename    string
	ImageData   string
	Width       int
	Height      int
}

func webserver() {

	tmpl := template.Must(template.ParseFS(content, "templates/*.html"))

	// Main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("web /")
		data := PageData{
			ImageLoaded: currentImage != nil,
			Filename:    currentFilename,
		}
		if currentImage != nil {
			data.Width = currentImage.Bounds().Dx()
			data.Height = currentImage.Bounds().Dy()

			// Encode current image to base64
			var buf bytes.Buffer
			png.Encode(&buf, currentImage)
			data.ImageData = base64.StdEncoding.EncodeToString(buf.Bytes())
		}
		tmpl.ExecuteTemplate(w, "index.html", data)
	})

	// Upload image
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		log.Println("web /upload")
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Error reading file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Read file into memory
		data, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Error reading file", http.StatusBadRequest)
			return
		}

		// Decode image
		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			http.Error(w, "Error decoding image", http.StatusBadRequest)
			return
		}

		// Convert to RGBA
		currentImage = imgproc.ConvertToRGBA(img)
		processedImage = imgproc.ConvertToRGBA(img)
		currentFilename = header.Filename

		// Return image immediately
		sendImageResponse(w, processedImage)
	})

	// Apply B&W conversion
	http.HandleFunc("/apply-bw", func(w http.ResponseWriter, r *http.Request) {
		log.Println("web /apply-bw")
		if currentImage == nil {
			http.Error(w, "No image loaded", http.StatusBadRequest)
			return
		}

		method := r.FormValue("bw-method")

		var result *image.RGBA
		switch method {
		case "BT.601":
			result = imgproc.BWBT601(currentImage)
		case "BT.709":
			result = imgproc.BWBT709(currentImage)
		case "BT.2020":
			result = imgproc.BWBT2020(currentImage)
		case "Luminance":
			result = imgproc.BWLuminance(currentImage)
		case "LuminanceGamma":
			result = imgproc.BWLuminanceGamma(currentImage)
		case "LuminanceGammaBT601":
			result = imgproc.BWLuminanceGammaBT601(currentImage)
		case "LuminanceGammaBT709":
			result = imgproc.BWLuminanceGammaBT709(currentImage)
		case "LuminanceGammaBT2020":
			result = imgproc.BWLuminanceGammaBT2020(currentImage)
		case "Simple":
			result = imgproc.BWSimple(currentImage)
		default:
			result = imgproc.BWBT601(currentImage)
		}

		processedImage = result
		sendImageResponse(w, result)
	})

	// Apply dithering
	http.HandleFunc("/apply-dither", func(w http.ResponseWriter, r *http.Request) {
		log.Println("web /apply-dither")
		if processedImage == nil {
			http.Error(w, "No image to dither", http.StatusBadRequest)
			return
		}

		method := r.FormValue("dither-method")

		if method == "None" {
			sendImageResponse(w, processedImage)
			return
		}

		var result *image.RGBA
		switch method {
		case "Floyd-Steinberg":
			result = imgproc.DitherFloydSteinberg(processedImage)
		case "False Floyd-Steinberg":
			result = imgproc.DitherFalseFloydSteinberg(processedImage)
		case "Atkinson":
			result = imgproc.DitherArkinson(processedImage)
		case "Burkes":
			result = imgproc.DitherBurkes(processedImage)
		case "Simple":
			result = imgproc.DitherSimple(processedImage)
		default:
			result = imgproc.DitherFloydSteinberg(processedImage)
		}

		processedImage = result
		sendImageResponse(w, result)
	})

	// Process all
	http.HandleFunc("/process-all", func(w http.ResponseWriter, r *http.Request) {
		log.Println("web /process-all")
		if currentImage == nil {
			http.Error(w, "No image loaded", http.StatusBadRequest)
			return
		}

		// Get parameters
		bwMethod := r.FormValue("bw-method")
		ditherMethod := r.FormValue("dither-method")

		// Apply B&W
		var result *image.RGBA
		switch bwMethod {
		case "BT.601":
			result = imgproc.BWBT601(currentImage)
		case "BT.709":
			result = imgproc.BWBT709(currentImage)
		case "BT.2020":
			result = imgproc.BWBT2020(currentImage)
		case "Luminance":
			result = imgproc.BWLuminance(currentImage)
		case "LuminanceGamma":
			result = imgproc.BWLuminanceGamma(currentImage)
		case "LuminanceGammaBT601":
			result = imgproc.BWLuminanceGammaBT601(currentImage)
		case "LuminanceGammaBT709":
			result = imgproc.BWLuminanceGammaBT709(currentImage)
		case "LuminanceGammaBT2020":
			result = imgproc.BWLuminanceGammaBT2020(currentImage)
		case "Simple":
			result = imgproc.BWSimple(currentImage)
		default:
			result = imgproc.BWBT601(currentImage)
		}

		// Apply dithering if not "None"
		if ditherMethod != "None" {
			switch ditherMethod {
			case "Floyd-Steinberg":
				result = imgproc.DitherFloydSteinberg(result)
			case "False Floyd-Steinberg":
				result = imgproc.DitherFalseFloydSteinberg(result)
			case "DitherJarvisJudiceNinke":
				result = imgproc.DitherJarvisJudiceNinke(result)
			case "Sierra3":
				result = imgproc.DitherSierra3(result)
			case "Sierra2":
				result = imgproc.DitherSierra2(result)
			case "SierraLite":
				result = imgproc.DitherSierraLite(result)
			case "Stucki":
				result = imgproc.DitherStucki(result)
			case "Atkinson":
				result = imgproc.DitherArkinson(result)
			case "Burkes":
				result = imgproc.DitherBurkes(result)
			case "Simple":
				result = imgproc.DitherSimple(result)
			}
		}

		processedImage = result
		sendImageResponse(w, result)
	})

	// Reset image
	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		log.Println("web /reset")
		if currentImage != nil {
			processedImage = imgproc.ConvertToRGBA(currentImage)
			sendImageResponse(w, processedImage)
		}
	})

	// Download processed image
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		log.Println("web /download")
		if processedImage == nil {
			http.Error(w, "No processed image", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Content-Disposition", `attachment; filename="processed.png"`)

		png.Encode(w, processedImage)
	})

	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		log.Println("web /result")
		if processedImage == nil {

			http.Error(w, "No processed image", http.StatusBadRequest)
			return
		}
		sendImageResponse(w, processedImage)
		//fmt.Fprint(w, "Hello from /result")
	})

	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func main() {

	webserver()

}

func sendImageResponse(w http.ResponseWriter, img *image.RGBA) {
	// Encode to PNG
	var buf bytes.Buffer
	png.Encode(&buf, img)

	// Convert to base64
	b64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	// Send image and dimensions
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
		<img src="data:image/png;base64,%s" class="w3-image" style="max-width:100%%; height:auto;" id="result-img">
		<div class="w3-margin-top">
			<p><strong>Dimensions:</strong> %dx%d pixels</p>
		</div>
	`, b64, img.Bounds().Dx(), img.Bounds().Dy())
}

func parseFloat(s string, def float64) float64 {
	if v, err := strconv.ParseFloat(s, 64); err == nil {
		return v
	}
	return def
}

func parseInt(s string, def int) int {
	if v, err := strconv.Atoi(s); err == nil {
		return v
	}
	return def
}
