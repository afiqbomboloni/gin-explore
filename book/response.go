package book

type BookResponse struct {
	Id          int64		`json:"id"`
	Title       string		`json:"title"`
	Description string		`json:"description"`
	Price       int64		`json:"price"`
	Rating      int64		`json:"rating"`
}