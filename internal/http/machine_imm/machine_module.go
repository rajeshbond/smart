package machine

import (
	"database/sql"

	"github.com/go-chi/jwtauth/v5"
	machinehandler "github.com/rajeshbond/smart/internal/http/machine_imm/machine_handler"
	machineservice "github.com/rajeshbond/smart/internal/http/machine_imm/machine_service"
	machinestore "github.com/rajeshbond/smart/internal/http/machine_imm/machine_store"
)

type MachineModule struct {
	MachineStore   *machinestore.MachineStore
	MachineService *machineservice.MachineService
	MachineHandler *machinehandler.MachineHandler
	tokenAuth      *jwtauth.JWTAuth
}

func NewMachineModule(db *sql.DB, tokenAuth *jwtauth.JWTAuth) *MachineModule {
	machineStore := machinestore.NewMachineStore(db)
	machineservice := machineservice.NewMachineSerice(machineStore)
	machineHandler := machinehandler.NewMachineHandler(machineservice, tokenAuth)

	return &MachineModule{
		tokenAuth:      tokenAuth,
		MachineStore:   machineStore,
		MachineService: machineservice,
		MachineHandler: machineHandler,
	}
}
