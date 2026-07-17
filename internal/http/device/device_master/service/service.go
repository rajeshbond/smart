/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : service.go
 *
 * DESCRIPTION :
 * Device Master Service
 *
 ******************************************************************************/

package service

import (
	"github.com/rajeshbond/smart/internal/http/device/device_master/store"
)

//------------------------------------------------------------------------------
// Service
//------------------------------------------------------------------------------

type Service struct {
	store *store.Store
}

//------------------------------------------------------------------------------
// Constructor
//------------------------------------------------------------------------------

func NewService(
	store *store.Store,
) *Service {

	return &Service{
		store: store,
	}
}