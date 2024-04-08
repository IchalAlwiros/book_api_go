package model

import "time"

type Book struct {
	ID        uint    `json:id gorm:primaryKey`
	Name      string  `json:name`
	Author    *string `json:author`
	CreatedAt time.Time `json:created_at`
	UpdatedAt time.Time `json:updated_at`
}