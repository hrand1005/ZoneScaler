package game

import (
	"github.com/hajimehoshi/ebiten/v2"
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
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawPlayerGridPosition(screen, g.Player)
	drawGameWorld(screen, g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return GridSize * GridCountX * 2, GridSize * GridCountY * 2 // Larger window size
}
