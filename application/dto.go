package application

type CreateBookDTO struct {
	Title string `json:"title" validate:"required"`
}

type BookDTO struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
