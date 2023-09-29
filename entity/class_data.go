package entity

import (
	"github.com/crysalia/common/enums"
	"github.com/google/uuid"
)

type InventoryItem struct {
	Slot int64
	Item string
}

type ClassData struct {
	Id        uuid.UUID
	Player    string
	Class     enums.Class
	Level     int64
	XP        float64
	Inventory []string
}
