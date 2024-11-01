package models

import (
	"time"

	"github.com/Danila331/rushh/app/store"
)

type CardApplication struct {
	ID     int
	UserID int
	Date   time.Time
	Sum    int
	Done   bool
}

type CardApplicationInteface interface {
	Create() error
	Update() error
	ReadById() (CardApplication, error)
	ReadAll() ([]CardApplication, error)
	Delete() error
}

func (c *CardApplication) Create() error {
	connect, err := store.ConnectDB()
	if err != nil {
		return err
	}
	defer connect.Close()
	_, err = connect.Exec("INSERT INTO cards (user_id, date, sum, done) VALUES ($1, $2, $3, $4)", c.UserID, c.Date, c.Sum, c.Done)
	if err != nil {
		return err
	}
	return nil
}

func (c *CardApplication) Update() error {
	connect, err := store.ConnectDB()
	if err != nil {
		return err
	}
	defer connect.Close()
	_, err = connect.Exec("UPDATE cards SET user_id = $1, date = $2, sum = $3, done = $4 WHERE id = $5", c.UserID, c.Date, c.Sum, c.Done, c.ID)
	if err != nil {
		return err
	}
	return nil
}

func (c *CardApplication) ReadById() (CardApplication, error) {
	connect, err := store.ConnectDB()
	if err != nil {
		return CardApplication{}, err
	}
	defer connect.Close()
	var card CardApplication

	err = connect.QueryRow("SELECT * FROM cards WHERE id = $1", c.ID).Scan(&card.ID, &card.UserID, &card.Date, &card.Sum, &card.Done)
	if err != nil {
		return CardApplication{}, err
	}
	return card, nil
}

func (c *CardApplication) ReadAll() ([]CardApplication, error) {
	connect, err := store.ConnectDB()
	if err != nil {
		return []CardApplication{}, err
	}
	defer connect.Close()
	var cards []CardApplication
	rows, err := connect.Query("SELECT * FROM cards")
	if err != nil {
		return []CardApplication{}, err
	}
	for rows.Next() {
		var card CardApplication
		err = rows.Scan(&card.ID, &card.UserID, &card.Date, &card.Sum, &card.Done)
		if err != nil {
			return []CardApplication{}, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func (c *CardApplication) Delete() error {
	connect, err := store.ConnectDB()
	if err != nil {
		return err
	}
	defer connect.Close()
	_, err = connect.Exec("DELETE FROM cards WHERE id = $1", c.ID)
	if err != nil {
		return err
	}
	return nil
}
