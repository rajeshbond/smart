/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : service.go
 *
 ******************************************************************************/

package service

import (
	"github.com/rajeshbond/smart/internal/http/device/device_master/store"
	"github.com/rajeshbond/smart/internal/mqtt/mqttadmin"
)

//------------------------------------------------------------------------------
// Service
//------------------------------------------------------------------------------

type Service struct {
	Store *store.Store

	mqtt mqttadmin.Service
}

//------------------------------------------------------------------------------
// Constructor
//------------------------------------------------------------------------------

func NewService(
	store *store.Store,
	mqtt mqttadmin.Service,
) *Service {

	return &Service{

		Store: store,

		mqtt: mqtt,
	}
}
