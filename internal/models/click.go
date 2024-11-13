package models

import "time"

type Click struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	BannerID  string    `json:"banner_id" db:"banner_id"`
}
