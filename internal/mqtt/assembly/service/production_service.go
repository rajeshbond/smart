package service

import (
	"context"

	dto "github.com/rajeshbond/smart/internal/mqtt/assembly/production_dto"
	"github.com/rajeshbond/smart/internal/mqtt/assembly/store"
)

type productionService struct {
	store store.ProductionStore
}

func NewProductionService(store store.ProductionStore) ProductionSerive {
	return &productionService{
		store: store,
	}
}

func (ser *productionService) Save(ctx context.Context, req *dto.ProductionDTO) error {
	return ser.store.Save(ctx, req)
}
