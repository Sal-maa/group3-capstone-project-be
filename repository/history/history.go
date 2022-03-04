package history

import (
	_entity "capstone/be/entity"
	"fmt"

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

func (hr *HistoryRepository) GetAllRequestHistoryOfUser(user_id int, page int) (count int, histories []_entity.UserRequestHistorySimplified, code int, err error) {
	if code, err = hr.checkUserExistence(user_id); err != nil {
		return count, histories, code, err
	}

	stmt, err := hr.db.Prepare(`
		SELECT b.id, b.request_time, c.name, a.name, a.image
		FROM borrowORreturn_requests b
		JOIN assets a
		ON b.asset_id = a.id
		JOIN categories c
		ON a.category_id = c.id
		WHERE b.deleted_at IS NULL
		  AND b.status = 'Approved by Admin'
		  AND b.activity = 'Return'
		  AND b.user_id = ?
		LIMIT 5 OFFSET ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return count, histories, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(user_id, (page-1)*5)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return count, histories, code, err
	}

	defer res.Close()

	for res.Next() {
		history := _entity.UserRequestHistorySimplified{}

		if err := res.Scan(&history.Id, &history.RequestDate, &history.Category, &history.AssetName, &history.AssetImage); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return count, histories, code, err
		}

		history.ActivityType = "Borrowing Asset"
		history.AssetImage = fmt.Sprintf("https://capstone-group3.s3.ap-southeast-1.amazonaws.com/%s", history.AssetImage)

		histories = append(histories, history)
	}

	count, code, err = hr.getHistoryCount(user_id)

	if err != nil {
		return count, histories, code, err
	}

	return count, histories, http.StatusOK, nil
}

func (hr *HistoryRepository) GetDetailRequestHistoryByRequestId(request_id int) (history _entity.UserRequestHistory, code int, err error) {
	stmt, err := hr.db.Prepare(`
		SELECT c.name, a.name, a.image, u.name, b.request_time, b.return_time, b.description, b.short_name
		FROM borrowORreturn_requests b
		JOIN assets a
		ON b.asset_id = a.id
		JOIN categories c
		ON a.category_id = c.id
		JOIN users u
		ON b.user_id = u.id
		WHERE b.deleted_at IS NULL
		  AND b.status = 'Approved by Admin'
		  AND b.activity = 'Return'
		  AND b.id = ?
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

	short_name := ""

	if res.Next() {
		if err := res.Scan(&history.Category, &history.AssetName, &history.AssetImage, &history.UserName, &history.RequestDate, &history.ReturnDate, &history.Description, &short_name); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return history, code, err
		}
	}

	if history == (_entity.UserRequestHistory{}) {
		log.Println("request not found")
		code, err = http.StatusBadRequest, errors.New("request not found")
		return history, code, err
	}

	history.Id = request_id
	history.AssetImage = fmt.Sprintf("https://capstone-group3.s3.ap-southeast-1.amazonaws.com/%s", history.AssetImage)
	history.Status = "Successfully returned"

	stock, err := hr.getAssetStock(short_name)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return history, code, err
	}

	history.StockLeft = stock

	return history, http.StatusOK, nil
}

func (hr *HistoryRepository) GetAllUsageHistoryOfAsset(short_name string) (asset _entity.AssetInfo, users []_entity.AssetUser, code int, err error) {
	asset, code, err = hr.getAssetDetail(short_name)

	if err != nil {
		return asset, users, code, err
	}

	if asset == (_entity.AssetInfo{}) {
		code, err = http.StatusBadRequest, errors.New("request not found")
		return asset, users, code, err
	}

	stmt, err := hr.db.Prepare(`
		SELECT u.name, b.request_time, b.activity
		FROM borrowORreturn_requests b
		JOIN users u
		ON b.user_id = u.id
		WHERE b.deleted_at IS NULL
		  AND b.status = 'Approved by Admin'
		  AND b.short_name = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return asset, users, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(short_name)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return asset, users, code, err
	}

	defer res.Close()

	for res.Next() {
		user := _entity.AssetUser{}

		if err := res.Scan(&user.Name, &user.RequestDate, &user.Status); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return asset, users, code, err
		}

		user.Status += "ed"

		users = append(users, user)
	}

	return asset, users, http.StatusOK, nil
}

func (hr *HistoryRepository) checkUserExistence(user_id int) (code int, err error) {
	stmt, err := hr.db.Prepare(`
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

func (hr *HistoryRepository) getHistoryCount(user_id int) (count int, code int, err error) {
	stmt, err := hr.db.Prepare(`
		SELECT COUNT(id)
		FROM borrowORreturn_requests
		WHERE deleted_at IS NULL
		  AND status = 'Approved by Admin'
		  AND activity = 'Return'
		  AND user_id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return count, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(user_id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return count, code, err
	}

	defer res.Close()

	if res.Next() {
		if err := res.Scan(&count); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return count, code, err
		}
	}

	return count, http.StatusOK, nil
}

func (hr *HistoryRepository) getAssetDetail(short_name string) (asset _entity.AssetInfo, code int, err error) {
	stmt, err := hr.db.Prepare(`
		SELECT c.name, a.name, a.image
		FROM assets a
		JOIN categories c
		ON a.category_id = c.id
		WHERE deleted_at IS NULL
		  AND short_name = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return asset, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(short_name)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return asset, code, err
	}

	defer res.Close()

	if res.Next() {
		if err = res.Scan(&asset.Category, &asset.AssetName, &asset.AssetImage); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return asset, code, err
		}
	}

	asset.AssetImage = fmt.Sprintf("https://capstone-group3.s3.ap-southeast-1.amazonaws.com/%s", asset.AssetImage)

	return asset, http.StatusOK, nil
}

func (hr *HistoryRepository) getAssetStock(short_name string) (stock int, err error) {
	stmt, err := hr.db.Prepare(`
		SELECT COUNT(id)
		FROM assets
		WHERE deleted_at IS NULL
		  AND status = 'Available'
		  AND short_name = ?
		GROUP BY short_name
	`)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Query(short_name)

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
