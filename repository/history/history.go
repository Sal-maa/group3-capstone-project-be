package history

import (
	_entity "capstone/be/entity"

	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type HistoryRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *HistoryRepository {
	return &HistoryRepository{db: db}
}

func (ur *HistoryRepository) GetById(id int) (user _entity.User, code int, err error) {
	stmt, err := ur.db.Prepare(`
		SELECT d.name, u.role, u.name, u.email, u.phone, u.password, u.gender, u.address, u.avatar
		FROM users u
		JOIN divisions d
		ON u.division_id = d.id
		WHERE u.deleted_at IS NULL AND u.id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return user, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return user, code, err
	}

	defer res.Close()

	if res.Next() {
		if err := res.Scan(&user.Division, &user.Role, &user.Name, &user.Email, &user.Phone, &user.Password, &user.Gender, &user.Address, &user.Avatar); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return user, code, err
		}
	}

	if user == (_entity.User{}) {
		log.Println("user not found")
		code, err = http.StatusBadRequest, errors.New("user not found")
		return user, code, err
	}

	user.Id = id
	user.Avatar = fmt.Sprintf("https://capstone-group3.s3.ap-southeast-1.amazonaws.com/%s", user.Avatar)

	return user, http.StatusOK, nil
}
