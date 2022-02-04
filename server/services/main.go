package services

import "bs-to-scrapper/server/services/db"

func DB() db.ServiceCollection {
	return db.ServiceCollection{}
}

func Network() NetworkService {
	return NetworkService{}
}
