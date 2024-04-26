package dto

type AddBookRequest struct {
	Name            string `json:"name" validate:"required"`
	Author          string `json:"author" validate:"required"`
	PublicationYear int    `json:"publication_year" validate:"required"`
}
