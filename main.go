package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/oshjoshu99/simplebank/api"
	db "github.com/oshjoshu99/simplebank/db/sqlc"
	"github.com/oshjoshu99/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(testDB)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Run(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
