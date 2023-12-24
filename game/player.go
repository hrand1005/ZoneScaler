package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

const (
	playerSize = 20 // Size of the player
)

type Player struct {
	X, Y float64 // Pixel position
}

func NewPlayer() *Player {
	return &Player{
		X: GridSize / 2, // Center the player in the first cell
		Y: GridSize / 2,
	}
}

func (p *Player) Update() {
	// Handle player movement
	speed := 2.0 // Speed of the player
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.X -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.X += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Y += speed
	}

	// Keep the player within the window bounds
	if p.X < 0 {
		p.X = 0
	}
	if p.Y < 0 {
		p.Y = 0
	}

	// Adjust the boundary check to account for the player's size
	maxPosX := float64(3*GridSize + GridSize - playerSize)
	maxPosY := float64(3*GridSize + GridSize - playerSize)

	if p.X > maxPosX {
		p.X = maxPosX
	}
	if p.Y > maxPosY {
		p.Y = maxPosY
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Draw the player
	ebitenutil.DrawRect(screen, p.X, p.Y, playerSize, playerSize, color.RGBA{0xff, 0x00, 0x00, 0xff})
}

func (p *Player) GridPosition() (int, int) {
	// Calculate the grid position
	gridX := int(p.X / GridSize)
	gridY := int(p.Y / GridSize)
	return gridX, gridY
}
