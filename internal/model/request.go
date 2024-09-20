package model

import "github.com/Piyanat1990/workflow/internal/constant"

type RequestItem struct {
	Title    string  `binding:"required"`
	Amount    float64 
	Quantity uint

}

type RequestFindItem struct {
	Statuses constant.ItemStatus `form:"status"`
}

type RequestUpdateItem struct {
	Status constant.ItemStatus
}

type RequestUpdateItems struct{
	Title    string 
	Amount    float64 
	Quantity uint

}

type RequestLogin struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}
