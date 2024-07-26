package usecase

import (
	"github.com/afteroffice/go-example-rest/internal/model"
	"github.com/afteroffice/go-example-rest/internal/repo"
)

type GetBookRequest struct {
	Limit int
}

func (uc *booksUseCase) GetBooks(req *GetBookRequest) ([]model.Book, error) {
	return uc.repo.GetBooks(&repo.GetBooksParam{
		Limit: req.Limit,
	}), nil
}
