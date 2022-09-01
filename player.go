package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	bikeWidth  = 80
	bikeHeight = 60
)

var (
	l1_x int
	l1_y int
	r1_x int
	r1_y int

	l2_x int
	l2_y int
	r2_x int
	r2_y int
)

//Player functions

func (g *Game) isJumpKeyJustPressed() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return true
	}
	return false
}

func (g *Game) handleMovement() {
	if g.isJumpKeyJustPressed() && !isJumping {
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

func (g *Game) anyHits() bool {
	for i := 0; i < g.logXs.Length(); i++ {
		if g.hit(i) {
			return true
		}
	}
	return false
}

func (g *Game) hit(logIndex int) bool {

	x0 := DefaultBikeX
	x1 := DefaultBikeX + bikeWidth

	y0 := g.yPos + bikeHeight
	y1 := g.yPos

	l1_x = x0
	l1_y = y0

	r1_x = x1
	r1_y = y1

	//get log box
	_, h := logImage.Size()
	log_x0 := g.logXs.ValueAt(logIndex)
	log_x1 := log_x0 + logWidth

	log_y0 := Baseline - float64(h)
	log_y1 := Baseline

	l2_x = int(log_x0)
	l2_y = int(log_y0)

	r2_x = int(log_x1)
	r2_y = int(log_y1)

	//determine if bike and log are overlapping
	// If one rectangle is on left side of other
	if l1_x > r2_x || l2_x > r1_x {
		return false

	}

	// If one rectangle is above other
	if r1_y > l2_y || r2_y > l1_y {
		return false
	}
	return true
}

func (g *Game) anyScore() bool {
	for i := 0; i < g.logXs.Length(); i++ {
		if g.score(i) {
			return true
		}
	}
	return false
}

func (g Game) score(logIndex int) bool {
	x0 := DefaultBikeX
	x1 := DefaultBikeX + bikeWidth

	y0 := g.yPos + bikeHeight
	y1 := g.yPos

	l1_x = x0
	l1_y = y0

	r1_x = x1
	r1_y = y1

	//get log box
	_, h := logImage.Size()
	log_x0 := g.logXs.ValueAt(logIndex)
	log_x1 := log_x0 + logWidth

	log_y0 := Baseline - float64(h)
	log_y1 := Baseline

	l2_x = int(log_x0)
	l2_y = int(log_y0)

	r2_x = int(log_x1)
	r2_y = int(log_y1)

	if (l1_x < r2_x && l2_x < r1_x) && (r1_y > l2_y || r2_y > l1_y) {
		return true
	}
	return false
}
