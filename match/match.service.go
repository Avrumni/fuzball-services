package match

import (
	"errors"
)

var matchs []Match

//func Query(query string, args ...interface{}) []*db.Entity {
//	var dbConnection = db.Get()
//
//	rows, err := dbConnection.Query(query, args)
//	if err != nil {
//		println(err.Error())
//	}
//
//	defer rows.Close()
//
//	entities := make([]*db.Entity, 0)
//
//
//	columns, _ := rows.ColumnTypes()
//
//	for rows.Next() {
//		entity := new(db.Entity)
//
//		for _, column := range columns {
//			entity.
//		}
//
//		err := rows.Scan(&entity.ID, &entity.Firstname, &entity.Lastname)
//		if err != nil {
//			println(err.Error())
//		}
//		entities = append(entities, entity)
//	}
//
//	return entities
//}

func GetAll () []*Match {
	matchs, err := daoGetAll()

	if (err == nil) {
		return matchs
	}

	println(err.Error())
	return nil
}

func GetById (id string) (*Match, error) {
	match, err := daoGetById(id)

	if (err == nil) {
		return match, nil
	}

	return nil, errors.New("Not Found")
}

func Create(match Match) *Match {
	updatedMatch, err := daoCreate(match)

	if (err != nil) {
		println("Delete erred: ", err.Error())
		return nil
	}

	return updatedMatch
}

func DeleteById(id string)  {
	err := daoDelete(id)

	if err != nil {
		println("Delete erred: ", err.Error())
	}
}
