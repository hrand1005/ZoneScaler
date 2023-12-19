package common

import (
	"fmt"
	"time"
)

// GameNode represents a single game server node
type GameNode struct {
	ID            string    `json:"id"`
	Address       string    `json:"address"`
	Load          int       `json:"load"`
	Regions       []string  `json:"regions"`
	IsActive      bool      `json:"isActive"`
	LastHeartbeat time.Time `json:"-"`
}

func (g *GameNode) String() string {
	return fmt.Sprintf("ID: %s\nAddress: %s\nLoad: %d\nRegions: %v\nIsActive: %t\nLastHeartbeat: %s",
		g.ID, g.Address, g.Load, g.Regions, g.IsActive, g.LastHeartbeat.Format(time.RFC3339))
}

func (g *GameNode) Copy() *GameNode {
	return &GameNode{
		ID:            g.ID,
		Address:       g.Address,
		Load:          g.Load,
		Regions:       append([]string{}, g.Regions...),
		IsActive:      g.IsActive,
		LastHeartbeat: g.LastHeartbeat,
	}
}
