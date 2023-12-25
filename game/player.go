package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

const (
	playerSize  = 20 // Size of the player
	PlayerSpeed = 4.0
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
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.X -= PlayerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.X += PlayerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Y -= PlayerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Y += PlayerSpeed
	}

	// Adjust the boundary check considering the zoom factor
	maxPosX := float64(GridSize*GridCountX) - float64(playerSize)
	maxPosY := float64(GridSize*GridCountY) - float64(playerSize)

	p.X = clamp(p.X, 0, maxPosX)
	p.Y = clamp(p.Y, 0, maxPosY)
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
