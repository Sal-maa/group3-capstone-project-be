package asset

import (
	_entity "capstone/be/entity"
)

type Asset interface {
	Create(assetData _entity.Asset) (createAsset _entity.Asset, code int, err error)
	GetById(id int) (asset _entity.Asset, code int, err error)
	GetAssetByCategory(category string, page int) (asset _entity.AssetSimplified, code int, err error)
	GetAll(page int) (asset []_entity.AssetSimplified, code int, err error)
	Update(assetData _entity.Asset) (updateAsset _entity.Asset, code int, err error)
	// Delete(id int) (code int, err error)
}
