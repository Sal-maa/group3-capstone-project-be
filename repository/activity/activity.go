package activity

import (
	_entity "capstone/be/entity"
	"errors"
	"fmt"
	"log"

	"database/sql"
	"net/http"
)

type ActivityRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *ActivityRepository {
	return &ActivityRepository{db: db}
}

func (ar ActivityRepository) GetAllRequestOfUser(user_id int) (activities []_entity.ActivitySimplified, code int, err error) {
	stmt, err := ar.db.Prepare(`
		SELECT b.id, a.name, a.image, b.status, b.request_time
		FROM borrowORreturn_requests b
		JOIN assets a
		ON b.asset_id = a.id
		WHERE b.deleted_at IS NULL
		  AND b.activity <> "Return"
		  AND b.status <> "Approved by Admin"
		  AND b.user_id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return activities, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query()

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return activities, code, err
	}

	defer res.Close()

	for res.Next() {
		activity := _entity.ActivitySimplified{}

		if err := res.Scan(&activity.Id, &activity.AssetName, &activity.AssetImage, &activity.Status, &activity.RequestDate); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return activities, code, err
		}

		activity.AssetImage = fmt.Sprintf("https://capstone-group3.s3.ap-southeast-1.amazonaws.com/%s", activity.AssetImage)

		activities = append(activities, activity)
	}

	return activities, http.StatusOK, nil
}

func (ar ActivityRepository) GetByRequestId(request_id int) (activity _entity.Activity, code int, err error) {
	stmt, err := ar.db.Prepare(`
		SELECT c.name, a.name, a.image, u.name, b.request_time, b.return_time, b.description, b.short_name
		FROM borrowORreturn_requests b
		JOIN assets a
		ON b.asset_id = a.id
		JOIN categories c
		ON a.category_id = c.id
		JOIN users u
		ON b.user_id = u.id
		WHERE b.deleted_at IS NULL
		  AND b.id = ?
	`)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return activity, code, err
	}

	defer stmt.Close()

	res, err := stmt.Query(request_id)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return activity, code, err
	}

	defer res.Close()

	short_name := ""

	if res.Next() {
		if err := res.Scan(&activity.Category, &activity.AssetName, &activity.AssetImage, &activity.UserName, &activity.RequestDate, &activity.ReturnDate, &activity.Description, short_name); err != nil {
			log.Println(err)
			code, err = http.StatusInternalServerError, errors.New("internal server error")
			return activity, code, err
		}
	}

	if activity == (_entity.Activity{}) {
		log.Println("request not found")
		code, err = http.StatusBadRequest, errors.New("request not found")
		return activity, code, err
	}

	activity.Id = request_id
	activity.AssetImage = fmt.Sprintf("https://capstone-group3.s3.ap-southeast-1.amazonaws.com/%s", activity.AssetImage)
	activity.Status = "Successfully returned"

	stock, err := ar.getAssetStock(short_name)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return activity, code, err
	}

	activity.StockLeft = stock

	return activity, http.StatusOK, nil
}

func (ar *ActivityRepository) getAssetStock(short_name string) (stock int, err error) {
	stmt, err := ar.db.Prepare(`
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
