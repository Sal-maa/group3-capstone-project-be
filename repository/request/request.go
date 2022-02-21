package request

import (
	"capstone/be/entity"
	_entity "capstone/be/entity"
	"database/sql"
)

type RequestRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *RequestRepository {
	return &RequestRepository{db: db}
}

func (rr *RequestRepository) CheckMaintenance(reqData entity.Borrow) (asset entity.Asset, err error) {
	stmt, err := rr.db.Prepare(`
		SELECT status FROM assets WHERE deleted_at IS NULL AND id = ?
	`)

	if err != nil {
		return asset, err
	}

	defer stmt.Close()

	res, err := stmt.Query(reqData.Asset.Id)

	if err != nil {
		return asset, err
	}

	defer res.Close()

	if res.Next() {
		if err := res.Scan(&asset.Status); err != nil {
			return asset, err
		}
	}
	return asset, nil
}

func (rr *RequestRepository) Borrow(reqData _entity.Borrow) (_entity.Borrow, error) {
	statement, err := rr.db.Prepare("INSERT INTO `borrow/return_requests` (updated_at, user_id, asset_id, activity, request_time, return_time, status, description) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return reqData, err
	}

	defer statement.Close()

	_, err = statement.Exec(reqData.UpdatedAt, reqData.User.Id, reqData.Asset.Id, reqData.Activity, reqData.RequestTime, reqData.ReturnTime, reqData.Status, reqData.Description)

	return reqData, err
}

func (rr *RequestRepository) GetCategoryId(newReq entity.CreateProcure) (id int, err error) {
	stmt, err := rr.db.Prepare(`
		SELECT id
		FROM categories
		WHERE name = ?
	`)

	if err != nil {
		return id, err
	}

	defer stmt.Close()

	res, err := stmt.Query(newReq.Category)

	if err != nil {
		return id, err
	}

	defer res.Close()

	if res.Next() {
		if err = res.Scan(&id); err != nil {
			return id, err
		}
	}

	return id, nil
}

func (rr *RequestRepository) Procure(reqData entity.Procure) (entity.Procure, error) {
	statement, err := rr.db.Prepare("INSERT INTO `procurement_requests` (updated_at, user_id, category_id, image, activity, request_time, status, description) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return reqData, err
	}

	defer statement.Close()

	_, err = statement.Exec(reqData.UpdatedAt, reqData.Category, reqData.Image, reqData.Activity, reqData.RequestTime, reqData.Status, reqData.Description)
	return reqData, err
}
