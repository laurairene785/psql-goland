package main

import (
	"errors"
	"time"
)

type ebook struct {
	id    int
	name  string
	date  time.Time
	study bool
}

// CreateEbook creates a new ebook in the DB
func CreateEbook(e *ebook) error {
	q := `INSERT INTO 
					ebooks (name, date, study)
					VALUES ($1, $2, $3)`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(e.name, e.date, e.study)
	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("We were specting 1 row afected")

	}
	return nil

}
