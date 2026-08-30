package service

import (
	"github.com/rajeshbond/smart/internal/mqtt/imm/store"
)

// ============================================================
// IMM SERVICE
// ============================================================

type ImmService struct {
	immStore *store.ImmProductionStore
}

func NewImmService(
	immStore *store.ImmProductionStore,
) ImmService {
	return ImmService{
		immStore: immStore,
	}
}
