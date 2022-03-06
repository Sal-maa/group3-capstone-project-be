package asset

import (
	_entity "capstone/be/entity"
)

type Asset interface {
	Create(assetData _entity.Asset) (code int, err error)
	GetAll(category string, status string) (assets []_entity.AssetSimplified, code int, err error)
	GetByShortName(short_name string) (total int, maintenance int, asset _entity.AssetSimplified, code int, err error)
	SetMaintenance(short_name string) (code int, err error)
	SetAvailable(short_name string) (code int, err error)
	GetCategoryId(category string) (id int, code int, err error)
	GetStats() (statistics _entity.Statistics, code int, err error)
}
