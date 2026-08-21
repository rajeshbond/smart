package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/http/assembly_master/dto"
)

func (s *Service) CreateAssemblyMaster(ctx context.Context, req dto.CreateAssemblyMasterRequest, claims *auth.UserClaims) (*dto.AssemblyMasterResponse, error) {
	// Basic Validation
	req.MachineID = strings.TrimSpace(req.MachineID)
	if req.MachineID == "" {
		return nil, fmt.Errorf("machine id is required")
	}

	req.AssemblyName = strings.TrimSpace(req.AssemblyName)

	if req.AssemblyName == "" {
		return nil, fmt.Errorf("assembly_name is required")
	}

	req.DeviceID = strings.TrimSpace(req.DeviceID)

	if req.DeviceID == "" {
		return nil, fmt.Errorf("device_id is required")
	}

	req.Variant = strings.TrimSpace(req.Variant)

	if req.Variant == "" {
		req.Variant = "none"
	}

	if req.HourTargetOutput < 0 {
		return nil, fmt.Errorf("hour_target_output cannot be negative")
	}

	return s.Store.CreateAssemblyMaster(ctx, req, claims.TenantID, claims.UserID)

}
