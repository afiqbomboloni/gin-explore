package book

import "encoding/json"

type BookRequest struct {
	Title string `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
	Description string `json:"description" binding:"required"`
	Rating json.Number `json:"rating" binding:"required"`
}

type BookUpdateRequest struct {
	Title string `json:"title"`
	Price json.Number `json:"price" binding:"number"`
	Description string `json:"description"`
	Rating json.Number `json:"rating"`
}

