package openlib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type GetBooksRespDTO struct {
	Name        string     `json:"name"`
	SubjectType string     `json:"subject_type"`
	Works       []*WorkDTO `json:"works"`
}

type WorkDTO struct {
	Title        string       `json:"title"`
	CoverID      int64        `json:"cover_id"`
	EditionCount int64        `json:"edition_count"`
	Key          string       `json:"key"`
	Authors      []*AuthorDTO `json:"authors"`
}

type AuthorDTO struct {
	Name string `json:"name`
	Key  string `json:"key"`
}

type openLibraryService struct {
}

type OpenLibraryServices interface {
	GetBooksBySubject(subject string) (*GetBooksRespDTO, error)
}

func NewOpenLibrary() OpenLibraryServices {
	return &openLibraryService{}
}

func (s *openLibraryService) GetBooksBySubject(subject string) (*GetBooksRespDTO, error) {
	var response GetBooksRespDTO

	url := "http://openlibrary.org/subjects/%s.json?"

	url = fmt.Sprintf(url, subject)
	// fmt.Println(url, "here")

	getReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request to Open Library : %v", err)
	}

	getReq.Header["Accept"] = []string{"application/json"}
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Do(getReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request to Open Library: %v", err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println(err)
		bodyBytes, err := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		return nil, fmt.Errorf("failed to decode response: %v [RESP BODY: %s]", err, bodyString)
	}

	return &response, nil
}
