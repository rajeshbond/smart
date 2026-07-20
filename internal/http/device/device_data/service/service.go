package service

import (
	"github.com/rajeshbond/smart/internal/http/device/device_data/store"
)

type Service struct {
	Store *store.Store
}

func NewService(store *store.Store) *Service {
	return &Service{Store: store}
}
