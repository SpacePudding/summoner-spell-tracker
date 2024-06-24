package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	BLUESIDE = 100
	REDSIDE  = 200
)

type OngoingMatch struct {
	Participants []Participants `json:"participants"`
}

type Participants struct {
	PuuId      string `json:"puuid"`
	TeamId     int    `json:"teamId"`
	Spell1Id   int    `json:"spell1Id"`
	Spell2Id   int    `json:"spell2Id"`
	ChampionId int    `json:"championId"`
}

type EnemyData struct {
	spell1Id   int
	spell2Id   int
	championId int
}

func obtainEnemyTeamAssets() ([]EnemyData, error) {
	ongoingMatch, err := obtainOngoingMatch()
	if err != nil {
		return []EnemyData{}, err
	}
	return extractEnemyData(ongoingMatch.Participants)
}

func obtainOngoingMatch() (OngoingMatch, error) {
	encryptedPUUID := os.Getenv("PUUID")
	urlEndpoint := fmt.Sprintf("https://euw1.api.riotgames.com/lol/spectator/v5/active-games/by-summoner/%s", encryptedPUUID)

	headers := map[string]string{
		"X-Riot-Token": os.Getenv("API_KEY"),
	}

	body, err := FetchJSON(urlEndpoint, headers)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return OngoingMatch{}, errors.New("game not found")
		}
		return OngoingMatch{}, err
	}

	var ongoingMatch OngoingMatch
	err = json.Unmarshal(body, &ongoingMatch)
	if err != nil {
		return OngoingMatch{}, err
	}

	return ongoingMatch, nil
}

func extractEnemyData(participants []Participants) ([]EnemyData, error) {
	enemyTeamId, err := determineEnemyTeamId(participants)
	if err != nil {
		return []EnemyData{}, err
	}

	enemyDataArr := []EnemyData{}

	for _, participant := range participants {
		if participant.TeamId == enemyTeamId {
			enemyData := EnemyData{
				spell1Id:   participant.Spell1Id,
				spell2Id:   participant.Spell2Id,
				championId: participant.ChampionId,
			}
			enemyDataArr = append(enemyDataArr, enemyData)
		}
	}
	return enemyDataArr, nil
}

func determineEnemyTeamId(participants []Participants) (int, error) {
	for _, particiant := range participants {
		if particiant.PuuId == os.Getenv("PUUID") {

			allyTeamId := particiant.TeamId

			switch allyTeamId {
			case BLUESIDE:
				return REDSIDE, nil
			case REDSIDE:
				return BLUESIDE, nil
			default:
				return 0, fmt.Errorf("unknown teamId %d", allyTeamId)
			}
		}
	}
	return 0, errors.New("your PUUID couldn't be found in the active game")
}
