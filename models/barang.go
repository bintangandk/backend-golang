package models

import "time"

type Barang struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Nama      string    `json:"nama"`
	Harga     float64   `json:"harga"`
	Stok      int       `json:"stok"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
