package models

import "time"

type CreateAdminTimerRequest struct {
	OrderTime          time.Time `json:"orderTime"`
	YoutastePhone      string    `json:"youtastePhone"`
	YoutastePassword   string    `json:"youtastePassword"`
	LieferandoUsername string    `json:"lieferandoUsername"`
	LieferandoPassword string    `json:"lieferandoPassword"`
}
