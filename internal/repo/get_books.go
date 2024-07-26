package repo

import (
	"github.com/afteroffice/go-example-rest/internal/model"
)

type GetBooksParam struct {
	Limit int
}

func (r *Repo) GetBooks(param *GetBooksParam) []model.Book {
	return r.datastore[0:param.Limit]
}
