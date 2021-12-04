package models

import "time"

type Product struct {
	Id        uint       `json:"id"`
	Name      string     `json:"name"`
	Price     uint       `json:"price"`
	Quantity  uint       `json:"quantity"`
	CreatedAt *time.Time `json:"created_at"`
}
