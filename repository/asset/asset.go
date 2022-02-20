package asset

import (
	_entity "capstone/be/entity"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type AssetRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *AssetRepository {
	return &AssetRepository{db: db}
}

func (ur AssetRepository) Create(assetData _entity.Asset) (createdAsset _entity.Asset, code int, err error) {
	stmt, err := ur.db.Prepare(`
	INSERT INTO assets (image, name, entry_date,status,address,description,quantity)
	VALUES (?, ?, TIMESTAMP,?,?,?,?)
`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return createdAsset, code, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(assetData.Image, assetData.Name, assetData.Entry_date, assetData.Status,
		assetData.Address, assetData.Description, assetData.Quantity)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return createdAsset, code, err
	}
	rowsAffected, _ := res.RowsAffected()

	if rowsAffected == 0 {
		log.Println(err)
		code, err = http.StatusInternalServerError, fmt.Errorf("asset not created")
		return createdAsset, code, err
	}
	id, err := res.LastInsertId()

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return createdAsset, code, err
	}

	createdAsset.Id = int(id)
	createdAsset.Image = assetData.Image
	createdAsset.Name = assetData.Name
	createdAsset.Entry_date = assetData.Entry_date
	createdAsset.Status = assetData.Status
	createdAsset.Address = assetData.Address
	createdAsset.Description = assetData.Description
	createdAsset.Quantity = assetData.Quantity

	return createdAsset, http.StatusOK, nil
}
