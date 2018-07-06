package team

import (
	"motome.com.au/fuzball-services/player"
)

var teams []Team

func GetAll () []*Team {
	teamsDao, err := daoGetAll()
	teams := make([]*Team, 0)

	for _, teamDto := range teamsDao {
		var team *Team
		team, err = mapTeam(teamDto)
		teams = append(teams, team)
	}

	if (err == nil) {
		return teams
	}

	println(err.Error())
	return nil
}

func GetById (id string) (*Team, error) {
	teamDto, err := daoGetById(id)

	var team *Team

	if (err == nil) {
		team, err = mapTeam(teamDto)
	}

	if (err == nil) {
		return team, nil
	}

	return nil, err
}

func Create(team Team) *Team {
	teamDto := TeamDto{
		ID: team.ID,
		Player1Id: team.Player1.ID,
		Player2Id: team.Player2.ID,
	}

	updatedTeamDto, err := daoCreate(teamDto)
	var updatedTeam *Team

	if (err == nil) {
		updatedTeam, err = mapTeam(updatedTeamDto)
	}

	if (err != nil) {
		println("Delete erred: ", err.Error())
	}

	return updatedTeam
}

func DeleteById(id string)  {
	err := daoDelete(id)

	if err != nil {
		println("Delete erred: ", err.Error())
	}
}

func mapTeam(dto *TeamDto) (*Team, error) {
	player1, err := player.GetById(dto.Player1Id)
	player2, err := player.GetById(dto.Player2Id)

	team := Team{
		ID: dto.ID,
		Player1: player1,
		Player2: player2,
	}

	return &team, err
}