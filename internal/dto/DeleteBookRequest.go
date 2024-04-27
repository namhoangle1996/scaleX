package dto

type DeleteBookRequest struct {
	Name string `json:"name" validate:"required"`
}
