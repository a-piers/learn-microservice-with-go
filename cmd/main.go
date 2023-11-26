package main

import (
	"lib/cmd/server"
	"lib/utils"
	"log"
	"net/http"

	_ "lib/docs"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(utils.ExceptionToString(err))
	}

	r := mux.NewRouter()

	service := server.NewServer()

	err := service.Storage.Ping()
	if err != nil {
		log.Fatalf("Error on pinging database. %s", utils.ExceptionToString(err))
	}

	defer func() {
		if err := service.Storage.Close(); err != nil {
			log.Fatalf("Error on closing database. %s", utils.ExceptionToString(err))
		}
	}()

	service.InitRouters(r)

	log.Printf("Server is listening on %s address...", service.ListenAddr)
	log.Fatal(http.ListenAndServe(service.ListenAddr, r))
}
