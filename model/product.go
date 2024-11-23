package model

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Description sql.NullString  `json:"description"`
	Price       decimal.Decimal `json:"price"`
	Category    string          `json:"category"`
	ImagePath   sql.NullString  `json:"image_path"`
	CreatedAt   time.Time       `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `json:"updatedAt" gorm:"autoUpdateTime"`
}

type ProductDTO struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Price       int       `json:"price"`
	Category    string    `json:"category"`
	ImagePath   *string   `json:"imagePath"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}

// FillFromModel: Mengisi ProductDTO dari Product
func (p *ProductDTO) FillFromModel(model Product) {
	p.ID = model.ID
	p.Name = model.Name
	p.Description = ToPointerString(model.Description)
	p.Price = int(model.Price.IntPart())
	p.Category = model.Category
	p.ImagePath = ToPointerString(model.ImagePath)
	p.CreatedAt = model.CreatedAt
	p.UpdatedAt = model.UpdatedAt
}

// ToModel: Mengonversi ProductDTO ke Product
func (p ProductDTO) ToModel() Product {
	return Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: ToNullString(p.Description),
		Price:       decimal.NewFromInt(int64(p.Price)),
		Category:    p.Category,
		ImagePath:   ToNullString(p.ImagePath),
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

// Fungsi utilitas
func ToPointerString(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func ToNullString(ptr *string) sql.NullString {
	if ptr != nil {
		return sql.NullString{String: *ptr, Valid: true}
	}
	return sql.NullString{Valid: false}
}
