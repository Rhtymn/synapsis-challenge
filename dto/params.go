package dto

type IDParams struct {
	ID int64 `uri:"id" binding:"required,numeric,min=1"`
}
