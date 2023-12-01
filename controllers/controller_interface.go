package controllers

import "net/http"

type IController interface {
	Login(http.ResponseWriter, *http.Request) error
	Register(http.ResponseWriter, *http.Request) error
}
