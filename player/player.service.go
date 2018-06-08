package player

import (
	"errors"
	"motome.com.au/fuzball-services/db"
)

var players []Player

func GetAll () []Player {
	return players;
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