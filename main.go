package main

import (
	"fmt"
	"os"

	"github.com/Worrameth/maimeebooks/config"
	"github.com/Worrameth/maimeebooks/modules/servers"
	"github.com/Worrameth/maimeebooks/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())

	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	fmt.Println(db)

	servers.NewServer(cfg, db).Start()
}
