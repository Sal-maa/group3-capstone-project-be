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
func (ur *AssetRepository) GetAll(page int) (assets []_entity.AssetSimplified, code int, err error) {
	var totalAsset int
	stmt, err := ur.db.Prepare(`
	select distinct a.id, a.code_asset,a.image, a.name,a.short_name,a.status,b.name,a.description,a.quantity 
	from assets a join categories b ON a.category_id = b.id
	group by a.name
	limit ? offset ?

	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return assets, code, err
	}

	limit := 8
	offset := (page - 1) * limit

	res, err := stmt.Query(limit, offset)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return assets, code, err
	}

	defer res.Close()

	for res.Next() {
		stmt2, err := ur.db.Prepare(`select count(e.id)
								from assets e 
								where e.deleted_at is null
								`)

		if err != nil {
			log.Println(err)
			return assets, code, err

		}
		res2, err := stmt2.Query()

		if err != nil {
			log.Println(err)
			return assets, code, err

		}

		defer res2.Close()

		for res2.Next() {
			err := res2.Scan(&totalAsset)

			if err != nil {
				log.Println(err)
				return assets, code, err
			}

		}
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
			}
			if asset.Status == "Borrowed" {
				countBorrow++
			}
		}
		asset.UserCount = userCount
		asset.StockAvailable = countBorrow
		asset.TotalData.TotalPage = totalAsset
		asset.Image = fmt.Sprintf("https://capstone-group3.s3.ap-southeast-1.amazonaws.com/%s", asset.Image)
		assets = append(assets, asset)
	}

	return assets, http.StatusOK, nil
}

func (r *AssetRepository) GetAssetByCategory(category string, page int) (asset _entity.AssetSimplified, code int, err error) {
	var totalAsset int
	stmt, err := r.db.Prepare(`
	select a.id, a.code_asset,a.image, a.name,a.short_name,a.status,b.name,a.description,a.quantity
	from assets a join categories b ON a.category_id = b.id
	where a.deleted_at is null and b.name = ? limit ? offset ?`)

	if err != nil {
		log.Println(err)
		return asset, totalAsset, err
	}

	limit := 5
	offset := (page - 1) * limit

	res, err := stmt.Query(category, limit, offset)

	if err != nil {
		log.Println(err)
		return asset, totalAsset, err
	}

	defer res.Close()
	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return asset, code, err
	}

	defer res.Close()

	if res.Next() {
		if err := res.Scan(&asset.Id, &asset.CodeAsset, &asset.Image, &asset.Name,
			&asset.Short_Name, &asset.Status, &asset.CategoryName,
			&asset.Description, &asset.Quantity); err != nil {

			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return asset, totalAsset, err
		}
	}

	// if asset == (_entity.AssetSimplified{}) {
	// 	log.Println("asset not found")
	// 	code, err = http.StatusBadRequest, errors.New("asset not found")
	// 	return asset, code, err
	// }
	return asset, totalAsset, nil
}

func (ur AssetRepository) Create(assetData _entity.Asset) (createdAsset _entity.AssetSimplified, code int, err error) {
	stmt, err := ur.db.Prepare(`
	INSERT INTO assets (code_asset,image, name, short_name, status, category_id, description, quantity)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return createdAsset, code, err
	}

	defer stmt.Close()
	res, err := stmt.Exec(assetData.CodeAsset, assetData.Image, assetData.Name,
		assetData.Short_Name, assetData.Status, assetData.Category.Id,
		assetData.Description, assetData.Quantity)

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
	createdAsset.CodeAsset = assetData.CodeAsset
	createdAsset.Image = assetData.Image
	createdAsset.Name = assetData.Name
	createdAsset.Short_Name = assetData.Short_Name
	createdAsset.Status = assetData.Status
	createdAsset.CategoryName = assetData.Category.Name
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

func (ur *AssetRepository) Update(assetData _entity.Asset) (updateAsset _entity.AssetSimplified, code int, err error) {
	stmt, err := ur.db.Prepare(`
		UPDATE assets
		SET image = ?, name = ?, short_name = ?, status = ?, description = ?,
		quantity = ?, updated_at = CURRENT_TIMESTAMP
		WHERE deleted_at IS NULL AND id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return updateAsset, code, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(assetData.Image, assetData.Name, assetData.Short_Name, assetData.Status,
		assetData.Description, assetData.Quantity, assetData.Id)

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
	if updateAsset.Status == "Asset Uder Maintenance" {
		updateAsset.Status = "Asset Uder Maintenance"
		updateAsset.Status = assetData.Status
	}
	updateAsset.Id = assetData.Id
	updateAsset.Image = assetData.Image
	updateAsset.Name = assetData.Name
	updateAsset.Description = assetData.Description
	updateAsset.Quantity = assetData.Quantity
	return updateAsset, http.StatusOK, nil
}
