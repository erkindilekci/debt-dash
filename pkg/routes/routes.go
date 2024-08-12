package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"

	"github.com/erkindilekci/debt-dash/pkg/handlers"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("static/"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	router.Get("/", handlers.GetAllCards)
	router.Get("/add-new-card", handlers.GetForm)
	router.Post("/add-new-card", handlers.CreateNewCard)
	router.Get("/add-expense/{id}", handlers.GetExpenseForm)
	router.Post("/add-expense/{id}", handlers.AddExpense)
	router.Post("/delete/{id}", handlers.DeleteCardById)
	router.Post("/delete-all-cards", handlers.DeleteAllCards)

	return router
}
