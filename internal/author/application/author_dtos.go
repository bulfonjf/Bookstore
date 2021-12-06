package application

type AuthorDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CreateAuthorDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateAuthorDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
