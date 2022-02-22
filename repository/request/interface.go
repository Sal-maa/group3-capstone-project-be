package request

import (
	_entity "capstone/be/entity"
)

type Request interface {
	GetAssetId(newReq _entity.CreateBorrow) (id int, err error)
	CheckMaintenance(reqData _entity.Borrow) (asset _entity.Asset, err error)
	Borrow(reqData _entity.Borrow) (_entity.Borrow, error)
	GetCategoryId(newReq _entity.CreateProcure) (id int, err error)
	Procure(reqData _entity.Procure) (_entity.Procure, error)
	GetUserDivision(id int) (divId int, err error)
	GetBorrowById(id int) (req _entity.Borrow, err error)
	UpdateBorrow(reqData _entity.Borrow) (_entity.Borrow, error)
	UpdateProcure(reqData _entity.Procure) (_entity.Procure, error)
	UpdateBorrowByAdmin(reqData _entity.Borrow) (_entity.Borrow, error)
	UpdateProcureByAdmin(reqData _entity.Procure) (_entity.Procure, error)
}
