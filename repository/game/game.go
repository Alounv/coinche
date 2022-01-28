package gamerepo

import (
	"coinche/usecases"
	"fmt"

	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
)

var gameSchema = `
CREATE TABLE game (
	id serial PRIMARY KEY NOT NULL,
	name text,
	createdAt timestamp NOT NULL DEFAULT now(),
	players text[]
)`

type GameRepository struct {
	usecases.GameRepositoryInterface
	db *sqlx.DB
}

func (s *GameRepository) CreatePlayerTableIfNeeded() {
	_, err := s.db.Exec(gameSchema)
	if err != nil {
		fmt.Print(err)
	}
}

func NewGameRepository(dsn string) *GameRepository {
	db := sqlx.MustOpen("pgx", dsn)

	return NewGameRepositoryFromDb(db)
}

func NewGameRepositoryFromDb(db *sqlx.DB) *GameRepository {
	gameRepository := GameRepository{db: db}
	gameRepository.CreatePlayerTableIfNeeded()

	return &gameRepository
}
