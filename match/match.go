package match

import "motome.com.au/fuzball-services/team"

type Match struct {
	ID        string   `json:"id,omitempty"`
	TeamA *team.Team `json:"teamA,omitempty"`
	TeamAScore int `json:"teamAScore,omitempty"`
	TeamB *team.Team `json:"teamB,omitempty"`
	TeamBScore int `json:"teamBScore,omitempty"`
}