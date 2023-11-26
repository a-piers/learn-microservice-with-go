package server

import (
	"fmt"
	"lib/cmd/api"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func (srv *Server) InitRouters(r *mux.Router) {
	r.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", os.Getenv("LISTENADDR"))),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	r.HandleFunc("/login", api.MakeHTTPHandleFunc(srv.ControllerInterface.Login)).Methods(http.MethodPost)
}
