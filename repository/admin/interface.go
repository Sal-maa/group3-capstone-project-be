package admin

import (
	_entity "capstone/be/entity"
)

type Admin interface {
	GetAllNewRequest(limit, offset int) (requests []_entity.RequestResponse, err error)
}
