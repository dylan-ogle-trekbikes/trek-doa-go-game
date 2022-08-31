package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

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
	ScreenWidth   = 800
	ScreenHeight  = 420
	Baseline      = ScreenHeight / 1.75
	Gravity       = 19
	Acceleration  = 2
	LogSpawnDelay = 750
)

var (
	bikeImage    *ebiten.Image
	bikeVy       float64
	bikeBaseline int

	logImage *ebiten.Image
	logWidth = 45.0
	logSpeed = 6.0

	isJumping = false

	lastSpawnedLog int64
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

//Log functions

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

func (g *Game) spawnLogs() {
	//wait at least 1-3 seconds before spawning a log
	if (time.Now().UnixMilli() - lastSpawnedLog) < LogSpawnDelay {
		return
	}
	//flip a coin if a log should spawn
	rand.Seed(time.Now().UnixNano())

	// flip the coin
	shouldSpawn := rand.Intn(100)
	println(time.Now().UnixMilli() - lastSpawnedLog)
	if shouldSpawn > 90 {
		g.createLog()
	}

}

func (g *Game) createLog() {
	g.logXs.Push(ScreenWidth)
	lastSpawnedLog = time.Now().UnixMilli()
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
		// g.removeLog()
	}
}

//Player functions

func (g *Game) isKeyJustPressed() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return true
	}
	return false
}

func (g *Game) handleMovement() {
	if g.isKeyJustPressed() && !isJumping {
		g.vy = -Gravity
		g.yPos++
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

	g.spawnLogs()
	g.moveLogs()

	g.handleMovement()

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
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
	return g
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
