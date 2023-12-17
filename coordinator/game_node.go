package main

import (
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
