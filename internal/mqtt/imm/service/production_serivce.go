package service

import (
	"context"
	"fmt"

	"github.com/rajeshbond/smart/internal/mqtt/imm/dto"
)

// ============================================================
// SAVE PRODUCTION
// ============================================================

func (s *ImmService) Save(
	ctx context.Context,
	req *dto.ProductionDTO,
) error {

	if req == nil {
		return fmt.Errorf("production request is nil")
	}

	return s.Save(ctx, req)
}
