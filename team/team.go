package team

import "motome.com.au/fuzball-services/player"

type Team struct {
	ID      string        `json:"id,omitempty"`
	Player1 *player.Player `json:"player1,omitempty"`
	Player2 *player.Player `json:"player2,omitempty"`
}
