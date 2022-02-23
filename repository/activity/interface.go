package activity

import (
	_entity "capstone/be/entity"
)

type Activity interface {
	GetAllActivityOfUser(user_id int) (activities []_entity.ActivitySimplified, code int, err error)
	GetDetailActivityByRequestId(request_id int) (activity _entity.Activity, code int, err error)
	CancelRequest(request_id int) (code int, err error)
	ReturnRequest(request_id int) (code int, err error)
}
