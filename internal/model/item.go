package model

import (
	"github.com/Piyanat1990/workflow/internal/constant"
)

type Item struct {
	ID       uint                `json:"id" gorm:"primaryKey"`
	Title    string              `json:"title"`
	Amount   float64             `json:"amount"`
	Quantity uint                `json:"quantity"`
	Status   constant.ItemStatus `json:"status"`
}
