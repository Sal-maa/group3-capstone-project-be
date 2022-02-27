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

func (ar *AdminRepository) GetAllAdmin(limit, offset int, status, category, date string) (requests []_entity.RequestResponse, err error) {
	query := ""
	if status == "all" {
		status = ""
	}
	if category == "all" {
		category = ""
	}
	query = `
	SELECT 
		b.id, b.user_id, u.name, a.id, a.name, a.image, c.name, b.activity, b.request_time, b.return_time, b.status, b.description
	FROM borrowORreturn_requests b
	JOIN users u 
		ON b.user_id = u.id
	JOIN assets a
		ON b.asset_id = a.id
	JOIN categories c
		ON a.category_id = c.id
	WHERE b.status LIKE ? AND c.name LIKE ? AND b.request_time LIKE ?
	ORDER BY b.request_time DESC
	LIMIT ? OFFSET ?`
	stmt, err := ar.db.Prepare(query)

	if err != nil {
		log.Println(err)
		return requests, err
	}

	defer stmt.Close()

	res, err := stmt.Query("%"+status+"%", "%"+category+"%", "%"+date+"%", limit, offset-1)

	if err != nil {
		log.Println(err)
		return requests, err
	}

	defer res.Close()

	for res.Next() {
		request := _entity.RequestResponse{}
		if err := res.Scan(&request.Id, &request.User.Id, &request.User.Name, &request.Asset.Id, &request.Asset.Name, &request.Asset.Image, &request.Asset.CategoryName, &request.Activity, &request.RequestTime, &request.ReturnTime, &request.Status, &request.Description); err != nil {
			log.Println(err)
			return requests, err
		}

		requests = append(requests, request)
	}

	return requests, nil
}

func (ar *AdminRepository) GetAllManager(divLogin, limit, offset int, status, category, date string) (requests []_entity.RequestResponse, err error) {
	query := ""
	if status == "all" {
		status = ""
	}
	if category == "all" {
		category = ""
	}
	query = `
			SELECT 
				b.id, b.user_id, u.name, a.id, a.name, a.image, c.name, b.activity, b.request_time, b.return_time, b.status, b.description
			FROM borrowORreturn_requests b
			JOIN users u 
				ON b.user_id = u.id 
			JOIN assets a
				ON b.asset_id = a.id
			JOIN categories c
				ON a.category_id = c.id
			WHERE u.division_id = ? AND b.status LIKE ? AND c.name LIKE ? AND b.request_time LIKE ?
			ORDER BY b.request_time DESC
			LIMIT ? OFFSET ?`

	stmt, err := ar.db.Prepare(query)

	if err != nil {
		log.Println(err)
		return requests, err
	}

	defer stmt.Close()

	res, err := stmt.Query(divLogin, "%"+status+"%", "%"+category+"%", "%"+date+"%", limit, offset-1)

	if err != nil {
		log.Println(err)
		return requests, err
	}

	defer res.Close()

	for res.Next() {
		request := _entity.RequestResponse{}

		if err := res.Scan(&request.Id, &request.User.Id, &request.User.Name, &request.Asset.Id, &request.Asset.Name, &request.Asset.Image, &request.Asset.CategoryName, &request.Activity, &request.RequestTime, &request.ReturnTime, &request.Status, &request.Description); err != nil {
			log.Println(err)
			return requests, err
		}

		requests = append(requests, request)
	}

	return requests, nil
}
