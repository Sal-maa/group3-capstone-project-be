package history

import (
	_entity "capstone/be/entity"

	"database/sql"
	"errors"
	"log"
	"net/http"
)

type HistoryRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *HistoryRepository {
	return &HistoryRepository{db: db}
}

func (hr *HistoryRepository) GetAllUsageHistoryOfUser(user_id int, page int) (histories []_entity.UserUsageHistorySimplified, code int, err error) {
	if code, err = hr.checkUserExistence(user_id); err != nil {
		return histories, code, err
	}

	stmt, err := hr.db.Prepare(`
		SELECT b.id, b.request_time, c.name, a.name, a.image
		FROM borrow_return_requests b
		JOIN assets a
		ON b.asset_id = a.id
		JOIN categories c
		ON a.category_id = c.id
		WHERE b.deleted_at IS NULL AND b.status = 'Approved by Admin' AND b.activity = 'Return' AND b.user_id = ?
		LIMIT 5 OFFSET ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return histories, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(user_id, (page-1)*5)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return histories, code, err
	}

	defer res.Close()

	for res.Next() {
		history := _entity.UserUsageHistorySimplified{}

		if err := res.Scan(&history.Id, &history.RequestDate, &history.Category, &history.AssetName, &history.AssetImage); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return histories, code, err
		}

		history.ActivityType = "Borrowing Asset"

		histories = append(histories, history)
	}

	return histories, http.StatusOK, nil
}

func (hr *HistoryRepository) GetDetailUsageHistoryByRequestId(request_id int) (history _entity.UserUsageHistory, code int, err error) {
	stmt, err := hr.db.Prepare(`
		SELECT c.name, a.name, a.image, u.name, b.request_time, b.return_time, b.description, b.asset_id
		FROM borrow_return_requests b
		JOIN assets a
		ON b.asset_id = a.id
		JOIN categories c
		ON a.category_id = c.id
		JOIN users u
		ON b.user_id = u.id
		WHERE b.deleted_at IS NULL AND b.status = 'Approved by Admin' AND b.activity = 'Return' AND b.id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return history, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(request_id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return history, code, err
	}

	defer res.Close()

	asset_id := 0

	if res.Next() {
		if err := res.Scan(&history.Category, &history.AssetName, &history.AssetImage, &history.UserName, &history.RequestDate, &history.ReturnDate, &history.Description, &asset_id); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return history, code, err
		}
	}

	if history == (_entity.UserUsageHistory{}) {
		log.Println("request not found")
		code, err = http.StatusBadRequest, errors.New("request not found")
		return history, code, err
	}

	history.Id = request_id
	history.Status = "Successfully returned"

	stock, err := hr.getAssetStock(asset_id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return history, code, err
	}

	history.StockLeft = stock

	return history, http.StatusOK, nil
}

func (hr *HistoryRepository) GetAllUsageHistoryOfAsset(asset_id int) (histories []_entity.AssetUsageHistory, code int, err error) {
	if code, err = hr.checkAssetExistence(asset_id); err != nil {
		return histories, code, err
	}

	stmt, err := hr.db.Prepare(`
		SELECT b.id, c.name, a.name, a.image, u.name, b.request_time, b.activity
		FROM borrow_return_requests b
		JOIN assets a
		ON b.asset_id = a.id
		JOIN categories c
		ON a.category_id = c.id
		JOIN users u
		ON b.user_id = u.id
		WHERE b.deleted_at IS NULL AND b.status = 'Approved by Admin' AND b.asset_id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return histories, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(asset_id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return histories, code, err
	}

	defer res.Close()

	for res.Next() {
		history := _entity.AssetUsageHistory{}

		if err := res.Scan(&history.Id, &history.Category, &history.AssetName, &history.AssetImage, &history.UserName, &history.RequestDate, &history.Status); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return histories, code, err
		}

		history.Status += "ed"

		histories = append(histories, history)
	}

	return histories, http.StatusOK, nil
}

func (hr *HistoryRepository) checkUserExistence(user_id int) (code int, err error) {
	stmt, err := hr.db.Prepare(`
		SELECT id
		FROM users
		WHERE deleted_at IS NULL AND id = ?
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
		code, err = http.StatusBadRequest, errors.New("user not exist")
		return code, err
	}

	return http.StatusOK, nil
}

func (hr *HistoryRepository) checkAssetExistence(asset_id int) (code int, err error) {
	stmt, err := hr.db.Prepare(`
		SELECT id
		FROM assets
		WHERE deleted_at IS NULL AND id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(asset_id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return code, err
	}

	defer res.Close()

	if !res.Next() {
		code, err = http.StatusBadRequest, errors.New("asset not exist")
		return code, err
	}

	return http.StatusOK, nil
}

func (hr *HistoryRepository) getAssetStock(asset_id int) (stock int, err error) {
	stmt, err := hr.db.Prepare(`
		SELECT status
		FROM assets
		WHERE id = ?
	`)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Query(asset_id)

	if err != nil {
		return 0, err
	}

	defer res.Close()

	status := ""

	if res.Next() {
		if err = res.Scan(&status); err != nil {
			return 0, err
		}
	}

	if status == "Asset under maintenance" {
		return 0, nil
	}

	stmt, err = hr.db.Prepare(`
		SELECT COUNT(id)
		FROM borrow_return_requests
		WHERE deleted_at IS NULL AND status = 'Waiting approval from Admin' OR status = 'Waiting approval from Manager' OR status = 'Approved by Manager' OR status = 'Approved by Admin' AND b.asset_id = ?
	`)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err = stmt.Query(asset_id)

	if err != nil {
		return 0, err
	}

	defer res.Close()

	if res.Next() {
		if err = res.Scan(&stock); err != nil {
			return 0, err
		}
	}

	return stock, nil
}
