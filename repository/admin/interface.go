package admin

import (
	_entity "capstone/be/entity"
)

type Admin interface {
	GetAllAdminWaitingApproval(limit, offset int, category, date, order string) (requests []_entity.RequestResponse, total int, err error)
	GetAllAdminReturned(limit, offset int, category, date, order string) (requests []_entity.RequestResponse, total int, err error)
	GetAllAdmin(limit, offset int, status, category, date, order string) (requests []_entity.RequestResponse, total int, err error)
	GetAllManagerReturned(divLogin, limit, offset int, category, date, order string) (requests []_entity.RequestResponse, total int, err error)
	GetAllManager(divLogin, limit, offset int, status, category, date, order string) (requests []_entity.RequestResponse, total int, err error)
	GetAllProcure(limit, offset int, status, category, date, order string) (requests []_entity.Procure, total int, err error)
	GetUserDivision(id int) (divId int, code int, err error)
}
