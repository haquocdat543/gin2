package main

import (
	"gin/src/config"
	"gin/src/db"
	"gin/src/router"
)

func main() {

	config.LoadEnv()
	dbConn := config.InitDB()

	if err := db.Migrate(dbConn); err != nil {
		panic("Migration failed: " + err.Error())
	}

	r := router.SetupRouter(dbConn)
	r.Run(":8080")
}
