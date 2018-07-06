package team

import (
	"motome.com.au/fuzball-services/db"
	"database/sql"
)

func daoGetAll() ([]*TeamDto, error) {
	var dbConnection = db.Get()
	var teams []*TeamDto

	rows, err := dbConnection.Query("SELECT * FROM team")
	if err == nil {
		defer rows.Close()
		teams, err = daoMap(rows)
	}

	return teams, err
}

func daoGetById(id string) (*TeamDto, error) {
	var dbConnection = db.Get()
	var teams []*TeamDto
	var team *TeamDto

	rows, err := dbConnection.Query("SELECT * FROM team where id = $1", id)

	if (err == nil) {
		teams, err = daoMap(rows)
		team = teams[0]
		rows.Close()
	}

	return team, err
}

func daoDelete(id string) error {
	var dbConnection = db.Get()

	rows, err := dbConnection.Query("delete from team where id=$1", id)

	rows.Close()

	return err
}

func daoCreate(team TeamDto) (*TeamDto, error) {
	var dbConnection = db.Get();
	var teams []*TeamDto
	var updatedTeam *TeamDto

	rows, err := dbConnection.Query(
		"insert into team(player_1_id, player_2_id) values($1, $2) returning *",
		team.Player1Id, team.Player2Id)

	if err == nil {
		rows.NextResultSet()
		teams, err = daoMap(rows)
		updatedTeam = teams[0]
		rows.Close()
	}

	return updatedTeam, err
}

func daoMap(rows *sql.Rows) ([]*TeamDto, error) {
	var err error
	teamsDto := make([]*TeamDto, 0)

	for rows.Next() {
		teamDto := new(TeamDto)

		err = rows.Scan(&teamDto.ID, &teamDto.Player1Id, &teamDto.Player2Id)
		if err != nil {
			println(err.Error())
		}

		teamsDto = append(teamsDto, teamDto)
	}

	return teamsDto, err
}