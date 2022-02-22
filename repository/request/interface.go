package request

import (
	_entity "capstone/be/entity"
)

type Request interface {
	CheckMaintenance(reqData _entity.Borrow) (asset _entity.Asset, err error)
	Borrow(reqData _entity.Borrow) (_entity.Borrow, error)
	GetCategoryId(newReq _entity.CreateProcure) (id int, err error)
	Procure(reqData _entity.Procure) (_entity.Procure, error)
	UpdateBorrow(reqData _entity.Borrow) (_entity.Borrow, error)
}
