package api

import (
	"context"

	"opendax-clean/finex/api/proto/apipb"
	"opendax-clean/finex/ent"
	"opendax-clean/finex/ent/asset"
	"opendax-clean/finex/ent/proto/entpb"
)

type AssetService struct {
	client *ent.Client
	apipb.UnimplementedAssetServiceServer
}

func NewAssetService(client *ent.Client) *AssetService {
	return &AssetService{
		client: client,
	}
}

func (svc *AssetService) Get(ctx context.Context, req *apipb.AssetGetRequest) (*entpb.Asset, error) {
	entity, err := svc.client.Asset.Get(context.Background(), req.GetId())
	if err != nil {
		return nil, err
	}
	return &entpb.Asset{
		Id:    entity.ID,
		Name:  entity.Name,
		Index: entity.Index,
	}, nil
}

func (svc *AssetService) List(ctx context.Context, req *apipb.AssetListRequest) (*apipb.AssetListResponse, error) {
	var entities []*ent.Asset
	var err error
	if ids := req.GetIds(); ids != nil {
		entities, err = svc.client.Asset.Query().Where(asset.IDIn(ids...)).All(ctx)
	} else {
		entities, err = svc.client.Asset.Query().All(ctx)
	}
	if err != nil {
		return nil, err
	}
	res := &apipb.AssetListResponse{
		Assets: make([]*entpb.Asset, len(entities)),
	}
	for i, e := range entities {
		res.Assets[i] = &entpb.Asset{
			Id:    e.ID,
			Name:  e.Name,
			Index: e.Index,
		}
	}
	return res, nil
}
