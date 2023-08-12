package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	buttonImage *ebiten.Image
	buttonRect  image.Rectangle
)

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if x >= buttonRect.Min.X && x <= buttonRect.Max.X &&
			y >= buttonRect.Min.Y && y <= buttonRect.Max.Y {
			fmt.Println("Hello")
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(buttonRect.Min.X), float64(buttonRect.Min.Y))
	screen.DrawImage(buttonImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return buttonRect.Dx(), buttonRect.Dy()
	// return 200, 200
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Button Example")

	img, _, err := ebitenutil.NewImageFromFile("button.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	buttonImage = img
	buttonRect = img.Bounds()

	game := &Game{}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
