package gamerepo

func (s *GameRepositary) CreateGame(name string) int {
	var id int

	err := s.db.QueryRow(
		`
		INSERT INTO game (name) 
		VALUES ($1) 
		RETURNING id
		`,
		name,
	).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}
