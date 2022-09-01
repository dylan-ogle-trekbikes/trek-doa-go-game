package main

type GameMode int

const (
	Start GameMode = iota
	Active
	Paused
	GameOver
)

func (m GameMode) isPaused() bool {
	return m == Paused
}

func (m GameMode) isActive() bool {
	return m == Active
}

func (m GameMode) isGameOver() bool {
	return m == GameOver
}

func (m GameMode) isStart() bool {
	return m == Start
}
