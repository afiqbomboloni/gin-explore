package book


type Service interface {
	FindAll() ([]Book, error)
	FindById(Id int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(Id int, bookUpdateRequest BookUpdateRequest) (Book, error)
	Delete(Id int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func(s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()

	return books,err

}

func(s *service) FindById(Id int) (Book, error) {
	book, err := s.repository.FindById(Id)
	return book,err
}

func(s *service) Create(bookRequest BookRequest) (Book, error) {

	price,_ := bookRequest.Price.Int64()
	rating,_ := bookRequest.Rating.Int64()

	book := Book {
		Title: bookRequest.Title,
		Price: price,
		Rating: rating,
		Description: bookRequest.Description,
	}

	newBook, err := s.repository.Create(book)

	return newBook, err


}

func(s *service) Update(Id int, bookUpdateRequest BookUpdateRequest) (Book, error) {

	book, err := s.repository.FindById(Id)
	price,_ := bookUpdateRequest.Price.Int64()
	rating,_ := bookUpdateRequest.Rating.Int64()

	
	book.Title = bookUpdateRequest.Title
	book.Price =  price
	book.Rating =  rating
	book.Description =  bookUpdateRequest.Description
	

	updatedBook, err := s.repository.Update(book)

	return updatedBook, err


}

func(s *service) Delete(Id int) (Book, error) {
	book, err := s.repository.FindById(Id)
	

	deletedBook, err := s.repository.Delete(book)

	return deletedBook,err
}