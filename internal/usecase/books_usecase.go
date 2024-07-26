package usecase

import (
	"github.com/afteroffice/go-example-rest/internal/model"
	"github.com/afteroffice/go-example-rest/internal/repo"
)

type BooksUCInterface interface {
	// CreateBook(req *dto.BookReqDTO) (*dto.GetBooksRespDTO, error)
	GetBooks(req *GetBookRequest) ([]model.Book, error)
	// UpdateBook(req *dto.BookReqDTO) (*dto.GetBooksRespDTO, error)
	// DeleteBook(req *dto.BookReqDTO) (*dto.GetBooksRespDTO, error)
}

type booksUseCase struct {
	repo repo.Repo
}

func NewBooksUseCase(r repo.Repo) *booksUseCase {
	return &booksUseCase{
		repo: r,
	}
}
