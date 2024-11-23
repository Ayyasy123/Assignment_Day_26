package model

import (
	"database/sql"
	"time"
)

type Order struct {
	ID        int            `json:"id"`
	ProductId sql.NullString `json:"productId"`
	Quantity  int            `json:"quantity"`
	OrderDate string         `json:"orderDate"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
}

type OrderDto struct {
	ID        int       `json:"id"`
	ProductId *string   `json:"productId" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required"`
	OrderDate string    `json:"orderDate"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (o *OrderDto) FillFromModel(model Order) {
	o.ID = model.ID
	if model.ProductId.Valid {
		o.ProductId = &model.ProductId.String
	}
	o.Quantity = model.Quantity
	o.OrderDate = model.OrderDate
	o.CreatedAt = model.CreatedAt
	o.UpdatedAt = model.UpdatedAt
}

func (o OrderDto) ToModel() Order {
	model := Order{
		ID:        o.ID,
		Quantity:  o.Quantity,
		OrderDate: o.OrderDate,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
	if o.ProductId != nil {
		model.ProductId.String = *o.ProductId
		model.ProductId.Valid = true
	}

	return model
}
