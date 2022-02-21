package request

import "capstone/be/entity"

type Request interface {
	CheckMaintenance(AssetId int) (asset entity.Asset, err error)
	Borrow(reqData entity.Borrow) (entity.Borrow, error)
	Procure(reqData entity.Procure) (entity.Procure, error)
}
