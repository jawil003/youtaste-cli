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

func Timer() *TimerService {
	return &TimerService{}
}

func Time() *TimeService {
	return &TimeService{}
}

func Scrapper() *ScrapperService {
	return &ScrapperService{}
}
