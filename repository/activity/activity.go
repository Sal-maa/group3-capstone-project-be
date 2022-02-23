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

func (ar ActivityRepository) GetAllActivityOfUser(user_id int) (activities []_entity.ActivitySimplified, code int, err error) {
	if code, err = ar.checkUserExistence(user_id); err != nil {
		return activities, code, err
	}

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

func (ar ActivityRepository) GetDetailActivityByRequestId(request_id int) (activity _entity.Activity, code int, err error) {
	stmt, err := ar.db.Prepare(`
		SELECT c.name, a.name, a.image, u.name, b.request_time, b.return_time, b.description, b.activity, b.status, b.short_name
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
		if err := res.Scan(&activity.Category, &activity.AssetName, &activity.AssetImage, &activity.UserName, &activity.RequestDate, &activity.ReturnDate, &activity.Description, &activity.ActivityType, &activity.Status, short_name); err != nil {
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

	stock, err := ar.getAssetStock(short_name)

	if err != nil {
		log.Println(err)
		code, err = http.StatusInternalServerError, errors.New("internal server error")
		return activity, code, err
	}

	activity.StockLeft = stock

	return activity, http.StatusOK, nil
}

func (ar ActivityRepository) CancelRequest(request_id int) (code int, err error) {
	stmt, err := ar.db.Prepare(`
		UPDATE borrowORreturn_requests
		SET status = 'Cancelled'
		WHERE deleted_at IS NULL AND id = ?
	`)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	defer stmt.Close()

	res, err := stmt.Exec(request_id)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	if rowsAffected == 0 {
		log.Println("rows affected is 0 while cancelling request")
		return http.StatusBadRequest, errors.New("request not cancelled")
	}

	stmt, err = ar.db.Prepare(`
		UPDATE borrowORreturn_requests
		SET updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	defer stmt.Close()

	_, err = stmt.Exec(request_id)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	return http.StatusOK, nil
}

func (ar ActivityRepository) ReturnRequest(request_id int) (code int, err error) {
	stmt, err := ar.db.Prepare(`
		UPDATE borrowORreturn_requests
		SET activity = 'Return', status = 'Waiting approval', return_time = CURRENT_TIMESTAMP
		WHERE deleted_at IS NULL AND id = ?
	`)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	defer stmt.Close()

	res, err := stmt.Exec(request_id)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	if rowsAffected == 0 {
		log.Println("rows affected is 0 while cancelling request")
		return http.StatusBadRequest, errors.New("request not cancelled")
	}

	stmt, err = ar.db.Prepare(`
		UPDATE borrowORreturn_requests
		SET updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	defer stmt.Close()

	_, err = stmt.Exec(request_id)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	return http.StatusOK, nil
}

func (ar *ActivityRepository) checkUserExistence(user_id int) (code int, err error) {
	stmt, err := ar.db.Prepare(`
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
