package model

import (
	"database/sql"
	"time"
)

type Inventory struct {
	ID        int            `json:"id"`
	ProductId sql.NullString `json:"productId"`
	Quantity  int            `json:"quantity"`
	Location  sql.NullString `json:"location"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
}

type InventoryDto struct {
	ID        int       `json:"id"`
	ProductId *string   `json:"productId"`
	Quantity  int       `json:"quantity"`
	Location  *string   `json:"location"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (i *InventoryDto) FillFromModel(model Inventory) {
	i.ID = model.ID
	if model.ProductId.Valid {
		i.ProductId = &model.ProductId.String
	}
	i.Quantity = model.Quantity
	if model.Location.Valid {
		i.Location = &model.Location.String
	}
	i.CreatedAt = model.CreatedAt
	i.UpdatedAt = model.UpdatedAt
}

func (i InventoryDto) ToModel() Inventory {
	model := Inventory{
		ID:        i.ID,
		Quantity:  i.Quantity,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
	if i.ProductId != nil {
		model.ProductId.String = *i.ProductId
		model.ProductId.Valid = true
	}
	if i.Location != nil {
		model.Location.String = *i.Location
		model.Location.Valid = true
	}

	return model
}
