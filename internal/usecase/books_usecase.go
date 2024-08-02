package usecase

import (
	"github.com/afteroffice/go-example-rest/internal/model"
	"github.com/afteroffice/go-example-rest/internal/repo"
)

type BooksUCInterface interface {
	GetBooks(req *GetBookRequest) ([]model.Book, error)
}

type booksUseCase struct {
	repo repo.Repo
}

func NewBooksUseCase(r repo.Repo) *booksUseCase {
	return &booksUseCase{
		repo: r,
	}
}
