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

func (ar AssetRepository) Create(assetData _entity.Asset) (code int, err error) {
	stmt, err := ar.db.Prepare(`
		INSERT INTO assets (code, name, short_name, category_id, description, status, image, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(assetData.Code, assetData.Name, assetData.ShortName, assetData.CategoryId, assetData.Description, assetData.Status, assetData.Image)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	return http.StatusOK, nil
}

func (ar *AssetRepository) GetAll() (assets []_entity.AssetSimplified, code int, err error) {
	stmt, err := ar.db.Prepare(`
		SELECT c.name, a.name, a.short_name, a.image, a.description
		FROM assets a
		JOIN categories c
		ON a.category_id = c.id
		WHERE a.deleted_at IS NULL
		GROUP BY a.short_name
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return assets, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query()

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return assets, code, err
	}

	defer res.Close()

	for res.Next() {
		asset := _entity.AssetSimplified{}

		if err := res.Scan(&asset.Category, &asset.Name, &asset.ShortName, &asset.Image, &asset.Description); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return assets, code, err
		}

		users, available, err := ar.getStatsByShortName(asset.ShortName)

		if err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return assets, code, err
		}

		asset.UserCount = users
		asset.StockAvailable = available
		asset.Image = fmt.Sprintf("https://capstone-group3.s3.ap-southeast-1.amazonaws.com/%s", asset.Image)

		assets = append(assets, asset)
	}

	return assets, http.StatusOK, nil
}

func (ar *AssetRepository) GetAssetsByCategory(category_id int) (assets []_entity.AssetSimplified, code int, err error) {
	stmt, err := ar.db.Prepare(`
		SELECT c.name, a.name, a.short_name, a.image, a.description
		FROM assets a
		JOIN categories c
		ON a.category_id = c.id
		WHERE a.deleted_at IS NULL
		  AND a.category_id = ?
		GROUP BY a.short_name
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return assets, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(category_id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return assets, code, err
	}

	defer res.Close()

	for res.Next() {
		asset := _entity.AssetSimplified{}

		if err := res.Scan(&asset.Category, &asset.Name, &asset.ShortName, &asset.Image, &asset.Description); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return assets, code, err
		}

		users, available, err := ar.getStatsByShortName(asset.ShortName)

		if err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return assets, code, err
		}

		asset.UserCount = users
		asset.StockAvailable = available
		asset.Image = fmt.Sprintf("https://capstone-group3.s3.ap-southeast-1.amazonaws.com/%s", asset.Image)

		assets = append(assets, asset)
	}

	return assets, http.StatusOK, nil
}

func (ar *AssetRepository) GetByShortName(short_name string) (total int, asset _entity.AssetSimplified, code int, err error) {
	stmt, err := ar.db.Prepare(`
		SELECT c.name, a.name, a.image, a.description, COUNT(a.id)
		FROM assets a
		JOIN categories c
		ON a.category_id = c.id
		WHERE deleted_at IS NULL
		  AND a.short_name = ?
		GROUP BY a.short_name
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return total, asset, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(short_name)
	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return total, asset, code, err
	}

	defer res.Close()

	if res.Next() {
		if err := res.Scan(&asset.Category, &asset.Name, &asset.Image, &asset.Description, &total); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return total, asset, code, err
		}
	}

	if asset == (_entity.AssetSimplified{}) {
		log.Println("asset not found")
		code, err = http.StatusBadRequest, errors.New("asset not found")
		return total, asset, code, err
	}

	asset.Image = fmt.Sprintf("https://capstone-group3.s3.ap-southeast-1.amazonaws.com/%s", asset.Image)

	return total, asset, http.StatusOK, nil
}

func (ar *AssetRepository) SetMaintenance(short_name string) (code int, err error) {
	stmt, err := ar.db.Prepare(`
		UPDATE assets
		SET status = 'Asset Under Maintenance', updated_at = CURRENT_TIMESTAMP
		WHERE deleted_at IS NULL
		  AND status = 'Available'
		  AND short_name = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(short_name)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	if rowsAffected == 0 {
		log.Println("rows affected is 0 while update asset")
		code, err = http.StatusBadRequest, errors.New("asset not updated")
		return code, err
	}

	return http.StatusOK, nil
}

func (ar *AssetRepository) SetAvailable(short_name string) (code int, err error) {
	stmt, err := ar.db.Prepare(`
		UPDATE assets
		SET status = 'Available', updated_at = CURRENT_TIMESTAMP
		WHERE deleted_at IS NULL
		  AND status = 'Asset Under Maintenance'
		  AND short_name = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(short_name)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	if rowsAffected == 0 {
		log.Println("rows affected is 0 while update asset")
		code, err = http.StatusBadRequest, errors.New("asset not updated")
		return code, err
	}

	return http.StatusOK, nil
}

func (ar *AssetRepository) GetStats() (statistics _entity.Statistics, code int, err error) {
	stmt, err := ar.db.Prepare(`
		SELECT status, count(status)
		FROM assets
		GROUP BY status
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return statistics, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query()

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return statistics, code, err
	}

	defer res.Close()

	for res.Next() {
		status, count := "", 0

		if err := res.Scan(&status, &count); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return statistics, code, err
		}

		switch status {
		case "Asset Under Maintenance":
			statistics.UnderMaintenance = count
		case "Available":
			statistics.Available = count
		case "Borrowed":
			statistics.Borrowed = count
		default:
			log.Println("illegal status")
		}
	}

	statistics.TotalAsset = statistics.Borrowed + statistics.UnderMaintenance + statistics.Available

	return statistics, http.StatusOK, nil
}

func (ar *AssetRepository) GetCategoryId(category string) (id int, code int, err error) {
	stmt, err := ar.db.Prepare(`
		SELECT id
		FROM categories
		WHERE name = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return id, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(category)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return id, code, err
	}

	defer res.Close()

	if res.Next() {
		if err = res.Scan(&id); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return id, code, err
		}
	}

	if id == 0 {
		log.Println("illegal category")
		code, err = http.StatusBadRequest, errors.New("category not exist")
		return id, code, err
	}

	return id, http.StatusOK, nil
}

func (ar *AssetRepository) getStatsByShortName(short_name string) (users int, available int, err error) {
	stmt, err := ar.db.Prepare(`
		SELECT status, COUNT(status)
		FROM assets
		WHERE deleted_at IS NULL
		  AND status <> 'Asset Under Maintenance'
		  AND short_name = ?
		GROUP BY status
	`)

	if err != nil {
		return users, available, err
	}

	defer stmt.Close()

	res, err := stmt.Query(short_name)

	if err != nil {
		return users, available, err
	}

	defer res.Close()

	for res.Next() {
		status, count := "", 0

		if err := res.Scan(&status, &count); err != nil {
			return users, available, err
		}

		if status == "Available" {
			available = count
		} else if status == "Borrowed" {
			users = count
		}
	}

	return users, available, nil
}
