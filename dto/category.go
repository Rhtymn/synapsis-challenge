package dto

import "github.com/Rhtymn/synapsis-challenge/domain"

type CategoryDTO struct {
	ID       int64   `json:"id,omitempty"`
	ImageURL *string `json:"image_url,omitempty"`
	Name     string  `json:"name,omitempty"`
	Slug     string  `json:"slug,omitempty"`
	ParentID *int64  `json:"parent_id,omitempty"`
}

func NewCategoryDTO(c domain.Category) CategoryDTO {
	return CategoryDTO{
		ID:       c.ID,
		ImageURL: c.ImageURL,
		Name:     c.Name,
		Slug:     c.Slug,
		ParentID: c.ParentID,
	}
}
