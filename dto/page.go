package dto

import "github.com/Rhtymn/synapsis-challenge/domain"

type PageInfo struct {
	CurrentPage  int   `json:"current_page"`
	ItemsPerPage int   `json:"items_per_page"`
	ItemCount    int64 `json:"item_count"`
	PageCount    int   `json:"page_count"`
}

func NewPageInfoResponse(p domain.PageInfo) PageInfo {
	return PageInfo{
		CurrentPage:  p.CurrentPage,
		ItemsPerPage: p.ItemsPerPage,
		ItemCount:    p.ItemCount,
		PageCount:    p.PageCount,
	}
}
