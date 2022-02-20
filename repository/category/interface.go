package category

import (
	_entity "capstone/be/entity"
)

type Category interface {
	// Create(assetData _entity.Asset) (createAsset _entity.Asset, code int, err error)
	// GetById(id int) (asset _entity.Asset, code int, err error)
	GetAll() (category []_entity.Category, code int, err error)
	// Update(assetData _entity.Asset) (updateAsset _entity.Asset, code int, err error)
	// Delete(id int) (code int, err error)
}
