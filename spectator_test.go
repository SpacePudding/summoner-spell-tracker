package main

import (
	"os"
	"reflect"
	"testing"
)

func TestDetermineEnemyTeamId(t *testing.T) {
	os.Setenv("PUUID", "mocked-puuid")

	tests := []struct {
		name           string
		participants   []Participants
		expectedTeamId int
		expectedError  bool
	}{
		{
			name: "Enemy team is red side",
			participants: []Participants{
				{PuuId: "mocked-puuid", TeamId: 100, Spell1Id: 1, Spell2Id: 2, ChampionId: 101},
				{PuuId: "enemy-1", TeamId: 200, Spell1Id: 3, Spell2Id: 4, ChampionId: 102},
				{PuuId: "enemy-2", TeamId: 200, Spell1Id: 5, Spell2Id: 6, ChampionId: 103},
			},
			expectedTeamId: 200,
			expectedError:  false,
		},
		{
			name: "Enemy team is blue side",
			participants: []Participants{
				{PuuId: "mocked-puuid", TeamId: 200, Spell1Id: 1, Spell2Id: 2, ChampionId: 101},
				{PuuId: "enemy-1", TeamId: 100, Spell1Id: 3, Spell2Id: 4, ChampionId: 102},
				{PuuId: "enemy-2", TeamId: 100, Spell1Id: 5, Spell2Id: 6, ChampionId: 103},
			},
			expectedTeamId: 100,
			expectedError:  false,
		},
		{
			name: "PUUID not found",
			participants: []Participants{
				{PuuId: "some-other-puuid", TeamId: 200, Spell1Id: 1, Spell2Id: 2, ChampionId: 101},
				{PuuId: "enemy-1", TeamId: 100, Spell1Id: 3, Spell2Id: 4, ChampionId: 102},
				{PuuId: "enemy-2", TeamId: 100, Spell1Id: 5, Spell2Id: 6, ChampionId: 103},
			},
			expectedTeamId: 0,
			expectedError:  true,
		},
		{
			name: "Unknown teamId",
			participants: []Participants{
				{PuuId: "mocked-puuid", TeamId: 400, Spell1Id: 1, Spell2Id: 2, ChampionId: 101},
				{PuuId: "enemy-1", TeamId: 110, Spell1Id: 3, Spell2Id: 4, ChampionId: 102},
				{PuuId: "enemy-2", TeamId: 120, Spell1Id: 5, Spell2Id: 6, ChampionId: 103},
			},
			expectedTeamId: 0,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enemyTeamId, err := determineEnemyTeamId(tt.participants)
			if (err != nil) != tt.expectedError {
				t.Fatalf("expected error: %v, got: %v", tt.expectedError, err)
			}
			if enemyTeamId != tt.expectedTeamId {
				t.Fatalf("expected team id: %v, got: %v", tt.expectedTeamId, enemyTeamId)
			}
		})
	}
}

func TestExtractEnemyData(t *testing.T) {
	os.Setenv("PUUID", "mocked-puuid")

	tests := []struct {
		name              string
		participants      []Participants
		expectedEnemyData []EnemyData
	}{
		{
			name: "Extract EnemyTeam data",
			participants: []Participants{
				{PuuId: "mocked-puuid", TeamId: 100, Spell1Id: 1, Spell2Id: 2, ChampionId: 101},
				{PuuId: "ally-2", TeamId: 100, Spell1Id: 3, Spell2Id: 4, ChampionId: 102},
				{PuuId: "ally-3", TeamId: 100, Spell1Id: 5, Spell2Id: 6, ChampionId: 103},
				{PuuId: "ally-4", TeamId: 100, Spell1Id: 1, Spell2Id: 2, ChampionId: 104},
				{PuuId: "ally-5", TeamId: 100, Spell1Id: 3, Spell2Id: 4, ChampionId: 105},
				{PuuId: "enemy-1", TeamId: 200, Spell1Id: 5, Spell2Id: 6, ChampionId: 106},
				{PuuId: "enemy-2", TeamId: 200, Spell1Id: 1, Spell2Id: 2, ChampionId: 107},
				{PuuId: "enemy-3", TeamId: 200, Spell1Id: 3, Spell2Id: 4, ChampionId: 108},
				{PuuId: "enemy-4", TeamId: 200, Spell1Id: 5, Spell2Id: 6, ChampionId: 109},
				{PuuId: "enemy-5", TeamId: 200, Spell1Id: 1, Spell2Id: 2, ChampionId: 110},
			},
			expectedEnemyData: []EnemyData{
				{Spell1Id: 5, Spell2Id: 6, ChampionId: 106},
				{Spell1Id: 1, Spell2Id: 2, ChampionId: 107},
				{Spell1Id: 3, Spell2Id: 4, ChampionId: 108},
				{Spell1Id: 5, Spell2Id: 6, ChampionId: 109},
				{Spell1Id: 1, Spell2Id: 2, ChampionId: 110},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enemyData, _ := extractEnemyData(tt.participants)
			if !reflect.DeepEqual(enemyData, tt.expectedEnemyData) {
				t.Fatalf("expected enemy data: %v, got: %v", tt.expectedEnemyData, enemyData)
			}
		})
	}
}
