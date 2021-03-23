package main

import (
	"fmt"

	"github.com/en1tan/anon-board/config"
	"github.com/en1tan/anon-board/database"
	"github.com/en1tan/anon-board/routes"
)

func main() {
	c := config.NewConfig()
	r := routes.NewRouter(c)
	conn := database.NewDatabaseConnection(c)

	fmt.Println(conn)

	r.Serve()
}
