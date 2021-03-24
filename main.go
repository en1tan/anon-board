package main

import (
	"github.com/en1tan/anon-board/config"
	"github.com/en1tan/anon-board/controllers"
	"github.com/en1tan/anon-board/database"
	"github.com/en1tan/anon-board/routes"
	"github.com/en1tan/anon-board/services"
)

func main() {
	c := config.NewConfig()

	r := routes.NewRouter(c)
	conn := database.NewDatabaseConnection(c)

	ts := services.NewThreadService(conn)
	tc := controllers.NewThreadController(ts)

	r.RegisterThreadRoutes(tc)

	r.Serve()
}
