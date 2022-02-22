package request

import (
	"capstone/be/entity"
	"database/sql"
)

type RequestRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *RequestRepository {
	return &RequestRepository{db: db}
}

func (rr *RequestRepository) CheckMaintenance(AssetId int) (asset entity.Asset, err error) {
	stmt, err := rr.db.Prepare(`
		SELECT status FROM assets WHERE deleted_at IS NULL AND id = ?
	`)

	if err != nil {
		return asset, err
	}

	defer stmt.Close()

	res, err := stmt.Query(AssetId)

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

func (rr *RequestRepository) Borrow(reqData entity.Borrow) (entity.Borrow, error) {
	_, err := rr.db.Exec("INSERT INTO `borrow/return_requests` (updated_at, user_id, asset_id, activity, request_time, return_time, status, description) VALUES(?,?,?,?,?,?,?,?)", reqData.UpdatedAt, reqData.User.Id, reqData.Asset.Id, reqData.Activity, reqData.RequestTime, reqData.ReturnTime, reqData.Status, reqData.Description)
	return reqData, err
}

func (rr *RequestRepository) Procure(reqData entity.Procure) (entity.Procure, error) {
	_, err := rr.db.Exec("INSERT INTO `procurement_requests` (updated_at, user_id, category_id, image, activity, request_time, status, description) VALUES(?,?,?,?,?,?,?,?)", reqData.UpdatedAt, reqData.CategoryId, reqData.Image, reqData.Activity, reqData.RequestTime, reqData.Status, reqData.Description)
	return reqData, err
}
