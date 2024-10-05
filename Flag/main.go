package main

import (
	"image/color"
	"math"

	"github.com/fogleman/gg"
)

func main() {
	// Define the width and height of the flag
	const width = 900
	const height = 600

	// Create a new context with the specified width and height
	dc := gg.NewContext(width, height)

	// Define the colors
	orange := color.RGBA{255, 153, 51, 255} // Top stripe - Orange
	white := color.RGBA{255, 255, 255, 255} // Middle stripe - White
	green := color.RGBA{19, 136, 8, 255}    // Bottom stripe - Green
	blue := color.RGBA{0, 0, 128, 255}      // Ashoka Chakra - Blue

	// Draw the orange stripe
	dc.SetColor(orange)
	dc.DrawRectangle(0, 0, width, height/3)
	dc.Fill()

	// Draw the white stripe
	dc.SetColor(white)
	dc.DrawRectangle(0, height/3, width, height/3)
	dc.Fill()

	// Draw the green stripe
	dc.SetColor(green)
	dc.DrawRectangle(0, 2*(height/3), width, height/3)
	dc.Fill()

	// Draw the Ashoka Chakra (24-spoke wheel)
	dc.SetColor(blue)
	dc.DrawCircle(width/2, height/2, height/6)
	dc.SetLineWidth(5)
	dc.Stroke()

	// Draw the 24 spokes of the Ashoka Chakra
	for i := 0; i < 24; i++ {
		angle := math.Pi * 2 * float64(i) / 24 // Convert i to radians
		x := width/2 + (height/6-5)*math.Cos(angle)
		y := height/2 + (height/6-5)*math.Sin(angle)
		dc.DrawLine(width/2, height/2, x, y)
		dc.Stroke()
	}

	// Save the image to a file
	dc.SavePNG("indian_flag.png")
}
