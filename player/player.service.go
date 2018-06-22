package player

import (
	"database/sql"
)

var players []Player

func Make (row *sql.Rows) (*Player, error) {
	player := new(Player)

	return player, row.Scan(&player.ID, &player.Firstname, &player.Lastname)
}

func Query (query string, maker func(rows sql.Rows) (struct, error), args ...string) []struct {
	var dbConnection = db.Get();

	rows, err := dbConnection.Query(query, args)
	if err != nil {
		println(err.Error())
	}

	defer rows.Close()

	players := make([]*Player, 0)
	for rows.Next() {
		player := new(Player)
		err := rows.Scan(&player.ID, &player.Firstname, &player.Lastname)
		if err != nil {
			println(err.Error())
		}
		players = append(players, player)
	}
	if err = rows.Err(); err != nil {
		println(err.Error())
	}
	return players
}

func GetAll () []Player {
	return Query("SELECT * FROM player", Make)
}

func GetById (id string) (Player, error) {
	for _, item := range players {
		if item.ID == id {
			return item, nil
		}
	}

	return Player{}, errors.New("Not Found")
}

func Create(player Player) Player {
	var dbConnection = db.Get();

	rows, err := dbConnection.Query("insert into player(firstname, lastname) values($1, $2)", player.Firstname, player.Lastname)

	strRows, _ := rows.Columns()

	println(1, strRows)

	if err != nil {
		println("Create erred: ", err.Error())
	}

	rows.Close()
	return player
}

func DeleteById(id string)  {
	var dbConnection = db.Get();

	rows, err := dbConnection.Query("delete from player where id=$1", id)

	if err != nil {
		println("Delete erred: ", err.Error())
	}

	rows.Close()
}


func doMagic(rows sql.Rows) {
	columns, err := rows.Columns()

	if err != nil {
		println("Ahh", err.Error())
		return
	}

	for rows.Next() {
		var row = make(map[string]string, len(columns))
		for column := range columns {
			row[column] =
		}
	}
}