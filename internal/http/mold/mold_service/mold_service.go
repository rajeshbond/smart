package moldservice

import moldstore "github.com/rajeshbond/smart/internal/http/mold/mold_store"

type MoldSerice struct {
	MoldStore *moldstore.MoldStore
}

func NewMoldService(moldStore *moldstore.MoldStore) *MoldSerice {
	return &MoldSerice{MoldStore: moldStore}
}
