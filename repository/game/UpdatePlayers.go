package gamerepo

import (
	"strings"
)

func (s *GameRepository) UpdatePlayers(id int, players []string) error {
	_, err := s.db.Exec(
		`
		UPDATE game
		SET players = string_to_array($1, ',')
		WHERE id = $2
		`,
		strings.Join(players, ","),
		id,
	)
	return err
}