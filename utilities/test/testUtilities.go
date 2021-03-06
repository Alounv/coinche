package testUtilities

import (
	"bytes"
	"coinche/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
)

func NewCreateGameRequest(test *testing.T, name string) *http.Request {
	route := fmt.Sprintf("/games/create?name=%s", url.QueryEscape(name))
	return GetNewRequest(test, route, http.MethodPost)
}

func NewGetGameRequest(test *testing.T, gameID int) *http.Request {
	route := fmt.Sprintf("/games/%d", gameID)
	return GetNewRequest(test, route, http.MethodGet)
}

func NewJoinGameRequest(test *testing.T, gameID int, playerName string) *http.Request {
	route := fmt.Sprintf("/games/%d/join?playerName=%s", gameID, url.QueryEscape(playerName))
	return GetNewRequest(test, route, http.MethodGet)
}

func GetNewRequest(test *testing.T, route string, method string) *http.Request {
	request, err := http.NewRequest(method, route, nil)
	FatalIfErr(err, test)
	return request
}

func CreateDb(connectionInfo string, dbName string) *sqlx.DB {
	userDb := sqlx.MustOpen("pgx", connectionInfo)
	_, err := userDb.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		fmt.Println("Database already existing, drop before creation", err)
		userDb.MustExec("DROP DATABASE " + dbName)
		userDb.MustExec("CREATE DATABASE " + dbName)
	}
	userDb.Close()

	db := sqlx.MustOpen("pgx", connectionInfo+" dbname="+dbName)
	return db
}

func DropDb(connectionInfo string, dbName string, db *sqlx.DB) {
	db.Close()

	userDb := sqlx.MustOpen("pgx", connectionInfo)
	userDb.MustExec("DROP DATABASE " + dbName)
	userDb.Close()
}

func DecodeToGames(buf *bytes.Buffer, test *testing.T) []domain.Game {
	var got []domain.Game
	err := json.NewDecoder(buf).Decode(&got)
	if err != nil {
		test.Fatalf("Unable to parse response from gameAPIs %q into %q, '%v'", buf, "slice of Game", err)
	}
	return got
}

func DecodeToGame(buf *bytes.Buffer, test *testing.T) domain.Game {
	var got domain.Game
	err := json.NewDecoder(buf).Decode(&got)
	if err != nil {
		test.Fatalf("Unable to parse response from gameAPIs %q into %q, '%v'", buf, "Game", err)
	}
	return got
}

func FatalIfErr(err error, test *testing.T) {
	if err != nil {
		test.Fatal(err)
	}
}
