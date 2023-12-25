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
	GridCountX = 16  // Number of grids horizontally
	GridCountY = 16  // Number of grids vertically
	ZoomFactor = 8.0
)

type Game struct {
	Player *Player
	Camera *Camera
}

func NewGame() *Game {
	return &Game{
		Player: NewPlayer(),
		Camera: NewCamera(),
	}
}

func (g *Game) Update() error {
	g.Player.Update() // Update the camera position to center on the player
	g.Camera.Update(g.Player.X, g.Player.Y)

	// Debugging output
	fmt.Printf("Player Position: (%f, %f)\n", g.Player.X, g.Player.Y)
	fmt.Printf("Camera Position: (%f, %f)\n", g.Camera.X, g.Camera.Y)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Apply camera transformations
	op := &ebiten.DrawImageOptions{}

	// Calculate the offset for the camera to keep the player centered
	Width, Height := ebiten.WindowSize()
	cameraOffsetX := g.Camera.X - float64(Width)/(2*ZoomFactor)
	cameraOffsetY := g.Camera.Y - float64(Height)/(2*ZoomFactor)

	// Translate first to follow the player
	op.GeoM.Translate(-cameraOffsetX, -cameraOffsetY)
	// Then scale to apply zoom
	op.GeoM.Scale(ZoomFactor, ZoomFactor)

	// Create an offscreen image for drawing the game world
	offscreen := ebiten.NewImage(GridSize*GridCountX, GridSize*GridCountY)
	drawGrid(offscreen)
	g.Player.Draw(offscreen)

	// Draw the offscreen image to the screen with camera adjustments
	screen.DrawImage(offscreen, op)

	// Display player's grid position without camera offset
	GridX, GridY := g.Player.GridPosition()
	positionText := fmt.Sprintf("Grid Position: %d, %d", GridX, GridY)
	ebitenutil.DebugPrintAt(screen, positionText, 10, 10) // Fixed position on screen
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return GridSize * GridCountX * 2, GridSize * GridCountY * 2 // Larger window size
}

func drawGrid(screen *ebiten.Image) {
	lineThickness := float32(1)  // Adjust the line thickness as needed
	gridLineColor := color.White // Define the grid line color

	for x := 0; x <= GridCountY; x++ {
		x1 := float32(x * GridSize)
		// Vertical lines
		vector.StrokeLine(screen, x1, 0, x1, float32(GridSize*GridCountX), lineThickness, gridLineColor, false)
	}

	for y := 0; y <= GridCountX; y++ {
		y1 := float32(y * GridSize)
		// Horizontal lines
		vector.StrokeLine(screen, 0, y1, float32(GridSize*GridCountY), y1, lineThickness, gridLineColor, false)
	}
}
