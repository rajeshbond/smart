package machineservice

import machinestore "github.com/rajeshbond/smart/internal/http/machine_imm/machine_store"

type MachineService struct {
	MachineStore machinestore.MachineStore
}

func NewMachineSerice(machineStore *machinestore.MachineStore) *MachineService {
	return &MachineService{MachineStore: *machineStore}
}
