package main

import (
	"github.com/daoraimi/devlab-backend/application"
	"github.com/daoraimi/devlab-backend/server"
)

func main() {
	s := server.New(&application.UserApp{})

	s.Run(":8000")
}
