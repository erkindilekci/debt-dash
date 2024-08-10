package db

import (
	"database/sql"
	"fmt"
	"github.com/erkindilekci/debt-dash/internal/models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() error {
	var err error

	db, err = sql.Open("postgres", "YOUR_POSTGRES_DATABASE_URL")
	if err != nil {
		return err
	}

	sqlStatement := `CREATE TABLE IF NOT EXISTS
		cards (
			id SERIAL PRIMARY KEY,
			name TEXT,
			card_limit INTEGER,
			current_debt NUMERIC,
			minimum_debt NUMERIC
		)`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil
}

func DbCreateCard(card *models.CreditCard) error {
	query, err := db.Prepare("INSERT INTO cards(name, card_limit, current_debt, minimum_debt) VALUES ($1, $2, $3, $4)")
	defer query.Close()

	if err != nil {
		return err
	}

	_, err = query.Exec(card.Name, card.Limit, card.CurrentTermDebt, card.MinimumDebt)

	if err != nil {
		return err
	}

	return nil
}

func DbFindNameById(id int) (string, error) {
	query, err := db.Prepare("SELECT name FROM cards WHERE id = $1")
	defer query.Close()

	var name string

	row := query.QueryRow(id)
	err = row.Scan(&name)
	if err != nil {
		return name, err
	}

	return name, nil
}

func DbGetAllCards() ([]*models.CreditCard, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query, err := db.Prepare("SELECT id, name, card_limit, current_debt, minimum_debt FROM cards")
	if err != nil {
		return nil, err
	}
	defer func() {
		if query != nil {
			query.Close()
		}
	}()

	result, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer result.Close()

	cards := make([]*models.CreditCard, 0)
	for result.Next() {
		data := new(models.CreditCard)
		err := result.Scan(&data.Id, &data.Name, &data.Limit, &data.CurrentTermDebt, &data.MinimumDebt)
		if err != nil {
			return nil, err
		}
		cards = append(cards, data)
	}
	return cards, nil
}

func DbAddExpense(id int, expense float64) error {
	query, err := db.Prepare(`
		UPDATE cards
		SET current_debt = current_debt + $1,
			minimum_debt = minimum_debt +
				CASE WHEN card_limit >= 25000 THEN $2 * 0.4 ELSE $3 * 0.3 END
		WHERE id = $4`)

	defer query.Close()

	if err != nil {
		return err
	}

	_, err = query.Exec(expense, expense, expense, id)

	if err != nil {
		return err
	}

	return nil
}

func DbDeleteCardById(id int) error {
	query, err := db.Prepare("DELETE FROM cards WHERE id=$1")
	defer query.Close()

	if err != nil {
		return err
	}

	_, err = query.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func DbDeleteAllCards() error {
	query, err := db.Prepare("DELETE FROM cards")
	defer query.Close()

	if err != nil {
		return err
	}

	_, err = query.Exec()
	if err != nil {
		return err
	}

	return nil
}
