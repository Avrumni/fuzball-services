package match

import (
	"motome.com.au/fuzball-services/db"
	"database/sql"
)

func daoGetAll() ([]*MatchDto, error) {
	var dbConnection = db.Get()
	var matchs []*MatchDto

	rows, err := dbConnection.Query("SELECT * FROM match")
	if err == nil {
		defer rows.Close()
		matchs, err = daoMap(rows)
	}

	return matchs, err
}

func daoGetById(id string) (*MatchDto, error) {
	var dbConnection = db.Get()
	var matchs []*MatchDto
	var match *MatchDto

	rows, err := dbConnection.Query("SELECT * FROM match where id = $1", id)

	if (err == nil) {
		matchs, err = daoMap(rows)
		match = matchs[0]
		rows.Close()
	}

	return match, err
}

func daoDelete(id string) error {
	var dbConnection = db.Get()

	rows, err := dbConnection.Query("delete from match where id=$1", id)

	rows.Close()

	return err
}

func daoCreate(match MatchDto) (*MatchDto, error) {
	var dbConnection = db.Get();
	var matchs []*MatchDto
	var updatedMatch *MatchDto

	rows, err := dbConnection.Query(
		"insert into match(team_a_id, team_a_score, team_b_id, team_b_score) values($1, $2, $3, $4) returning *",
		match.TeamAId, match.TeamAScore, match.TeamBId, match.TeamBScore)

	if err == nil {
		rows.NextResultSet()
		matchs, err = daoMap(rows)
		updatedMatch = matchs[0]
		rows.Close()
	}

	return updatedMatch, err
}

func daoMap(rows *sql.Rows) ([]*MatchDto, error) {
	var err error
	matchs := make([]*MatchDto, 0)

	for rows.Next() {
		match := new(MatchDto)

		err = rows.Scan(&match.ID, &match.TeamAId, &match.TeamAScore, &match.TeamBId, &match.TeamBScore)
		if err != nil {
			println(err.Error())
		}

		matchs = append(matchs, match)
	}

	return matchs, err
}