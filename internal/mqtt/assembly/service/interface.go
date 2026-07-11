package service

import (
	"context"

	dto "github.com/rajeshbond/smart/internal/mqtt/assembly/production_dto"
)

type ProductionSerive interface {
	Save(ctx context.Context, req *dto.ProductionDTO) error
}
