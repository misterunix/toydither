package imgproc

import (
	"image"
	"image/color"
	"math"
)

// converts an image to black and white using the BT.601 standard
//
// imgIn is passed as a pointer for speed, but is not modified.
func BWBT601(imgIn *image.RGBA) *image.RGBA {

	imgOut := image.NewRGBA(imgIn.Bounds())

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	for y := range hh {
		for x := range ww {
			r, g, b, _ := imgIn.At(x, y).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8

			gg := (float32(r) * 0.299) +
				(float32(g) * 0.587) +
				(float32(b) * 0.114)

			gray := uint8(gg)

			imgOut.Set(x, y, color.Gray{gray})
		}
	}

	return imgOut
}

// converts an image to black and white using the Luminance method with
// gamma correction and BT.601
//
// imgIn is passed as a pointer for speed, but is not modified.
func BWLuminanceGammaBT601(imgIn *image.RGBA) *image.RGBA {
	imgOut := image.NewRGBA(imgIn.Bounds())

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	for y := range hh {
		for x := range ww {
			r, g, b, _ := imgIn.At(x, y).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8

			rL := float64(r) / 255.0
			gL := float64(g) / 255.0
			bL := float64(b) / 255.0

			rG := math.Pow(float64(rL), 2.2)
			gG := math.Pow(float64(gL), 2.2)
			bG := math.Pow(float64(bL), 2.2)

			rgN := rG * 255
			ggN := gG * 255
			bgN := bG * 255
			// Luminance coefficients
			gg := (float32(rgN) * 0.299) +
				(float32(ggN) * 0.587) +
				(float32(bgN) * 0.114)

			// Gamma correction
			// gg = (gg / 255.0)
			// gg = (gg * gg * gg)

			gray := uint8(gg)

			imgOut.Set(x, y, color.Gray{gray})
		}
	}

	return imgOut
}

// converts an image to black and white using the BT.709 standard
//
// imgIn is passed as a pointer for speed, but is not modified.
func BWBT709(imgIn *image.RGBA) *image.RGBA {
	imgOut := image.NewRGBA(imgIn.Bounds())

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	for y := range hh {
		for x := range ww {
			r, g, b, _ := imgIn.At(x, y).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8

			gg := (float32(r) * 0.2126) +
				(float32(g) * 0.7152) +
				(float32(b) * 0.0722)

			gray := uint8(gg)

			imgOut.Set(x, y, color.Gray{gray})
		}
	}

	return imgOut
}

// converts an image to black and white using the Luminance method with
// gamma correction and BT.709
//
// imgIn is passed as a pointer for speed, but is not modified.
func BWLuminanceGammaBT709(imgIn *image.RGBA) *image.RGBA {
	imgOut := image.NewRGBA(imgIn.Bounds())

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	for y := range hh {
		for x := range ww {
			r, g, b, _ := imgIn.At(x, y).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8

			rL := float64(r) / 255.0
			gL := float64(g) / 255.0
			bL := float64(b) / 255.0

			rG := math.Pow(float64(rL), 2.2)
			gG := math.Pow(float64(gL), 2.2)
			bG := math.Pow(float64(bL), 2.2)

			rgN := rG * 255
			ggN := gG * 255
			bgN := bG * 255

			gg := (float32(rgN) * 0.2126) +
				(float32(ggN) * 0.7152) +
				(float32(bgN) * 0.0722)

			gray := uint8(gg)

			imgOut.Set(x, y, color.Gray{gray})
		}
	}

	return imgOut
}

// converts an image to black and white using the BT.2020 standard
//
// imgIn is passed as a pointer for speed, but is not modified.
func BWBT2020(imgIn *image.RGBA) *image.RGBA {
	imgOut := image.NewRGBA(imgIn.Bounds())

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	for y := range hh {
		for x := range ww {
			r, g, b, _ := imgIn.At(x, y).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8

			// Luma coefficients

			gg := (float32(r) * 0.2627) +
				(float32(g) * 0.6780) +
				(float32(b) * 0.0593)

			gray := uint8(gg)

			imgOut.Set(x, y, color.Gray{gray})
		}
	}

	return imgOut

}

// converts an image to black and white using the Luminance method with
// gamma correction and BT.2020
//
// imgIn is passed as a pointer for speed, but is not modified.
func BWLuminanceGammaBT2020(imgIn *image.RGBA) *image.RGBA {
	imgOut := image.NewRGBA(imgIn.Bounds())

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	for y := range hh {
		for x := range ww {
			r, g, b, _ := imgIn.At(x, y).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8

			rL := float64(r) / 255.0
			gL := float64(g) / 255.0
			bL := float64(b) / 255.0

			rG := math.Pow(float64(rL), 2.2)
			gG := math.Pow(float64(gL), 2.2)
			bG := math.Pow(float64(bL), 2.2)

			rgN := rG * 255
			ggN := gG * 255
			bgN := bG * 255
			// Luminance coefficients
			gg := (float32(rgN) * 0.2627) +
				(float32(ggN) * 0.6780) +
				(float32(bgN) * 0.0593)

			// Gamma correction
			// gg = (gg / 255.0)
			// gg = (gg * gg * gg)

			gray := uint8(gg)

			imgOut.Set(x, y, color.Gray{gray})
		}
	}

	return imgOut
}

// converts an image to black and white using the Luminance method
// (ITU-R BT.709 coefficients)
//
// imgIn is passed as a pointer for speed, but is not modified.
func BWLuminance(imgIn *image.RGBA) *image.RGBA {
	imgOut := image.NewRGBA(imgIn.Bounds())

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	for y := range hh {
		for x := range ww {
			r, g, b, _ := imgIn.At(x, y).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8

			// Luminance coefficients
			gg := (float32(r) * 0.2126) +
				(float32(g) * 0.7152) +
				(float32(b) * 0.0722)

			gray := uint8(gg)

			imgOut.Set(x, y, color.Gray{gray})
		}
	}

	return imgOut
}

//
//
//

// converts an image to black and white using the Luminance method with
// gamma correction. This method uses the BT.709 coeff
//
// imgIn is passed as a pointer for speed, but is not modified.
func BWLuminanceGamma(imgIn *image.RGBA) *image.RGBA {
	imgOut := image.NewRGBA(imgIn.Bounds())

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	for y := range hh {
		for x := range ww {
			r, g, b, _ := imgIn.At(x, y).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8

			rL := float64(r) / 255.0
			gL := float64(g) / 255.0
			bL := float64(b) / 255.0

			rG := math.Pow(float64(rL), 2.2)
			gG := math.Pow(float64(gL), 2.2)
			bG := math.Pow(float64(bL), 2.2)

			rgN := rG * 255
			ggN := gG * 255
			bgN := bG * 255

			gg := (float32(rgN) * 0.2126) +
				(float32(ggN) * 0.7152) +
				(float32(bgN) * 0.0722)

			// Gamma correction
			// gg = (gg / 255.0)
			// gg = (gg * gg * gg)

			gray := uint8(gg)

			imgOut.Set(x, y, color.Gray{gray})
		}
	}

	return imgOut
}

// converts an image to black and white using the simple average method
// (r+g+b)/3
//
// imgIn is passed as a pointer for speed, but is not modified.
func BWSimple(imgIn *image.RGBA) *image.RGBA {
	imgOut := image.NewRGBA(imgIn.Bounds())

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	for y := range hh {
		for x := range ww {
			r, g, b, _ := imgIn.At(x, y).RGBA()
			rf := float64(uint8(r >> 8))
			bf := float64(uint8(b >> 8))
			gf := float64(uint8(g >> 8))
			grayf := (rf + gf + bf) / 3.0 // float64 simple average

			gray := uint8(grayf)

			imgOut.Set(x, y, color.Gray{gray})
		}
	}
	return imgOut

}

// convert an image using ratios for RGB values. It is up to the caller to ensure
// that the sum of the ratios equals 1.0
//
// I made this to keep from having to write the same code over and over again for red, green or blue.
//
// imgIn is passed as a pointer for speed, but is not modified.
func BWRGBRatio(imgIn *image.RGBA, rRatio float64, gRatio float64, bRatio float64) *image.RGBA {

	imgOut := image.NewRGBA(imgIn.Bounds())

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	for y := range hh {
		for x := range ww {
			r, g, b, _ := imgIn.At(x, y).RGBA()
			rf := float64(uint8(r >> 8))
			bf := float64(uint8(b >> 8))
			gf := float64(uint8(g >> 8))
			grayf := (rf*rRatio + gf*gRatio + bf*bRatio) // float64 simple average

			gray := uint8(grayf)

			imgOut.Set(x, y, color.Gray{gray})
		}
	}
	return imgOut
}
