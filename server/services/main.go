package services

import "bs-to-scrapper/server/services/database"

func DB() *database.Database {
	return database.Init()
}
