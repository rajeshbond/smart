package service

import "github.com/rajeshbond/smart/internal/http/assembly_master/store"

type Service struct {
	Store *store.Store
}

func NewService(store *store.Store) *Service {
	return &Service{Store: store}
}
