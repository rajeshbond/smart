package service

import (
	"context"

	"github.com/rajeshbond/smart/internal/mqtt/imm/dto"
)

type ImmServiceInterface interface {
	Save(ctx context.Context, req *dto.ProductionDTO) error
}
