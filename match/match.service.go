package match

import (
	"motome.com.au/fuzball-services/team"
)

var matchs []Match

func GetAll () []*Match {
	matchsDao, err := daoGetAll()
	matchs := make([]*Match, 0)

	for _, matchDto := range matchsDao {
		var match *Match
		match, err = mapMatch(matchDto)
		matchs = append(matchs, match)
	}

	if (err == nil) {
		return matchs
	}

	println(err.Error())
	return nil
}

func GetById (id string) (*Match, error) {
	matchDto, err := daoGetById(id)

	var match *Match

	if (err == nil) {
		match, err = mapMatch(matchDto)
	}

	if (err == nil) {
		return match, nil
	}

	return nil, err
}

func Create(match Match) *Match {
	matchDto := MatchDto{
		ID: match.ID,
		TeamAId: match.TeamA.ID,
		TeamAScore: match.TeamAScore,
		TeamBId: match.TeamB.ID,
		TeamBScore: match.TeamBScore,
	}

	updatedMatchDto, err := daoCreate(matchDto)
	var updatedMatch *Match

	if (err == nil) {
		updatedMatch, err = mapMatch(updatedMatchDto)
	}

	if (err != nil) {
		println("Delete erred: ", err.Error())
	}

	return updatedMatch
}

func DeleteById(id string)  {
	err := daoDelete(id)

	if err != nil {
		println("Delete erred: ", err.Error())
	}
}

func mapMatch(dto *MatchDto) (*Match, error) {
	teamA, err := team.GetById(dto.TeamAId)
	teamB, err := team.GetById(dto.TeamBId)

	match := Match{
		ID: dto.ID,
		TeamA: teamA,
		TeamAScore: dto.TeamAScore,
		TeamB: teamB,
		TeamBScore: dto.TeamBScore,
	}

	return &match, err
}