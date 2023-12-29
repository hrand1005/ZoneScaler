package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

func drawPlayerGridPosition(screen *ebiten.Image, player *Player) {
	gridX, gridY := player.GridPosition()
	sampleText := fmt.Sprintf("Grid Position: %d, %d", gridX, gridY)
	text.Draw(screen, sampleText, mplusBigFont, 15, 150, color.White)
}

func drawGameWorld(screen *ebiten.Image, gm *Game) {
	op := &ebiten.DrawImageOptions{}
	applyCameraTransformations(op, gm.Camera, gm.Player)

	offscreen := ebiten.NewImage(GridSize*GridCountX, GridSize*GridCountY)
	drawGrid(offscreen)
	gm.Player.Draw(offscreen)

	screen.DrawImage(offscreen, op)
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

func applyCameraTransformations(op *ebiten.DrawImageOptions, camera *Camera, player *Player) {
	Width, Height := ebiten.WindowSize()
	cameraOffsetX := camera.X - float64(Width)/(2*ZoomFactor)
	cameraOffsetY := camera.Y - float64(Height)/(2*ZoomFactor)

	op.GeoM.Translate(-cameraOffsetX, -cameraOffsetY)
	op.GeoM.Scale(ZoomFactor, ZoomFactor)
}
