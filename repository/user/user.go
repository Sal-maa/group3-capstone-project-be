package user

import (
	_entity "capstone/be/entity"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) LoginByEmail(email string) (loginUser _entity.User, code int, err error) {
	id, err := ur.checkEmailExistence(email)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return loginUser, code, err
	}

	if id == 0 {
		log.Println("email not found while login by email")
		code, err = http.StatusBadRequest, errors.New("email not found")
		return loginUser, code, err
	}

	stmt, err := ur.db.Prepare(`
		SELECT role, name, password, avatar
		FROM users
		WHERE deleted_at IS NULL AND id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return loginUser, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return loginUser, code, err
	}

	defer res.Close()

	if res.Next() {
		if err = res.Scan(&loginUser.Role, &loginUser.Name, &loginUser.Password, &loginUser.Avatar); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return loginUser, code, err
		}
	}

	loginUser.Id = int(id)
	loginUser.Avatar = fmt.Sprintf("https://capstone-group3.s3.ap-southeast-1.amazonaws.com/%s", loginUser.Avatar)

	return loginUser, http.StatusOK, nil
}

func (ur *UserRepository) LoginByPhone(phone string) (loginUser _entity.User, code int, err error) {
	id, err := ur.checkPhoneExistence(phone)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return loginUser, code, err
	}

	if id == 0 {
		log.Println("phone not found while login by phone")
		code, err = http.StatusBadRequest, errors.New("phone not found")
		return loginUser, code, err
	}

	stmt, err := ur.db.Prepare(`
		SELECT role, name, password, avatar
		FROM users
		WHERE deleted_at IS NULL AND id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return loginUser, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return loginUser, code, err
	}

	defer res.Close()

	if res.Next() {
		if err = res.Scan(&loginUser.Role, &loginUser.Name, &loginUser.Password, &loginUser.Avatar); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return loginUser, code, err
		}
	}

	loginUser.Id = int(id)
	loginUser.Avatar = fmt.Sprintf("https://capstone-group3.s3.ap-southeast-1.amazonaws.com/%s", loginUser.Avatar)

	return loginUser, http.StatusOK, nil
}

func (ur *UserRepository) GetById(id int) (user _entity.User, code int, err error) {
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

func (ur *UserRepository) GetAll() (users []_entity.UserSimplified, code int, err error) {
	stmt, err := ur.db.Prepare(`
		SELECT u.id, d.name, u.role, u.name, u.email, u.phone, u.gender, u.address, u.avatar
		FROM users u
		JOIN divisions d
		ON u.division_id = d.id
		WHERE u.deleted_at IS NULL
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return users, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query()

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return users, code, err
	}

	defer res.Close()

	for res.Next() {
		user := _entity.UserSimplified{}

		if err := res.Scan(&user.Id, &user.Division, &user.Role, &user.Name, &user.Email, &user.Phone, &user.Gender, &user.Address, &user.Avatar); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return users, code, err
		}

		user.Avatar = fmt.Sprintf("https://capstone-group3.s3.ap-southeast-1.amazonaws.com/%s", user.Avatar)

		users = append(users, user)
	}

	return users, http.StatusOK, nil
}

func (ur *UserRepository) Update(userData _entity.User) (updatedUser _entity.UserSimplified, code int, err error) {
	id, err := ur.checkEmailExistence(userData.Email)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updatedUser, code, err
	}

	if id != 0 && id != int64(userData.Id) {
		log.Println("email already used by other user while update user")
		code, err = http.StatusBadRequest, errors.New("email already exist")
		return updatedUser, code, err
	}

	id, err = ur.checkPhoneExistence(userData.Phone)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updatedUser, code, err
	}

	if id != 0 && id != int64(userData.Id) {
		log.Println("phone already used by other user while update user")
		code, err = http.StatusBadRequest, errors.New("phone already exist")
		return updatedUser, code, err
	}

	id, err = ur.getDivisionId(userData.Division)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updatedUser, code, err
	}

	if id == 0 {
		log.Println("division id not found while update user")
		code, err = http.StatusBadRequest, errors.New("division not found")
		return updatedUser, code, err
	}

	stmt, err := ur.db.Prepare(`
		UPDATE users
		SET division_id = ?, name = ?, email = ?, phone = ?, password = ?, gender = ?, address = ?, avatar = ?
		WHERE deleted_at IS NULL AND id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updatedUser, code, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(id, userData.Name, userData.Email, userData.Phone, userData.Password, userData.Gender, userData.Address, userData.Avatar, userData.Id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updatedUser, code, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updatedUser, code, err
	}

	if rowsAffected == 0 {
		log.Println("rows affected is 0 while update user")
		code, err = http.StatusBadRequest, errors.New("user not updated")
		return updatedUser, code, err
	}

	stmt, err = ur.db.Prepare(`
		UPDATE users
		SET updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updatedUser, code, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userData.Id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updatedUser, code, err
	}

	updatedUser.Id = userData.Id
	updatedUser.Division = userData.Division
	updatedUser.Role = userData.Role
	updatedUser.Name = userData.Name
	updatedUser.Email = userData.Email
	updatedUser.Phone = userData.Phone
	updatedUser.Gender = userData.Gender
	updatedUser.Address = userData.Address
	updatedUser.Avatar = fmt.Sprintf("https://capstone-group3.s3.ap-southeast-1.amazonaws.com/%s", userData.Avatar)

	return updatedUser, http.StatusOK, nil
}

func (ur *UserRepository) checkEmailExistence(email string) (id int64, err error) {
	stmt, err := ur.db.Prepare(`
		SELECT id
		FROM users
		WHERE deleted_at IS NULL AND email = ?
	`)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Query(email)

	if err != nil {
		return 0, err
	}

	defer res.Close()

	if res.Next() {
		if err = res.Scan(&id); err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (ur *UserRepository) checkPhoneExistence(phone string) (id int64, err error) {
	stmt, err := ur.db.Prepare(`
		SELECT id
		FROM users
		WHERE deleted_at IS NULL AND phone = ?
	`)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Query(phone)

	if err != nil {
		return 0, err
	}

	defer res.Close()

	if res.Next() {
		if err = res.Scan(&id); err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (ur *UserRepository) getDivisionId(division string) (id int64, err error) {
	stmt, err := ur.db.Prepare(`
		SELECT id
		FROM divisions
		WHERE name = ?
	`)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Query(division)

	if err != nil {
		return 0, err
	}

	defer res.Close()

	if res.Next() {
		if err = res.Scan(&id); err != nil {
			return 0, err
		}
	}

	return id, nil
}
