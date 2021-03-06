package repository

import (
	"coinche/domain"
)

func (s *GameRepository) CreateGame(game domain.Game) (int, error) {
	tx := s.db.MustBegin()

	var gameID int
	err := tx.QueryRow(
		`
		INSERT INTO game (name) 
		VALUES ($1) 
		RETURNING id
		`,
		game.Name,
	).Scan(&gameID)
	if err != nil {
		return 0, err
	}

	for playerName, player := range game.Players {
		_, err := tx.Exec(
			`
			INSERT INTO player (name, team, gameid) 
			VALUES ($1, $2, $3)
			`,
			playerName,
			player.Team,
			gameID,
		)
		if err != nil {
			return 0, err
		}
	}

	return gameID, tx.Commit()
}
