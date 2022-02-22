package history

import (
	_entity "capstone/be/entity"
)

type History interface {
	GetAllUsageHistoryOfUser(user_id int, page int) (histories []_entity.UserUsageHistorySimplified, code int, err error)
	GetDetailUsageHistoryByRequestId(request_id int) (history _entity.UserUsageHistory, code int, err error)
	GetAllUsageHistoryOfAsset(asset_id int) (histories []_entity.AssetUsageHistory, code int, err error)
}
