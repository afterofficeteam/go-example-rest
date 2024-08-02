package main

import (
	"github.com/afteroffice/go-example-rest/external/openlib"
	"github.com/afteroffice/go-example-rest/internal/handler"
	"github.com/afteroffice/go-example-rest/internal/repo"
	"github.com/afteroffice/go-example-rest/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// init modules
	openLib := openlib.NewOpenLibrary()
	bookResp, err := openLib.GetBooksBySubject("computer_science")
	if err != nil {
		panic(err)
	}

	// init repo
	r := repo.NewRepo()
	r.SeedInit(bookResp)

	// init usecase & handler
	uc := usecase.NewBooksUseCase(r)
	h := handler.NewHandler(uc)

	// start service
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/books", h.GetBooks)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
