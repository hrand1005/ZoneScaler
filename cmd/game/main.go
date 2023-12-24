package main

import (
	"github.com/phuslu/log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hrand1005/ZoneScaler/game"
)

func main() {

	// Ebiten game loop runs in the main thread
	ebitenGame := game.NewGame()
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(ebitenGame); err != nil {
		log.Fatal().Err(err).Msg("Ebiten RunGame failed")
	}
}
