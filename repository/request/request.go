package request

import (
	_entity "capstone/be/entity"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"
)

type RequestRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *RequestRepository {
	return &RequestRepository{db: db}
}

func (rr *RequestRepository) Borrow(reqData _entity.Borrow) (code int, err error) {
	if code, err = rr.checkUserExistence(reqData.User.Id); err != nil {
		return code, err
	}

	if code, err = rr.checkAssetExistence(reqData.Asset.ShortName); err != nil {
		return code, err
	}

	reqData.Asset.Id, code, err = rr.getAvailableAssetId(reqData.Asset.ShortName)

	if err != nil {
		return code, err
	}

	stmt, err := rr.db.Prepare(`
		INSERT INTO borrowORreturn_requests (updated_at, user_id, asset_id, activity, request_time, return_time, status, description) 
		VALUES (CURRENT_TIMESTAMP, ?, ?, "Borrow", CURRENT_TIMESTAMP, ?, ?, ?)
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(reqData.User.Id, reqData.Asset.Id, reqData.ReturnTime, reqData.Status, reqData.Description)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	if code, err = rr.setAssetBorrowed(reqData.Asset.Id); err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, nil
}

func (rr *RequestRepository) Procure(reqData _entity.Procure) (code int, err error) {
	if code, err = rr.checkAdminExistence(reqData.User.Id); err != nil {
		return code, err
	}

	categoryId, code, err := rr.getCategoryId(reqData.Category)

	if err != nil {
		return code, err
	}

	stmt, err := rr.db.Prepare(`
		INSERT INTO procurement_requests (updated_at, user_id, category_id, activity, image, request_time, status, description)
		VALUES (CURRENT_TIMESTAMP, ?, ?,?, ?, CURRENT_TIMESTAMP, ?, ?)
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(reqData.User.Id, categoryId, reqData.Activity, reqData.Image, reqData.Status, reqData.Description)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	return http.StatusOK, nil
}

func (rr *RequestRepository) checkUserExistence(user_id int) (code int, err error) {
	stmt, err := rr.db.Prepare(`
		SELECT id
		FROM users
		WHERE deleted_at IS NULL
		  AND id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(user_id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	defer res.Close()

	if !res.Next() {
		code, err = http.StatusBadRequest, errors.New("user not found")
		return code, err
	}

	return http.StatusOK, nil
}

func (rr *RequestRepository) checkAssetExistence(short_name string) (code int, err error) {
	stmt, err := rr.db.Prepare(`
		SELECT id 
		FROM assets 
		WHERE deleted_at IS NULL
		  AND short_name = ?
		ORDER BY id ASC LIMIT 1
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(short_name)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	defer res.Close()

	if !res.Next() {
		code, err = http.StatusBadRequest, errors.New("asset not found")
		return code, err
	}

	return http.StatusOK, nil
}

func (rr *RequestRepository) getAvailableAssetId(short_name string) (assetId int, code int, err error) {
	stmt, err := rr.db.Prepare(`
		SELECT id 
		FROM assets 
		WHERE status = "Available"
		  AND deleted_at IS NULL
		  AND short_name = ?
		ORDER BY id ASC LIMIT 1
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return assetId, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(short_name)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return assetId, code, err
	}

	defer res.Close()

	if res.Next() {
		if err = res.Scan(&assetId); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return assetId, code, err
		}
	}

	if assetId == 0 {
		log.Println("asset not available")
		code, err = http.StatusBadRequest, errors.New("asset not available")
		return assetId, code, err
	}

	return assetId, http.StatusOK, nil
}

func (rr *RequestRepository) checkAdminExistence(admin_id int) (code int, err error) {
	stmt, err := rr.db.Prepare(`
		SELECT id
		FROM users
		WHERE deleted_at IS NULL
		  AND role = 'Administrator'
		  AND id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(admin_id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	defer res.Close()

	if !res.Next() {
		code, err = http.StatusBadRequest, errors.New("admin not found")
		return code, err
	}

	return http.StatusOK, nil
}

func (rr *RequestRepository) getCategoryId(category string) (categoryId int, code int, err error) {
	stmt, err := rr.db.Prepare(`
		SELECT id 
		FROM categories 
		WHERE UPPER(name) = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return categoryId, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(strings.ToUpper(category))

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return categoryId, code, err
	}

	defer res.Close()

	if res.Next() {
		if err = res.Scan(&categoryId); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return categoryId, code, err
		}
	}

	if categoryId == 0 {
		log.Println(err)
		code, err = http.StatusBadRequest, errors.New("category not found")
		return categoryId, code, err
	}

	return categoryId, http.StatusOK, nil
}

func (rr *RequestRepository) setAssetBorrowed(assetId int) (code int, err error) {
	stmt, err := rr.db.Prepare(`
		UPDATE assets 
		SET updated_at = CURRENT_TIMESTAMP, status = "Borrowed" 
		WHERE status = "Available"
		  AND deleted_at IS NULL
		  AND id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(assetId)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	return http.StatusOK, err
}

func (rr *RequestRepository) GetBorrowById(id int) (req _entity.Borrow, code int, err error) {
	stmt, err := rr.db.Prepare(`
		SELECT id, user_id, asset_id, activity, request_time, return_time, status, description 
		FROM borrowORreturn_requests
		WHERE deleted_at IS NULL
		  AND id = ? 
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return req, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return req, code, err
	}

	defer res.Close()

	if res.Next() {
		if err := res.Scan(&req.Id, &req.User.Id, &req.Asset.Id, &req.Activity, &req.RequestTime, &req.ReturnTime, &req.Status, &req.Description); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return req, code, err
		}
	}

	if req == (_entity.Borrow{}) {
		log.Println("borrow request not found")
		code, err = http.StatusBadRequest, errors.New("request not found")
		return req, code, err
	}

	return req, http.StatusOK, nil
}

func (rr *RequestRepository) GetUserDivision(id int) (divId int, code int, err error) {
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

	return divId, http.StatusBadRequest, nil
}

func (rr *RequestRepository) UpdateBorrow(reqData _entity.Borrow) (updatedReq _entity.Borrow, code int, err error) {
	stmt, err := rr.db.Prepare(`
		UPDATE borrowORreturn_requests
		SET status = ?, return_time= ?, updated_at = CURRENT_TIMESTAMP
		WHERE deleted_at IS NULL
		  AND id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updatedReq, code, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(reqData.Status, reqData.ReturnTime, reqData.Id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updatedReq, code, err
	}

	if code, err = rr.setAvailable(reqData.Asset.Id); err != nil {
		return updatedReq, code, err
	}

	return reqData, http.StatusOK, nil
}

func (rr *RequestRepository) GetProcureById(id int) (req _entity.Procure, code int, err error) {
	stmt, err := rr.db.Prepare(`
		SELECT p.id, p.user_id, c.name, p.activity, p.image, p.request_time, p.status, p.description 
		FROM procurement_requests p
		JOIN categories c
		ON p.category_id = c.id
		WHERE p.deleted_at IS NULL
		  AND p.id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return req, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return req, code, err
	}

	defer res.Close()

	if res.Next() {
		if err := res.Scan(&req.Id, &req.User.Id, &req.Category, &req.Activity, &req.Image, &req.RequestTime, &req.Status, &req.Description); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return req, code, err
		}
	}

	if req == (_entity.Procure{}) {
		log.Println("procure request not found")
		code, err = http.StatusBadRequest, errors.New("request not found")
		return req, code, err
	}

	return req, http.StatusOK, nil
}

func (rr *RequestRepository) UpdateProcure(reqData _entity.Procure) (updatedReq _entity.Procure, code int, err error) {
	stmt, err := rr.db.Prepare(`
		UPDATE procurement_requests
		SET status = ?, updated_at = CURRENT_TIMESTAMP
		WHERE deleted_at IS NULL
		  AND id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updatedReq, code, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(reqData.Status, reqData.Id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updatedReq, code, err
	}

	return reqData, http.StatusOK, nil
}

func (rr *RequestRepository) setAvailable(assetId int) (code int, err error) {
	stmt, err := rr.db.Prepare(`
		UPDATE assets
		SET status = 'Available', updated_at = CURRENT_TIMESTAMP
		WHERE deleted_at IS NULL
		  AND id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(assetId)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	return http.StatusOK, nil
}

// ===========================

func (rr *RequestRepository) UpdateBorrowByAdmin(reqData _entity.Borrow) (_entity.Borrow, error) {
	statement, err := rr.db.Prepare(`
	UPDATE borrowOrreturn_requests
		SET status = ?, updated_at = CURRENT_TIMESTAMP
		WHERE status = "Approve by Manager" AND deleted_at IS NULL AND id = ?`)
	if err != nil {
		log.Println(err)
		return reqData, err
	}

	defer statement.Close()

	_, err = statement.Exec(reqData.Status, reqData.Id)

	return reqData, err
}

func (rr *RequestRepository) UpdateProcureByAdmin(reqData _entity.Procure) (_entity.Procure, error) {
	statement, err := rr.db.Prepare(`
	UPDATE procurement_requests
		SET status = ?, updated_at = CURRENT_TIMESTAMP
		WHERE status = "Approve by Manager" AND deleted_at IS NULL AND id = ?`)
	if err != nil {
		log.Println(err)
		return reqData, err
	}

	defer statement.Close()

	_, err = statement.Exec(reqData.Status, reqData.Id)

	return reqData, err
}

func (rr *RequestRepository) ReturnAdmin(reqData _entity.Borrow) (updatedReq _entity.Borrow, code int, err error) {
	stmt, err := rr.db.Prepare(`
		UPDATE borrowORreturn_requests
		SET activity = ?, updated_at = CURRENT_TIMESTAMP
		WHERE deleted_at IS NULL
		  AND id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updatedReq, code, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(reqData.Activity, reqData.Id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updatedReq, code, err
	}

	return reqData, http.StatusOK, nil
}
