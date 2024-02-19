package main

import (
	"os"

	"github.com/maxexq/parksoi-shop/config"
	"github.com/maxexq/parksoi-shop/modules/servers"
	"github.com/maxexq/parksoi-shop/pkg/databases"
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

	servers.NewServer(cfg, db).Start()

}
