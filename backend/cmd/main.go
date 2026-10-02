package main

import (
	"github.com/kingofhandsomes/silkscreen-website/libs/config"
	"github.com/kingofhandsomes/silkscreen-website/libs/db"
)

func main() {
	cfg := config.MustLoad("")

	pool := db.MustPool(
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Database,
	)
	defer pool.Close()

}
