package request

import "capstone/be/entity"

type Request interface {
	CheckMaintenance(reqData entity.Borrow) (asset entity.Asset, err error)
	Borrow(reqData entity.Borrow) (entity.Borrow, error)
	GetCategoryId(newReq entity.CreateProcure) (id int, err error)
	Procure(reqData entity.Procure) (entity.Procure, error)
}
