package history

import (
	_entity "capstone/be/entity"
)

type History interface {
	GetAllRequestHistoryOfUser(user_id int, page int) (count int, histories []_entity.UserRequestHistorySimplified, code int, err error)
	GetDetailRequestHistoryByRequestId(request_id int) (history _entity.UserRequestHistory, code int, err error)
	GetAllUsageHistoryOfAsset(short_name string) (asset _entity.AssetInfo, users []_entity.AssetUser, code int, err error)
}
