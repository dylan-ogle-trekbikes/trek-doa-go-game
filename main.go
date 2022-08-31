package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"trekdoagame/utils"
)

type Game struct {
	//player position
	yPos int
	vy   float64

	//active logs position
	logXs utils.Stack

	isGameOver bool
}

const (
	ScreenWidth  = 800
	ScreenHeight = 420
	Baseline     = ScreenHeight / 1.75
	Gravity      = 19
	Acceleration = 2
)

var (
	bikeImage    *ebiten.Image
	bikeVy       float64
	bikeBaseline int

	logImage *ebiten.Image
	logWidth = 45.0
	logSpeed = 6.0

	isJumping = false
)

// Custom game functions

func (g *Game) init() {
	bikeImg, _, err := ebitenutil.NewImageFromFile(("./images/bike_image.png"))
	if err != nil {
		log.Fatal(err)
	}
	logImg, _, err := ebitenutil.NewImageFromFile(("./images/square_log.png"))
	if err != nil {
		log.Fatal(err)
	}
	bikeImage = ebiten.NewImageFromImage(bikeImg)
	logImage = ebiten.NewImageFromImage(logImg)

	_, h := bikeImage.Size()
	calculatedBaseline := Baseline - h
	g.yPos = calculatedBaseline
	bikeBaseline = calculatedBaseline

}

func (g *Game) drawBike(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	//move to left side of screen
	options.GeoM.Translate(25, float64(g.yPos))

	screen.DrawImage(bikeImage, options)
}

func (g *Game) drawLog(screen *ebiten.Image, pos float64) {
	options := &ebiten.DrawImageOptions{}
	_, h := logImage.Size()
	options.GeoM.Translate(pos, Baseline-float64(h))
	screen.DrawImage(logImage, options)
}

func (g *Game) drawLogs(screen *ebiten.Image) {
	logsToDraw := g.logXs.Length()
	for i := 0; i < logsToDraw; i++ {
		logPos := g.logXs.ValueAt(i)
		g.drawLog(screen, logPos)
	}
}

func (g *Game) createLog() {
	g.logXs.Push(ScreenWidth - logWidth)
}

func (g *Game) removeLog() {
	g.logXs.Pop()
}

func (g *Game) isLogOffScreen(logPos float64) bool {
	if logPos+logWidth < 0 {
		return true
	}
	return false
}

func (g *Game) moveLogs() {
	if !g.logXs.IsEmpty() {
		for i := 0; i < g.logXs.Length(); i++ {
			g.moveLog(i)
		}
	}
}

func (g *Game) moveLog(logIndex int) {
	g.logXs[logIndex] = g.logXs[logIndex] - logSpeed
	if g.isLogOffScreen(g.logXs[logIndex]) {
		g.removeLog()
	}
}

func (g *Game) isKeyJustPressed() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return true
	}
	return false
}

func (g *Game) handleMovement() {
	if g.isKeyJustPressed() && !isJumping {
		g.vy = -Gravity
		g.yPos += 5
		isJumping = true
	}

	g.yPos += int(g.vy)

	// Gravity
	g.vy += float64(Acceleration * .75)

	//upper limit on gravity
	if g.vy > Gravity {
		g.vy = Gravity
	}

	//stop moving position if bike hits 0
	if g.yPos <= 0 {
		g.yPos = 0
	}

	//stop velocity if hits baseline
	if g.yPos >= bikeBaseline {
		g.yPos = bikeBaseline
		g.vy = 0
		isJumping = false
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {

	g.moveLogs()

	g.handleMovement()

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	g.drawBike(screen)
	g.drawLogs(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	g := &Game{}
	g.init()
	g.createLog()
	return g
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
