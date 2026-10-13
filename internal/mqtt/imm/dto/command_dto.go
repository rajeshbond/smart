package dto

// ============================================================
// IMM COMMAND REQUEST
// ============================================================
//
// This is the HTTP request received by the Go API.
//
// Example:
//
// POST /api/v1/imm/devices/04@xoom/command
//
// {
//     "command": "SET_MOLD_NO",
//     "mold_no": "MOLD-002"
// }
// ============================================================

type IMMCommandRequest struct {
	Command   string `json:"command"`
	MoldNo    string `json:"mold_no,omitempty"`
	MachineID string `json:"machine_id,omitempty"`
}

// ============================================================
// IMM COMMAND RESPONSE
// ============================================================

type IMMCommandResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	DeviceID string `json:"device_id"`
	Topic    string `json:"topic"`
}
