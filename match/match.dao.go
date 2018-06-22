package match

import (
	"motome.com.au/fuzball-services/db"
	"database/sql"
)

func daoGetAll() ([]*Match, error) {
	var dbConnection = db.Get()
	var matchs []*Match

	rows, err := dbConnection.Query("SELECT * FROM match")
	if err == nil {
		defer rows.Close()
		matchs, err = daoMap(rows)
	}

	return matchs, err
}

func daoGetById(id string) (*Match, error) {
	var dbConnection = db.Get()
	var matchs []*Match
	var match *Match

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

func daoCreate(match Match) (*Match, error) {
	var dbConnection = db.Get();
	var matchs []*Match
	var updatedMatch *Match

	rows, err := dbConnection.Query(
		"insert into match(firstname, lastname) values($1, $2) returning *",
		match.Firstname, match.Lastname)

	if err == nil {
		rows.NextResultSet()
		matchs, err = daoMap(rows)
		updatedMatch = matchs[0]
		rows.Close()
	}

	return updatedMatch, err
}

func daoMap(rows *sql.Rows) ([]*Match, error) {
	var err error
	matchs := make([]*Match, 0)

	for rows.Next() {
		match := new(Match)

		err = rows.Scan(&match.ID, &match.Firstname, &match.Lastname)
		if err != nil {
			println(err.Error())
		}

		matchs = append(matchs, match)
	}

	return matchs, err
}