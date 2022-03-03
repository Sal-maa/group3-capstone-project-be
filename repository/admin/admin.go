package admin

import (
	_entity "capstone/be/entity"
	"database/sql"
	"errors"
	"log"
	"net/http"
)

type AdminRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (ar *AdminRepository) GetAllAdmin(limit, offset int, status, category, date string) (requests []_entity.RequestResponse, total int, err error) {
	query := ""
	if category == "all" {
		category = ""
	}

	if status == "all" {
		status = "%%"
	} else {
		status = "%" + status + "%"
	}

	query = `
	SELECT 
		b.id, b.user_id, u.name, u.role, d.name, a.id, a.name, a.image, c.name, b.activity, b.request_time, b.return_time, b.status, b.description
	FROM borrowORreturn_requests b
	JOIN users u 
		ON b.user_id = u.id
	JOIN assets a
		ON b.asset_id = a.id
	JOIN categories c
		ON a.category_id = c.id
	JOIN divisions d
		ON d.id = u.division_id
	WHERE b.status LIKE ? AND c.name LIKE ? AND b.request_time LIKE ?
	ORDER BY b.request_time DESC
	LIMIT ? OFFSET ?`
	stmt, err := ar.db.Prepare(query)

	if err != nil {
		log.Println(err)
		return requests, 0, err
	}

	defer stmt.Close()

	res, err := stmt.Query(status, "%"+category+"%", "%"+date+"%", limit, offset-1)

	if err != nil {
		log.Println(err)
		return requests, 0, err
	}

	defer res.Close()

	for res.Next() {
		request := _entity.RequestResponse{}
		if err := res.Scan(&request.Id, &request.User.Id, &request.User.Name, &request.User.Role, &request.User.Division, &request.Asset.Id, &request.Asset.Name, &request.Asset.Image, &request.Asset.CategoryName, &request.Activity, &request.RequestTime, &request.ReturnTime, &request.Status, &request.Description); err != nil {
			log.Println(err)
			return requests, 0, err
		}

		requests = append(requests, request)
	}

	total, err = ar.countRecordBorrow(status)
	if err != nil {
		return requests, total, err
	}

	return requests, total, nil
}

func (ar *AdminRepository) GetAllManager(divLogin, limit, offset int, status, category, date string) (requests []_entity.RequestResponse, total int, err error) {
	query := ""
	if status == "all" {
		status = ""
	}
	status = status + "%Manager"

	if category == "all" {
		category = ""
	}

	query = `
			SELECT 
				b.id, b.user_id, u.name, u.role, d.name,a.id, a.name, a.image, c.name, b.activity, b.request_time, b.return_time, b.status, b.description
			FROM borrowORreturn_requests b
			JOIN users u 
				ON b.user_id = u.id 
			JOIN assets a
				ON b.asset_id = a.id
			JOIN categories c
				ON a.category_id = c.id
			JOIN divisions d
				ON d.id = u.division_id
			WHERE u.division_id = ? AND b.status LIKE ? AND c.name LIKE ? AND b.request_time LIKE ?
			ORDER BY b.request_time DESC
			LIMIT ? OFFSET ?`

	stmt, err := ar.db.Prepare(query)

	if err != nil {
		log.Println(err)
		return requests, 0, err
	}

	defer stmt.Close()

	res, err := stmt.Query(divLogin, status, "%"+category+"%", "%"+date+"%", limit, offset-1)

	if err != nil {
		log.Println(err)
		return requests, 0, err
	}

	defer res.Close()

	for res.Next() {
		request := _entity.RequestResponse{}

		if err := res.Scan(&request.Id, &request.User.Id, &request.User.Name, &request.User.Role, &request.User.Division, &request.Asset.Id, &request.Asset.Name, &request.Asset.Image, &request.Asset.CategoryName, &request.Activity, &request.RequestTime, &request.ReturnTime, &request.Status, &request.Description); err != nil {
			log.Println(err)
			return requests, 0, err
		}

		requests = append(requests, request)
	}
	total, err = ar.countRecordBorrow(status)
	if err != nil {
		return requests, total, err
	}
	return requests, total, nil
}

func (ar *AdminRepository) GetAllProcureManager(limit, offset int, status, category, date string) (requests []_entity.Procure, total int, err error) {
	query := ""
	if status == "all" {
		status = "%Manager"
	} else {
		status = status + "%Manager"
	}
	if category == "all" {
		category = ""
	}
	query = `
			SELECT 
				p.id, p.user_id, u.name, u.role, d.name, c.name, p.image, p.activity, p.request_time, p.status, p.description
			FROM procurement_requests p
			JOIN users u 
				ON p.user_id = u.id 
			JOIN categories c
				ON p.category_id = c.id
			JOIN divisions d
				ON d.id = u. division_id
			WHERE p.status LIKE ? AND c.name LIKE ? AND p.request_time LIKE ?
			ORDER BY p.request_time DESC
			LIMIT ? OFFSET ?`

	stmt, err := ar.db.Prepare(query)

	if err != nil {
		log.Println(err)
		return requests, 0, err
	}

	defer stmt.Close()

	res, err := stmt.Query("%"+status+"%", "%"+category+"%", "%"+date+"%", limit, offset-1)

	if err != nil {
		log.Println(err)
		return requests, 0, err
	}

	defer res.Close()

	for res.Next() {
		request := _entity.Procure{}
		if err := res.Scan(&request.Id, &request.User.Id, &request.User.Name, &request.User.Role, &request.User.Division, &request.Category, &request.Image, &request.Activity, &request.RequestTime, &request.Status, &request.Description); err != nil {
			log.Println(err)
			return requests, 0, err
		}

		requests = append(requests, request)
	}

	total, err = ar.countRecordProcure(status)
	if err != nil {
		return requests, total, err
	}
	return requests, total, nil
}

func (rr *AdminRepository) GetUserDivision(id int) (divId int, code int, err error) {
	stmt, err := rr.db.Prepare(`
		SELECT division_id
		FROM users 
		WHERE deleted_at IS NULL
		  AND id = ? 
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return divId, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return divId, code, err
	}

	defer res.Close()

	if res.Next() {
		if err := res.Scan(&divId); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return divId, code, err
		}
	}

	if divId == 0 {
		log.Println("user id not found")
		code, err = http.StatusBadRequest, errors.New("user not found")
		return divId, code, err
	}

	return divId, http.StatusOK, nil
}

func (ar *AdminRepository) countRecordBorrow(status string) (total int, err error) {
	stmt, err := ar.db.Prepare(`
	SELECT COUNT(status) 
	FROM borrowORreturn_requests
	WHERE status LIKE ? 
	`)

	if err != nil {
		log.Println(err)
		return total, err
	}

	defer stmt.Close()

	res, err := stmt.Query(status)

	if err != nil {
		log.Println(err)
		return total, err
	}

	defer res.Close()

	if res.Next() {
		if err := res.Scan(&total); err != nil {
			log.Println(err)
			return total, err
		}
	}

	return total, nil
}

func (ar *AdminRepository) countRecordProcure(status string) (total int, err error) {
	stmt, err := ar.db.Prepare(`
	SELECT COUNT(status) 
	FROM procurement_requests
	WHERE status LIKE ? 
	`)

	if err != nil {
		log.Println(err)
		return total, err
	}

	defer stmt.Close()

	res, err := stmt.Query(status)

	if err != nil {
		log.Println(err)
		return total, err
	}

	defer res.Close()

	if res.Next() {
		if err := res.Scan(&total); err != nil {
			log.Println(err)
			return total, err
		}
	}

	return total, nil
}
