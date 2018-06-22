package team

import (
	"motome.com.au/fuzball-services/db"
	"database/sql"
)

func daoGetAll() ([]*Team, error) {
	var dbConnection = db.Get()
	var teams []*Team

	rows, err := dbConnection.Query("SELECT * FROM team")
	if err == nil {
		defer rows.Close()
		teams, err = daoMap(rows)
	}

	return teams, err
}

func daoGetById(id string) (*Team, error) {
	var dbConnection = db.Get()
	var teams []*Team
	var team *Team

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

func daoCreate(team Team) (*Team, error) {
	var dbConnection = db.Get();
	var teams []*Team
	var updatedTeam *Team

	rows, err := dbConnection.Query(
		"insert into team(firstname, lastname) values($1, $2) returning *",
		team.Firstname, team.Lastname)

	if err == nil {
		rows.NextResultSet()
		teams, err = daoMap(rows)
		updatedTeam = teams[0]
		rows.Close()
	}

	return updatedTeam, err
}

func daoMap(rows *sql.Rows) ([]*Team, error) {
	var err error
	teams := make([]*Team, 0)

	for rows.Next() {
		team := new(Team)

		err = rows.Scan(&team.ID, &team.Firstname, &team.Lastname)
		if err != nil {
			println(err.Error())
		}

		teams = append(teams, team)
	}

	return teams, err
}