package admin

import (
	_entity "capstone/be/entity"
)

type Admin interface {
	GetAllNewRequest(limit, offset int, status, date string) (requests []_entity.RequestResponse, err error)
}
