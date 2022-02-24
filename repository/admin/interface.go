package admin

import (
	_entity "capstone/be/entity"
)

type Admin interface {
	GetAllNewRequest() (requests []_entity.RequestResponse, err error)
}
