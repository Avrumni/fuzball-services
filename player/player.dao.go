package player

import (
	"motome.com.au/fuzball-services/db"
	"database/sql"
)

func daoGetAll() ([]*Player, error) {
	var dbConnection = db.Get()
	var players []*Player

	rows, err := dbConnection.Query("SELECT * FROM player")
	if err == nil {
		defer rows.Close()
		players, err = daoMap(rows)
	}

	return players, err
}

func daoGetById(id string) (*Player, error) {
	var dbConnection = db.Get()
	var players []*Player
	var player *Player

	rows, err := dbConnection.Query("SELECT * FROM player where id = $1", id)

	if (err == nil) {
		players, err = daoMap(rows)
		player = players[0]
		rows.Close()
	}

	return player, err
}

func daoDelete(id string) error {
	var dbConnection = db.Get()

	rows, err := dbConnection.Query("delete from player where id=$1", id)

	rows.Close()

	return err
}

func daoCreate(player Player) (*Player, error) {
	var dbConnection = db.Get();
	var players []*Player
	var updatedPlayer *Player

	rows, err := dbConnection.Query(
		"insert into player(firstname, lastname) values($1, $2) returning *",
		player.Firstname, player.Lastname)

	if err == nil {
		rows.NextResultSet()
		players, err = daoMap(rows)
		updatedPlayer = players[0]
		rows.Close()
	}

	return updatedPlayer, err
}

func daoMap(rows *sql.Rows) ([]*Player, error) {
	var err error
	players := make([]*Player, 0)

	for rows.Next() {
		player := new(Player)

		err = rows.Scan(&player.ID, &player.Firstname, &player.Lastname)
		if err != nil {
			println(err.Error())
		}

		players = append(players, player)
	}

	return players, err
}