package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	//player position
	yPos int
	vy   int

	//active logs position
	logXs []int

	isGameOver bool
}

const (
	ScreenWidth  = 800
	ScreenHeight = 420
	Baseline     = ScreenHeight / 1.75
)

var (
	bikeImage *ebiten.Image
	logImage  *ebiten.Image
)

// Custom game functions

func (g *Game) init() {
	bikeImg, _, err := ebitenutil.NewImageFromFile(("./images/bike_image.png"))
	if err != nil {
		log.Fatal(err)
	}
	logImg, _, err := ebitenutil.NewImageFromFile(("./images/log.png"))
	if err != nil {
		log.Fatal(err)
	}
	bikeImage = ebiten.NewImageFromImage(bikeImg)
	logImage = ebiten.NewImageFromImage(logImg)

}

func (g *Game) drawBike(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	_, h := bikeImage.Size()
	//move to left side of screen
	options.GeoM.Translate(0, Baseline-float64(h))
	screen.DrawImage(bikeImage, options)
}

func (g *Game) drawLog(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	w, h := logImage.Size()
	options.GeoM.Translate(ScreenWidth-float64(w), Baseline-float64(h))
	screen.DrawImage(logImage, options)
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	g.drawBike(screen)
	g.drawLog(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func main() {

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
