package repo

import (
	"github.com/afteroffice/go-example-rest/external/openlib"
	"github.com/afteroffice/go-example-rest/internal/model"
	"github.com/google/uuid"
)

type RepoInterface interface {
	// GetBooks(req *dto.BookReqDTO) (*dto.GetBooksRespDTO, error)
}

type Repo struct {
	datastore []model.Book
}

func NewRepo() Repo {
	return Repo{
		datastore: []model.Book{},
	}
}

func (r *Repo) SeedInit(bookResp *openlib.GetBooksRespDTO) {
	for _, work := range bookResp.Works {
		newBook := model.Book{
			Id:      uuid.New(),
			Title:   work.Title,
			URL:     work.Key,
			Authors: []model.Author{},
		}

		for _, author := range work.Authors {
			newBook.Authors = append(newBook.Authors, model.Author{
				Name: author.Name,
				URL:  author.Key,
			})
		}

		r.datastore = append(r.datastore, newBook)
	}
}
