package common

type PlayerData struct {
	PlayerID   int    `json:"player_id"`
	PlayerName string `json:"player_name"`
	Score      int    `json:"score"`
}
