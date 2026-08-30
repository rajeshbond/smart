package store

import (
	"context"

	"github.com/rajeshbond/smart/internal/mqtt/imm/dto"
)

type ProductionStore interface {
	Save(ctx context.Context, req *dto.ProductionDTO) error
}
