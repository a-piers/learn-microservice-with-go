package controllers

import "lib/database"

type BaseController struct {
	Storage database.IDatabase
}
