package category

import (
	_entity "capstone/be/entity"
)

type Category interface {
	GetAll() (category []_entity.Category, code int, err error)
}
