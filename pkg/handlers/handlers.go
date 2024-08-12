package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"

	"github.com/erkindilekci/debt-dash/pkg/db"
	"github.com/erkindilekci/debt-dash/pkg/models"
	"github.com/erkindilekci/debt-dash/pkg/render"
)

func GetAllCards(w http.ResponseWriter, r *http.Request) {
	type IndexPageData struct {
		Cards                []*models.CreditCard
		TotalDebt            float64
		TotalCurrentTermDebt float64
	}

	cards, err := db.DbGetAllCards()
	if err != nil {
		Catch(err)
	}

	var totalMinimumDebt float64
	var totalCurrentTermDebt float64
	for _, card := range cards {
		totalMinimumDebt += card.MinimumDebt
		totalCurrentTermDebt += card.CurrentTermDebt
	}

	totalMinimumDebt, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", totalMinimumDebt), 64)
	totalCurrentTermDebt, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", totalCurrentTermDebt), 64)

	render.RenderTemplate(w, "index.page.gohtml", IndexPageData{cards, totalMinimumDebt, totalCurrentTermDebt})
}

func CreateNewCard(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("cardName")
	cardLimit, _ := strconv.Atoi(r.FormValue("cardLimit"))
	currentTermDebt, _ := strconv.ParseFloat(r.FormValue("currentTermDebt"), 64)

	var minimumDebt float64
	if cardLimit >= 25000 {
		minimumDebt = currentTermDebt * 0.4
	} else {
		minimumDebt = currentTermDebt * 0.3
	}

	minimumDebt, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", minimumDebt), 64)

	newCard := models.CreditCard{
		Name:            name,
		Limit:           cardLimit,
		CurrentTermDebt: currentTermDebt,
		MinimumDebt:     minimumDebt,
	}

	err := db.DbCreateCard(&newCard)
	Catch(err)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func AddExpense(w http.ResponseWriter, r *http.Request) {
	cardId := chi.URLParam(r, "id")
	if cardId == "" {
		http.Error(w, "No id specified", http.StatusBadRequest)
		return
	}

	intId, err := strconv.Atoi(cardId)
	if err != nil {
		Catch(err)
	}

	expense, _ := strconv.ParseFloat(r.FormValue("expense"), 64)
	err = db.DbAddExpense(intId, expense)
	Catch(err)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeleteCardById(w http.ResponseWriter, r *http.Request) {
	cardId := chi.URLParam(r, "id")
	if cardId == "" {
		http.Error(w, "No id specified", http.StatusBadRequest)
		return
	}

	intId, err := strconv.Atoi(cardId)
	if err != nil {
		Catch(err)
	}

	err = db.DbDeleteCardById(intId)
	if err != nil {
		Catch(err)
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteAllCards(w http.ResponseWriter, r *http.Request) {
	err := db.DbDeleteAllCards()
	if err != nil {
		Catch(err)
	}

	w.WriteHeader(http.StatusOK)
}

func GetForm(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "form.page.gohtml", nil)
}

func GetExpenseForm(w http.ResponseWriter, r *http.Request) {
	cardId := chi.URLParam(r, "id")
	if cardId == "" {
		http.Error(w, "No id specified", http.StatusBadRequest)
		return
	}

	intId, _ := strconv.Atoi(cardId)

	name, err := db.DbFindNameById(intId)
	if err != nil {
		Catch(err)
	}

	render.RenderTemplate(w, "expense-form.page.gohtml", name)
}

func Catch(err error) {
	if err != nil {
		log.Println(err)
	}
}
