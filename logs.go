package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

//Log functions

func (g *Game) drawLog(screen *ebiten.Image, xPos float64) {
	options := &ebiten.DrawImageOptions{}
	_, h := logImage.Size()
	yPos := Baseline - float64(h)
	options.GeoM.Translate(xPos, yPos)
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
