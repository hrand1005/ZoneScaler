package worker

import (
	"encoding/json"
	"net/http"

	"github.com/hrand1005/ZoneScaler/common"
	"github.com/phuslu/log"
)

// PlayerDataHandler is a stand-in HTTP handler for a game client update.
func PlayerDataHandler(w http.ResponseWriter, r *http.Request) {
	var playerData common.PlayerData
	err := json.NewDecoder(r.Body).Decode(&playerData)
	if err != nil {
		http.Error(w, "Failed to parse player data", http.StatusBadRequest)
		return
	}
	log.Info().Msgf("Received player data: %+v", playerData)

	w.WriteHeader(http.StatusOK)
	log.Info().Msg("Player data received successfully")
}
