package admin

import (
	_entity "capstone/be/entity"
)

type Admin interface {
	GetAllAdmin(limit, offset int, status, category, date string) (requests []_entity.RequestResponse, err error)
	GetAllManager(divLogin, limit, offset int, status, category, date string) (requests []_entity.RequestResponse, err error)
}
