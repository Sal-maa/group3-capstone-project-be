package admin

import (
	_entity "capstone/be/entity"
	"database/sql"
	"log"
)

type AdminRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (ar *AdminRepository) GetAllNewRequest(limit, offset int) (requests []_entity.RequestResponse, err error) {
	stmt, err := ar.db.Prepare(`
	SELECT 
		b.id, b.user_id, u.name, a.id, a.name, a.image, c.name, b.activity, b.request_time, b.return_time, b.status, b.description
	FROM borrowORreturn_requests b
	JOIN users u 
		ON b.user_id = u.id
	JOIN assets a
		ON b.asset_id = a.id
	JOIN category c
		ON a.category_id = c.id
	WHERE b.request_time = DATE(NOW())
	ORDER BY b.request_time DESC
	LIMIT ? OFFSET ?
	`)

	if err != nil {
		log.Println(err)
		return requests, err
	}

	defer stmt.Close()

	res, err := stmt.Query(limit, offset)

	if err != nil {
		log.Println(err)
		return requests, err
	}

	defer res.Close()

	for res.Next() {
		request := _entity.RequestResponse{}

		if err := res.Scan(&request.Id, &request.User.Id, &request.User.Name, &request.Asset.Id, &request.Asset.Image, &request.Asset.CategoryId, &request.Activity, &request.RequestTime, &request.ReturnTime, &request.Status, &request.Description); err != nil {
			log.Println(err)
			return requests, err
		}

		requests = append(requests, request)
	}

	return requests, nil
}
