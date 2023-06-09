package cmd

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func generateLifeBar(totalHearts int, emptyPercentage float64, outputFile string) error {
	filledHearts := int((1 - emptyPercentage) * float64(totalHearts))

	heartWidth := 20
	heartHeight := 20
	barWidth := totalHearts * heartWidth
	barHeight := heartHeight

	// Create a new RGBA image
	img := image.NewRGBA(image.Rect(0, 0, barWidth, barHeight))

	// Set the background color
	backgroundColor := color.RGBA{255, 255, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{backgroundColor}, image.Point{}, draw.Src)

	// Draw filled hearts
	filledHeartColor := color.RGBA{255, 0, 0, 255}
	for i := 0; i < filledHearts; i++ {
		x := i * heartWidth
		y := 0
		drawHeart(img, x, y, filledHeartColor)
	}

	// Draw empty hearts
	emptyHeartColor := color.RGBA{0, 0, 0, 255}
	for i := filledHearts; i < totalHearts; i++ {
		x := i * heartWidth
		y := 0
		drawHeart(img, x, y, emptyHeartColor)
	}

	// Create the output file
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the image as PNG and write it to the file
	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}

func drawHeart(img *image.RGBA, x, y int, c color.RGBA) {
	// Draw a heart shape at the specified position
	// You can implement your own heart drawing algorithm here

	// Example: Draw a rectangle
	rect := image.Rect(x, y, x+20, y+20)
	draw.Draw(img, rect, &image.Uniform{c}, image.Point{}, draw.Src)
}
