package services

import "bs-to-scrapper/server/services/db"

func DB() db.ServiceCollection {
	return db.ServiceCollection{}
}

func Network() NetworkService {
	return NetworkService{}
}

func JWT() JwtService {
	return JwtService{}
}

func User() UserService {
	return UserService{}
}

var timerService *TimerService

func Timer() *TimerService {

	if timerService == nil {
		timerService = &TimerService{}
	}

	return timerService
}
