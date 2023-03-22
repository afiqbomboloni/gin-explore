package handler

import (
	"fmt"
	"golang-api/book"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func(h *bookHandler) GetBooks(ctx *gin.Context) {
	books, err := h.bookService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _,b := range books {
		bookResponse := convertToResponse(b)

		booksResponse = append(booksResponse, bookResponse)
	}

	ctx.JSON(http.StatusOK,gin.H{
		"data": booksResponse,
	})
}

func(h *bookHandler) GetBook(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	bookResponse := convertToResponse(b)
	if b.Id == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "No Content",
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusOK,gin.H{
			"data": bookResponse,
		})
	}

	

	

}	


func(h *bookHandler) CreateBook(ctx *gin.Context) {
	var bookRequest book.BookRequest

	err := ctx.ShouldBindJSON(&bookRequest)

	if err != nil {

		errorMessages := []string{}
		for _,e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, is %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
			
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something error",
		})
	}

	bookResponse := convertToResponse(book)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": bookResponse,
	})
}

func(h *bookHandler) UpdateBook(ctx *gin.Context) {
	

	var bookUpdateRequest book.BookUpdateRequest

	err := ctx.ShouldBind(&bookUpdateRequest)

	
	

	if err != nil {

		errorMessages := []string{}
		for _,e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, is %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
			return
			
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error": errorMessages,
		})

	}


	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := h.bookService.Update(id, bookUpdateRequest)


	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Id Not Found",
		})

		return
	}

	bookResponse := convertToResponse(book)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": bookResponse,
	})
}


func(h *bookHandler) DeleteBook(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Delete(id)
	if book.Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":"Id Not found",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	

	bookResponse:= convertToResponse(book)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data" : bookResponse,
	})
}

func convertToResponse(b book.Book) book.BookResponse{
	return book.BookResponse {
		Id: b.Id,
		Title: b.Title,
		Price: b.Price,
		Description: b.Description,
		Rating: b.Rating,
}
}