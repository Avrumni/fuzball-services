package player

import "errors"

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
	player.ID = "UUID"
	players = append(players, player)
	return player
}

func DeleteById(id string)  {
	for index, item := range players {
		if item.ID == id {
			players = append(players[:index], players[index+1:]...)
			break
		}
	}
}