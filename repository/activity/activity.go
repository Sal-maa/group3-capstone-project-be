package activity

import (
	_entity "capstone/be/entity"

	"database/sql"
	"net/http"
)

type ActivityRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *ActivityRepository {
	return &ActivityRepository{db: db}
}

func (ar ActivityRepository) GetAll() (activities []_entity.ActivitySimplified, code int, err error) {

	return activities, http.StatusOK, nil
}
