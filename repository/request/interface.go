package request

import "capstone/be/entity"

type Request interface {
	Borrow(reqData entity.Borrow) (entity.Borrow, error)
	Procure(reqData entity.Procure) (entity.Procure, error)
}
