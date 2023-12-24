package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

const (
	GridSize   = 160 // Size of each grid cell
	GridCountX = 4   // Number of grids horizontally
	GridCountY = 4   // Number of grids vertically
)

type Game struct {
	Player *Player
}

func NewGame() *Game {
	return &Game{
		Player: NewPlayer(),
	}
}

func (g *Game) Update() error {
	g.Player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawGrid(screen)
	g.Player.Draw(screen)

	// Display the player's grid position
	gridX, gridY := g.Player.GridPosition()
	positionText := fmt.Sprintf("Grid Position: %d, %d", gridX, gridY)
	ebitenutil.DebugPrint(screen, positionText)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return GridSize * 4, GridSize * 4 // Fixed size for a 4x4 grid
}

func drawGrid(screen *ebiten.Image) {
	lineThickness := float32(1)  // Adjust the line thickness as needed
	gridLineColor := color.White // Define the grid line color

	for x := 0; x <= 4; x++ {
		x1 := float32(x * GridSize)
		// Vertical lines
		vector.StrokeLine(screen, x1, 0, x1, float32(GridSize*4), lineThickness, gridLineColor, false)
	}

	for y := 0; y <= 4; y++ {
		y1 := float32(y * GridSize)
		// Horizontal lines
		vector.StrokeLine(screen, 0, y1, float32(GridSize*4), y1, lineThickness, gridLineColor, false)
	}
}
