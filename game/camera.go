package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

const CameraFollowThreshold = 100.0 // Distance from center before the camera starts following

type Camera struct {
	X, Y float64 // Camera position
}

func NewCamera() *Camera {
	return &Camera{}
}

func (c *Camera) Update(playerX, playerY float64) {
	screenWidth, screenHeight := ebiten.WindowSize()

	// Calculate the centered camera position
	centeredCameraX := playerX - float64(screenWidth)/2
	centeredCameraY := playerY - float64(screenHeight)/2

	// Determine how far away the player is from the center of the camera's current position
	deltaX := centeredCameraX - c.X
	deltaY := centeredCameraY - c.Y

	// If the player is further than the CameraFollowThreshold from the camera's current position, adjust the camera
	if math.Abs(deltaX) > CameraFollowThreshold {
		if deltaX > 0 {
			c.X += deltaX - CameraFollowThreshold
		} else {
			c.X += deltaX + CameraFollowThreshold
		}
	}
	if math.Abs(deltaY) > CameraFollowThreshold {
		if deltaY > 0 {
			c.Y += deltaY - CameraFollowThreshold
		} else {
			c.Y += deltaY + CameraFollowThreshold
		}
	}

	// Clamp the camera position to the game world boundaries
	maxCameraX := float64(GridSize*GridCountX) - float64(screenWidth)
	maxCameraY := float64(GridSize*GridCountY) - float64(screenHeight)

	c.X = clamp(c.X, 0, maxCameraX)
	c.Y = clamp(c.Y, 0, maxCameraY)
}

// clamp helper function
func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
