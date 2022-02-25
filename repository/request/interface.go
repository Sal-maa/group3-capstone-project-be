package request

import (
	_entity "capstone/be/entity"
)

type Request interface {
	GetAssetId(newReq _entity.CreateBorrow) (id int, err error)
	CheckMaintenance(assetId int) (statAsset string, err error)
	Borrow(reqData _entity.Borrow) (_entity.Borrow, error)
	UpdateAssetStatus(assetId int) (assetUpdate string, err error)
	GetCategoryId(category string) (id int, err error)
	GetCategoryIdAsset(assetId int) (id int, err error)
	GetEmployeeId(name string) (id int, err error)
	AddCategory(category string) (string, error)
	Procure(reqData _entity.Procure) (_entity.Procure, error)
	GetUserDivision(id int) (divId int, err error)
	GetBorrowById(id int) (req _entity.Borrow, err error)
	GetProcureById(id int) (req _entity.Procure, err error)
	UpdateBorrow(reqData _entity.Borrow) (_entity.Borrow, error)
	UpdateProcure(reqData _entity.Procure) (_entity.Procure, error)
	UpdateBorrowByAdmin(reqData _entity.Borrow) (_entity.Borrow, error)
	UpdateProcureByAdmin(reqData _entity.Procure) (_entity.Procure, error)
}
