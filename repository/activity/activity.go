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
	return activity, http.StatusOK, nil
}
