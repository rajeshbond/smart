package service

import (
	"github.com/rajeshbond/smart/internal/http/device/device_data/shiftprovider"
	"github.com/rajeshbond/smart/internal/http/device/device_data/store"
)

type Service struct {
	Store         *store.Store
	shiftProvider shiftprovider.ShiftProvider
}

func NewService(store *store.Store, shiftProvider shiftprovider.ShiftProvider) *Service {
	return &Service{
		Store:         store,
		shiftProvider: shiftProvider,
	}
}
