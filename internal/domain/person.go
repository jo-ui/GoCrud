package domain

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/google/uuid"
)

type Person struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name    string    `json:"name" binding:"required"`
	Age     int       `json:"age" binding:"required"`
	Hobbies Hobbies   `json:"hobbies" gorm:"type:json;serializer:json"`
}

type Hobbies []string

func (h *Hobbies) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), h)
}

func (h Hobbies) Value() (driver.Value, error) {
	return json.Marshal(h)
}
