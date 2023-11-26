package server

import (
	"fmt"
	"lib/controllers"
	"lib/database"
	"lib/utils"
	"log"
	"os"
)

type Server struct {
	ControllerInterface controllers.IController
	Storage             database.IDatabase
	ListenAddr          string
}

func NewServer() *Server {
	storage, err := database.NewStorage()
	if err != nil {
		log.Fatal(utils.ExceptionToString(err))
	}

	return &Server{
		ControllerInterface: controllers.BaseController{Storage: storage},
		Storage:             storage,
		ListenAddr:          fmt.Sprint(":", os.Getenv("LISTENADDR")),
	}
}
