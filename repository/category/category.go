package category

import (
	_entity "capstone/be/entity"
	"database/sql"
	"errors"
	"log"
	"net/http"
)

type CategoryRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (ur *CategoryRepository) GetAll() (categories []_entity.Category, code int, err error) {
	stmt, err := ur.db.Prepare(`
	select * from categories
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return categories, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query()

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return categories, code, err
	}

	defer res.Close()

	for res.Next() {
		category := _entity.Category{}

		if err := res.Scan(&category, &category.Name); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return categories, code, err
		}

		categories = append(categories, category)
	}

	return categories, http.StatusOK, nil
}
