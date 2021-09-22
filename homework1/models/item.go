package models

import "time"

//Item - struct
type Item struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"` // цена в копеках,чтобы не было плавающих чисел
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
