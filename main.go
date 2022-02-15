package main

import (
	"coinche/api"
	gamerepo "coinche/repository/game"
	"coinche/usecases"
	"coinche/utilities/env"
	"fmt"
	"os"
)

func main() {
	env.LoadEnv("")
	connectionInfo := os.Getenv("SQLX_POSTGRES_INFO")
	dbName := os.Getenv("DB_NAME")
	addr := os.Getenv("PORT")

	dsn := connectionInfo + " dbname=" + dbName
	gameRepository, err := gamerepo.NewGameRepository(dsn)
	if err != nil {
		panic(err)
	}
	gameUsecases := usecases.NewGameUsecases(gameRepository)

	router, _ := api.SetupRouter(gameUsecases)

	fmt.Println("Listening on ", addr)
	err = router.Run(addr)
	if err != nil {
		panic(err)
	}
}
