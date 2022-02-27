package request

import (
	_entity "capstone/be/entity"
)

type Request interface {
	Borrow(reqData _entity.Borrow) (code int, err error)
	Procure(reqData _entity.Procure) (code int, err error)
	GetBorrowById(id int) (req _entity.Borrow, code int, err error)
	GetUserDivision(id int) (divId int, code int, err error)
	UpdateBorrow(reqData _entity.Borrow) (req _entity.Borrow, code int, err error)
	GetProcureById(id int) (req _entity.Procure, code int, err error)
	UpdateProcure(reqData _entity.Procure) (req _entity.Procure, code int, err error)

	UpdateBorrowByAdmin(reqData _entity.Borrow) (_entity.Borrow, error)
	UpdateProcureByAdmin(reqData _entity.Procure) (_entity.Procure, error)
}
