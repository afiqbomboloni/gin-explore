package book

import "time"

type Book struct {
	Id          int64	
	Title       string	
	Description string	
	Price       int64		
	Rating      int64		
	CreatedAt   time.Time
	UpdatedAt 	time.Time

}