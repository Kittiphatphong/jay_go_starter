package paginates

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PaginateRequest struct {
	Item int `json:"item" validate:"required""`
	Page int `json:"page" validate:"required"`
}

type PaginatedResponse struct {
	Data        interface{} `json:"data"`
	Total       int         `json:"total"`
	PerPage     int         `json:"per_page"`
	CurrentPage int         `json:"current_page"`
	LastPage    int         `json:"last_page"`
}

func Paginate(db *gorm.DB, model interface{}, paginate PaginateRequest) (*PaginatedResponse, error) {
	var total int64
	db.Model(model).Count(&total)
	lastPage := (int(total) + paginate.Item - 1) / paginate.Item
	offset := (paginate.Page - 1) * paginate.Item
	result := db.Preload(clause.Associations).Limit(paginate.Item).Offset(offset).Find(model)
	if result.Error != nil {
		return nil, result.Error
	}
	pagination := &PaginatedResponse{
		Total:       int(total),
		PerPage:     paginate.Item,
		CurrentPage: paginate.Page,
		LastPage:    lastPage,
		Data:        model,
	}
	return pagination, nil
}
