package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/opentype"

	"trekdoagame/utils"
)

type Game struct {
	//player position
	yPos int
	//player y velocity
	vy float64
	//stack of logs
	logXs utils.Stack
	//current game state (start, paused, game over)
	CurrentMode GameMode
}

const (
	ScreenWidth   = 800
	ScreenHeight  = 420
	Baseline      = ScreenHeight / 1.75
	Gravity       = 15
	Acceleration  = 1
	LogSpawnDelay = 750
	DefaultBikeX  = 25
	startText     = "Press any button to start"
	endText       = "Game Over! Press 'r' to restart"
)

var (
	MainFont font.Face

	bikeImage    *ebiten.Image
	bikeVy       float64
	bikeBaseline int

	logImage       *ebiten.Image
	logWidth       = 45.0
	logHeight      = 45.0
	logSpeed       = 6.0
	lastSpawnedLog int64

	isJumping = false
	hasScored = false

	score     int
	highScore int
)

// Custom game functions

func (g *Game) initGame() {
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
	g.logXs.Pop()
}

func (g *Game) init() {

	//font setup
	tt, err := opentype.Parse(goitalic.TTF)
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	MainFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		log.Fatalf("NewFace: %v", err)
	}

	//game setup
	g.initGame()
}

func (g *Game) drawBike(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	//move to left side of screen
	options.GeoM.Translate(DefaultBikeX, float64(g.yPos))

	screen.DrawImage(bikeImage, options)
}

func (g *Game) drawStart(screen *ebiten.Image) {
	text.Draw(screen, startText, MainFont, 20, 80, color.White)
}

func (g *Game) drawScore(screen *ebiten.Image) {
	scoreText := fmt.Sprintf("Score: %d", score)
	highScoreText := fmt.Sprintf("High Score: %d", highScore)
	text.Draw(screen, scoreText, MainFont, ScreenWidth-250, 40, color.Black)
	text.Draw(screen, highScoreText, MainFont, ScreenWidth-250, 80, color.Black)
}

func (g *Game) drawGame(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	g.drawBike(screen)
	g.drawLogs(screen)
	g.drawScore(screen)
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	text.Draw(screen, endText, MainFont, 20, 80, color.White)
}

// Update proceeds the game state.
func (g *Game) Update() error {

	if len(inpututil.PressedKeys()) > 0 && g.CurrentMode.isStart() {
		g.CurrentMode = Active
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {

		if g.CurrentMode.isActive() {
			g.CurrentMode = Paused
		} else {
			g.CurrentMode = Active
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		println(g.CurrentMode)
		println(g.CurrentMode.isGameOver())
		if g.CurrentMode.isGameOver() {
			g.CurrentMode = Active
			g.logXs.Clear()
		}
	}

	if g.CurrentMode.isPaused() || g.CurrentMode.isGameOver() || g.CurrentMode.isStart() {
		return nil
	}

	if g.anyScore() {
		if !hasScored {
			score++
			hasScored = true
		}

	} else if hasScored {
		hasScored = false
	}

	if g.anyHits() {
		g.CurrentMode = GameOver
		if score > highScore {
			highScore = score
		}
		score = 0
	}

	g.spawnLogs()
	g.moveLogs()

	g.handleMovement()

	return nil
}

// Draw draws the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.CurrentMode.isStart() {
		g.drawStart(screen)
		return
	} else if g.CurrentMode.isActive() || g.CurrentMode.isPaused() {
		g.drawGame(screen)
	} else if g.CurrentMode.isGameOver() {
		g.drawGameOver(screen)
	}

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// Instantiate a new game
func NewGame() *Game {
	g := &Game{}
	g.init()
	g.CurrentMode = Start
	return g
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
