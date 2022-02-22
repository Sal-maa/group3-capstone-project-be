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
func (ur *AssetRepository) GetAll() (assets []_entity.AssetSimplified, code int, err error) {
	stmt, err := ur.db.Prepare(`
	select a.id, a.code_asset,a.image, a.name,a.short_name,a.status,b.name,a.description,a.quantity 
	from assets a join categories b ON a.category_id = b.id
	limit ? offset ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return assets, code, err
	}

	var limit int
	var page int
	limit = 8
	page = page % limit
	offset := (page - 1) * limit

	res, err := stmt.Query(limit, offset)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return assets, code, err
	}

	defer res.Close()

	for res.Next() {
		asset := _entity.AssetSimplified{}
		if err := res.Scan(&asset.Id, &asset.CodeAsset, &asset.Image, &asset.Name,
			&asset.Short_Name, &asset.Status, &asset.CategoryName,
			&asset.Description, &asset.Quantity); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return assets, code, err
		}
		var userCount int
		var countBorrow int

		for i := 0; i < len(asset.Status); i++ {
			if asset.Status == "Available" {
				userCount++
				// UserCount = asset.Quantity - UserCount
			}
			if asset.Status == "Borrowed" {
				countBorrow++
				// countBorrow = asset.Quantity - UserCount
			}
		}
		asset.UserCount = userCount
		asset.StockAvailable = countBorrow
		asset.Image = fmt.Sprintf("https://capstone-group3.s3.ap-southeast-1.amazonaws.com/%s", asset.Image)
		assets = append(assets, asset)
	}

	return assets, http.StatusOK, nil
}

func (ur AssetRepository) Create(assetData _entity.Asset) (createdAsset _entity.Asset, code int, err error) {
	stmt, err := ur.db.Prepare(`
	INSERT INTO assets (image, name,status,category_id,description,quantity)
	VALUES (?, ?, ?, ?, ?, ?)
`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return createdAsset, code, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(assetData.Image, assetData.Name, assetData.Status,
		assetData.CategoryId, assetData.Description, assetData.Quantity)

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
	createdAsset.Status = assetData.Status
	createdAsset.CategoryId = assetData.CategoryId
	createdAsset.Description = assetData.Description
	createdAsset.Quantity = assetData.Quantity

	return createdAsset, http.StatusOK, nil
}

func (ur *AssetRepository) GetById(id int) (asset _entity.Asset, code int, err error) {
	stmt, err := ur.db.Prepare(`
	select a.id, a.image, a.name,a.status,a.category_id,
	a.description,a.quantity 
	from assets a JOIN 
	categories b ON a.category_id = b.id
	where a.deleted_at IS NULL AND a.id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return asset, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(id)
	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return asset, code, err
	}

	defer res.Close()

	if res.Next() {
		if err := res.Scan(&asset.Id, &asset.Image, &asset.Name, &asset.Status,
			&asset.CategoryId, &asset.Description, &asset.Quantity); err != nil {

			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return asset, code, err
		}
	}

	if asset == (_entity.Asset{}) {
		log.Println("asset not found")
		code, err = http.StatusBadRequest, errors.New("asset not found")
		return asset, code, err
	}
	return asset, http.StatusOK, nil
}

func (ur *AssetRepository) Update(assetData _entity.Asset) (updateAsset _entity.Asset, code int, err error) {
	stmt, err := ur.db.Prepare(`
		UPDATE assets
		SET image = ?, name = ?, status = ?, description = ?,
		quantity = ?, updated_at = CURRENT_TIMESTAMP
		WHERE deleted_at IS NULL AND id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updateAsset, code, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(assetData.Image, assetData.Name, assetData.Status, assetData.Description, assetData.Quantity, assetData.Id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updateAsset, code, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updateAsset, code, err
	}

	if rowsAffected == 0 {
		log.Println("rows affected is 0 while asset product")
		code, err = http.StatusBadRequest, errors.New("asset not updated")
		return updateAsset, code, err
	}

	updateAsset.Id = assetData.Id
	updateAsset.Image = assetData.Image
	updateAsset.Name = assetData.Name
	updateAsset.Status = assetData.Status
	updateAsset.Description = assetData.Description
	updateAsset.Quantity = assetData.Quantity
	return updateAsset, http.StatusOK, nil
}
