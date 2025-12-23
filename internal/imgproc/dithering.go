package imgproc

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand/v2"
)

// Helper function to calculate the error diffusion center around the pixel
func errorDiffustion(img *image.RGBA, errorTerm, errorMul float64, x, y int) {
	SetPixel(img, x, y, GetPixel(img, x, y)+int(errorMul*errorTerm))
}

// Arkinson dithering implementation
//
// imgIn is passed as a pointer for speed, but is not modified.
func DitherArkinson(imgIn *image.RGBA) *image.RGBA {

	// use draw to create, copy the input image to a new image
	// and then modify that image in place
	imgMod := image.NewRGBA(imgIn.Bounds())
	draw.Draw(imgMod, imgMod.Bounds(), imgIn, imgIn.Bounds().Min, draw.Src)

	// because of the way look ahead dithering works, it needs to
	// modify te same image as it is reading from

	ww := imgMod.Bounds().Dx()
	hh := imgMod.Bounds().Dy()

	for y := range hh {
		for x := range ww {

			// keep everything in float as long as possible

			var diffErr float64

			pixel := GetPixel(imgMod, x, y)
			if pixel >= 127 {
				diffErr = float64(DiffGray(255, pixel))
			} else {
				diffErr = float64(DiffGray(0, pixel))
			}

			diffErr = (float64(diffErr) / 8.0)

			// I think this is cleaner, if not slower
			errorDiffustion(imgMod, diffErr, 1.0, x+1, y)
			errorDiffustion(imgMod, diffErr, 1.0, x+2, y)
			errorDiffustion(imgMod, diffErr, 1.0, x+1, y+1)
			errorDiffustion(imgMod, diffErr, 1.0, x+2, y+1)
			errorDiffustion(imgMod, diffErr, 1.0, x, y+1)
			errorDiffustion(imgMod, diffErr, 1.0, x, y+2)
			errorDiffustion(imgMod, diffErr, 1.0, x-1, y+1)

			if pixel >= 127 {
				SetPixel(imgMod, x, y, 255)
			} else {
				SetPixel(imgMod, x, y, 0)
			}

		}
	}

	return imgMod

}

// Burkes dithering implementation
//
// imgIn is passed as a pointer for speed, but is not modified.
func DitherBurkes(imgIn *image.RGBA) *image.RGBA {

	// use draw to create, copy the input image to a new image
	// and then modify that image in place
	imgMod := image.NewRGBA(imgIn.Bounds())
	draw.Draw(imgMod, imgMod.Bounds(), imgIn, imgIn.Bounds().Min, draw.Src)

	//imgMod := image.NewRGBA(imgIn.Bounds())

	// because of the way look ahead dithering works, it needs to
	// modify te same image as it is reading from

	ww := imgMod.Bounds().Dx()
	hh := imgMod.Bounds().Dy()

	for y := range hh {
		for x := range ww {

			// keep everything in float as long as possible

			var diffErr float64

			pixel := GetPixel(imgMod, x, y)
			if pixel >= 127 {
				diffErr = float64(DiffGray(255, pixel))
			} else {
				diffErr = float64(DiffGray(0, pixel))
			}

			diffErr = (float64(diffErr) / 32.0)

			// I think this is cleaner, if not slower
			errorDiffustion(imgMod, diffErr, 8.0, x+1, y)
			errorDiffustion(imgMod, diffErr, 4.0, x+2, y)
			errorDiffustion(imgMod, diffErr, 2.0, x-2, y+1)
			errorDiffustion(imgMod, diffErr, 4.0, x-1, y+1)
			errorDiffustion(imgMod, diffErr, 8.0, x, y+1)
			errorDiffustion(imgMod, diffErr, 4.0, x+1, y+1)
			errorDiffustion(imgMod, diffErr, 2.0, x+2, y+1)

			if pixel >= 127 {
				SetPixel(imgMod, x, y, 255)
			} else {
				SetPixel(imgMod, x, y, 0)
			}

		}
	}

	return imgMod

}

// False Floyd-Steinberg dithering implementation
//
// imgIn is passed as a pointer for speed, but is not modified.
func DitherFalseFloydSteinberg(imgIn *image.RGBA) *image.RGBA {

	// use draw to create, copy the input image to a new image
	// and then modify that image in place
	imgMod := image.NewRGBA(imgIn.Bounds())
	draw.Draw(imgMod, imgMod.Bounds(), imgIn, imgIn.Bounds().Min, draw.Src)

	//imgMod := image.NewRGBA(imgIn.Bounds())

	// because of the way look ahead dithering works, it needs to
	// modify te same image as it is reading from

	ww := imgMod.Bounds().Dx()
	hh := imgMod.Bounds().Dy()

	for y := range hh {
		for x := range ww {

			// keep everything in float as long as possible

			var diffErr float64

			pixel := GetPixel(imgMod, x, y)
			if pixel >= 127 {
				diffErr = float64(DiffGray(255, pixel))
			} else {
				diffErr = float64(DiffGray(0, pixel))
			}

			diffErr = (float64(diffErr) / 8.0)

			// I think this is cleaner, if not slower
			errorDiffustion(imgMod, diffErr, 3.0, x+1, y)
			errorDiffustion(imgMod, diffErr, 2.0, x-1, y+1)
			errorDiffustion(imgMod, diffErr, 3.0, x, y+1)

			if pixel >= 127 {
				SetPixel(imgMod, x, y, 255)
			} else {
				SetPixel(imgMod, x, y, 0)
			}

		}
	}

	return imgMod

}

//
//
//

// dither using the Floyd-Steinberg implementation
func DitherFloydSteinberg(imgIn *image.RGBA) *image.RGBA {

	imgMod := image.NewRGBA(imgIn.Bounds())
	draw.Draw(imgMod, imgMod.Bounds(), imgIn, imgIn.Bounds().Min, draw.Src)

	// because of the way look ahead dithering works, it needs to
	// modify te same image as it is reading from

	ww := imgMod.Bounds().Dx()
	hh := imgMod.Bounds().Dy()

	for y := range hh {
		for x := range ww {

			// keep everything in float as long as possible

			var diffErr float64

			pixel := GetPixel(imgMod, x, y)
			if pixel >= 127 {
				diffErr = float64(DiffGray(255, pixel))
			} else {
				diffErr = float64(DiffGray(0, pixel))
			}

			diffErr = (float64(diffErr) / 16.0)

			// this is cleaner, if not slower
			errorDiffustion(imgMod, diffErr, 7.0, x+1, y)
			errorDiffustion(imgMod, diffErr, 3.0, x-1, y+1)
			errorDiffustion(imgMod, diffErr, 5.0, x, y+1)
			errorDiffustion(imgMod, diffErr, 1.0, x+1, y+1)

			if pixel >= 127 {
				SetPixel(imgMod, x, y, 255)
			} else {
				SetPixel(imgMod, x, y, 0)
			}

		}
	}

	return imgMod

}

// Jarvis-Judice-Ninke dithering implementation
//
// imgIn is passed as a pointer for speed, but is not modified.
func DitherJarvisJudiceNinke(imgIn *image.RGBA) *image.RGBA {

	// use draw to create, copy the input image to a new image
	// and then modify that image in place
	imgMod := image.NewRGBA(imgIn.Bounds())
	draw.Draw(imgMod, imgMod.Bounds(), imgIn, imgIn.Bounds().Min, draw.Src)

	//imgMod := image.NewRGBA(imgIn.Bounds())

	// because of the way look ahead dithering works, it needs to
	// modify te same image as it is reading from

	ww := imgMod.Bounds().Dx()
	hh := imgMod.Bounds().Dy()

	for y := range hh {
		for x := range ww {

			// keep everything in float as long as possible

			var diffErr float64

			pixel := GetPixel(imgMod, x, y)
			if pixel >= 127 {
				diffErr = float64(DiffGray(255, pixel))
			} else {
				diffErr = float64(DiffGray(0, pixel))
			}

			diffErr = (float64(diffErr) / 48.0)

			// I think this is cleaner, if not slower
			errorDiffustion(imgMod, diffErr, 7.0, x+1, y)
			errorDiffustion(imgMod, diffErr, 5.0, x+2, y)
			errorDiffustion(imgMod, diffErr, 3.0, x-2, y+1)
			errorDiffustion(imgMod, diffErr, 5.0, x-1, y+1)
			errorDiffustion(imgMod, diffErr, 7.0, x, y+1)
			errorDiffustion(imgMod, diffErr, 5.0, x+1, y+1)
			errorDiffustion(imgMod, diffErr, 3.0, x+2, y+1)
			errorDiffustion(imgMod, diffErr, 1.0, x-2, y+2)
			errorDiffustion(imgMod, diffErr, 3.0, x-1, y+2)
			errorDiffustion(imgMod, diffErr, 5.0, x, y+2)
			errorDiffustion(imgMod, diffErr, 3.0, x+1, y+2)
			errorDiffustion(imgMod, diffErr, 1.0, x+2, y+2)

			if pixel >= 127 {
				SetPixel(imgMod, x, y, 255)
			} else {
				SetPixel(imgMod, x, y, 0)
			}

		}
	}

	return imgMod

}

// dither an image using random probability. Produces ugly images.
//
// imgIn is passed as a pointer for speed, but is not modified.
func DitherRandom(imgIn *image.RGBA) *image.RGBA {

	// use draw to create, copy the input image to a new image
	// and then modify that image in place
	imgMod := image.NewRGBA(imgIn.Bounds())
	draw.Draw(imgMod, imgMod.Bounds(), imgIn, imgIn.Bounds().Min, draw.Src)

	//imgMod := image.NewRGBA(imgIn.Bounds())

	ww := imgMod.Bounds().Dx()
	hh := imgMod.Bounds().Dy()

	var gray uint8
	for y := range hh {
		for x := range ww {
			r, _, _, _ := imgMod.At(x, y).RGBA()
			rn := uint8(rand.IntN(255))
			rv := uint8(r >> 8)

			if rv > rn {
				gray = 255
			} else {
				gray = 0
			}
			SetPixel(imgMod, x, y, int(gray))
			//imgIn.Set(x, y, color.Gray{gray})
		}
	}

	return imgMod

}

// DitherSierra1 performs a Sierra 1 dithering algorithm on the input image.
//
// imgIn is passed as a pointer for speed, but is not modified.
func DitherSierra1(imgIn *image.RGBA) *image.RGBA {

	// use draw to create, copy the input image to a new image
	// and then modify that image in place
	imgMod := image.NewRGBA(imgIn.Bounds())
	draw.Draw(imgMod, imgMod.Bounds(), imgIn, imgIn.Bounds().Min, draw.Src)

	//    imgMod := image.NewRGBA(imgIn.Bounds())

	// because of the way look ahead dithering works, it needs to
	// modify te same image as it is reading from

	ww := imgMod.Bounds().Dx()
	hh := imgMod.Bounds().Dy()

	for y := range hh {
		for x := range ww {

			// keep everything in float as long as possible

			var diffErr float64

			pixel := GetPixel(imgMod, x, y)
			if pixel >= 127 {
				diffErr = float64(DiffGray(255, pixel))
			} else {
				diffErr = float64(DiffGray(0, pixel))
			}

			diffErr = (float64(diffErr) / 32.0)

			// I think this is cleaner, if not slower
			errorDiffustion(imgMod, diffErr, 5.0, x+1, y)
			errorDiffustion(imgMod, diffErr, 3.0, x+2, y)
			errorDiffustion(imgMod, diffErr, 2.0, x-2, y+1)
			errorDiffustion(imgMod, diffErr, 4.0, x-1, y+1)
			errorDiffustion(imgMod, diffErr, 5.0, x, y+1)
			errorDiffustion(imgMod, diffErr, 4.0, x+1, y+1)
			errorDiffustion(imgMod, diffErr, 2.0, x+2, y+1)
			errorDiffustion(imgMod, diffErr, 2.0, x-1, y+2)
			errorDiffustion(imgMod, diffErr, 3.0, x, y+2)
			errorDiffustion(imgMod, diffErr, 2.0, x+1, y+2)

			if pixel >= 127 {
				SetPixel(imgMod, x, y, 255)
			} else {
				SetPixel(imgMod, x, y, 0)
			}

		}
	}

	return imgMod

}

//
//
//

func DitherSierra2(imgIn *image.RGBA) *image.RGBA {

	imgMod := image.NewRGBA(imgIn.Bounds())
	draw.Draw(imgMod, imgMod.Bounds(), imgIn, imgIn.Bounds().Min, draw.Src)

	// because of the way look ahead dithering works, it needs to
	// modify te same image as it is reading from

	ww := imgMod.Bounds().Dx()
	hh := imgMod.Bounds().Dy()

	for y := range hh {
		for x := range ww {

			// keep everything in float as long as possible

			var diffErr float64

			pixel := GetPixel(imgMod, x, y)
			if pixel >= 127 {
				diffErr = float64(DiffGray(255, pixel))
			} else {
				diffErr = float64(DiffGray(0, pixel))
			}

			diffErr = (float64(diffErr) / 16.0)

			// I think this is cleaner, if not slower
			errorDiffustion(imgMod, diffErr, 4.0, x+1, y)
			errorDiffustion(imgMod, diffErr, 3.0, x+2, y)
			errorDiffustion(imgMod, diffErr, 1.0, x-2, y+1)
			errorDiffustion(imgMod, diffErr, 2.0, x-1, y+1)
			errorDiffustion(imgMod, diffErr, 3.0, x, y+1)
			errorDiffustion(imgMod, diffErr, 2.0, x+1, y+1)
			errorDiffustion(imgMod, diffErr, 1.0, x+2, y+1)

			if pixel >= 127 {
				SetPixel(imgMod, x, y, 255)
			} else {
				SetPixel(imgMod, x, y, 0)
			}

		}
	}

	return imgMod

}

//
//
//

func DitherSierra2_4a(imgIn *image.RGBA) *image.RGBA {

	imgMod := image.NewRGBA(imgIn.Bounds())
	draw.Draw(imgMod, imgMod.Bounds(), imgIn, imgIn.Bounds().Min, draw.Src)

	// because of the way look ahead dithering works, it needs to
	// modify te same image as it is reading from

	ww := imgMod.Bounds().Dx()
	hh := imgMod.Bounds().Dy()

	for y := range hh {
		for x := range ww {

			// keep everything in float as long as possible

			var diffErr float64

			pixel := GetPixel(imgMod, x, y)
			if pixel >= 127 {
				diffErr = float64(DiffGray(255, pixel))
			} else {
				diffErr = float64(DiffGray(0, pixel))
			}

			diffErr = (float64(diffErr) / 4.0)

			// I think this is cleaner, if not slower

			errorDiffustion(imgMod, diffErr, 2.0, x+1, y)
			errorDiffustion(imgMod, diffErr, 1.0, x-1, y+1)
			errorDiffustion(imgMod, diffErr, 1.0, x, y+1)

			if pixel >= 127 {
				SetPixel(imgMod, x, y, 255)
			} else {
				SetPixel(imgMod, x, y, 0)
			}

		}
	}

	return imgMod

}

// dither an image using only two states. Produces ugly images.
// value <= 127 = 0, values > 127 = 255
//
// imgIn is passed as a pointer for speed, but is not modified.
func DitherSimple(imgIn *image.RGBA) *image.RGBA {

	// use draw to create, copy the input image to a new image
	// and then modify that image in place
	imgMod := image.NewRGBA(imgIn.Bounds())
	draw.Draw(imgMod, imgMod.Bounds(), imgIn, imgIn.Bounds().Min, draw.Src)

	//imgOut := image.NewRGBA(imgIn.Bounds())

	//	imgMod := image.NewRGBA(imgIn.Bounds())

	ww := imgMod.Bounds().Dx()
	hh := imgMod.Bounds().Dy()

	var gray uint8
	for y := range hh {
		for x := range ww {
			r, _, _, _ := imgMod.At(x, y).RGBA()
			rv := uint8(r >> 8)

			if rv > 127 {
				gray = 255
			} else {
				gray = 0
			}

			imgMod.Set(x, y, color.Gray{gray})

		}
	}

	return imgMod

}

// DitherStucki applies the Stucki dithering algorithm to an image.
//
// imgIn is passed as a pointer for speed, but is not modified.
func DitherStucki(imgIn *image.RGBA) *image.RGBA {

	// use draw to create, copy the input image to a new image
	// and then modify that image in place
	imgMod := image.NewRGBA(imgIn.Bounds())
	draw.Draw(imgMod, imgMod.Bounds(), imgIn, imgIn.Bounds().Min, draw.Src)

	// because of the way look ahead dithering works, it needs to
	// modify te same image as it is reading from

	//	imgMod := image.NewRGBA(imgIn.Bounds())

	ww := imgMod.Bounds().Dx()
	hh := imgMod.Bounds().Dy()

	for y := range hh {
		for x := range ww {

			// keep everything in float as long as possible

			var diffErr float64

			pixel := GetPixel(imgMod, x, y)
			if pixel >= 127 {
				diffErr = float64(DiffGray(255, pixel))
			} else {
				diffErr = float64(DiffGray(0, pixel))
			}

			diffErr = (float64(diffErr) / 32.0)

			// I think this is cleaner, if not slower
			errorDiffustion(imgMod, diffErr, 8.0, x+1, y)
			errorDiffustion(imgMod, diffErr, 4.0, x+2, y)
			errorDiffustion(imgMod, diffErr, 2.0, x-2, y+1)
			errorDiffustion(imgMod, diffErr, 4.0, x-1, y+1)
			errorDiffustion(imgMod, diffErr, 8.0, x, y+1)
			errorDiffustion(imgMod, diffErr, 4.0, x+1, y+1)
			errorDiffustion(imgMod, diffErr, 2.0, x+2, y+1)
			errorDiffustion(imgMod, diffErr, 1.0, x-2, y+2)
			errorDiffustion(imgMod, diffErr, 2.0, x-1, y+2)
			errorDiffustion(imgMod, diffErr, 4.0, x, y+2)
			errorDiffustion(imgMod, diffErr, 2.0, x+1, y+2)
			errorDiffustion(imgMod, diffErr, 1.0, x+2, y+2)

			if pixel >= 127 {
				SetPixel(imgMod, x, y, 255)
			} else {
				SetPixel(imgMod, x, y, 0)
			}

		}
	}

	return imgMod

}

// Unknown-1 dithering implementation
//
// imgIn is passed as a pointer for speed, but is not modified.
func DitherUnknown1(imgIn *image.RGBA) *image.RGBA {

	// use draw to create, copy the input image to a new image
	// and then modify that image in place
	imgMod := image.NewRGBA(imgIn.Bounds())
	draw.Draw(imgMod, imgMod.Bounds(), imgIn, imgIn.Bounds().Min, draw.Src)

	//imgMod := image.NewRGBA(imgIn.Bounds())

	// because of the way look ahead dithering works, it needs to
	// modify te same image as it is reading from

	ww := imgMod.Bounds().Dx()
	hh := imgMod.Bounds().Dy()

	for y := range hh {
		for x := range ww {

			// keep everything in float as long as possible

			var diffErr float64

			pixel := GetPixel(imgMod, x, y)
			if pixel >= 127 {
				diffErr = float64(DiffGray(255, pixel))
			} else {
				diffErr = float64(DiffGray(0, pixel))
			}

			diffErr = (float64(diffErr) / 12.0)

			// I think this is cleaner, if not slower
			errorDiffustion(imgMod, diffErr, 2.0, x+2, y)
			errorDiffustion(imgMod, diffErr, 1.0, x+1, y)
			errorDiffustion(imgMod, diffErr, 2.0, x-1, y+1)
			errorDiffustion(imgMod, diffErr, 2.0, x, y+1)
			errorDiffustion(imgMod, diffErr, 2.0, x+1, y+1)
			errorDiffustion(imgMod, diffErr, 1.0, x-2, y+2)
			errorDiffustion(imgMod, diffErr, 1.0, x, y+2)
			errorDiffustion(imgMod, diffErr, 1.0, x+2, y+2)

			if pixel >= 127 {
				SetPixel(imgMod, x, y, 255)
			} else {
				SetPixel(imgMod, x, y, 0)
			}

		}
	}

	return imgMod

}
