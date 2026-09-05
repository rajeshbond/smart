package machinehandler

import (
	"github.com/go-chi/jwtauth/v5"
	machineservice "github.com/rajeshbond/smart/internal/http/machine_imm/machine_service"
)

type MachineHandler struct {
	MachineService *machineservice.MachineService
	tokenAuth      *jwtauth.JWTAuth
}

func NewMachineHandler(machineService *machineservice.MachineService, tokenAuth *jwtauth.JWTAuth) *MachineHandler {
	return &MachineHandler{
		tokenAuth:      tokenAuth,
		MachineService: machineService,
	}
}
