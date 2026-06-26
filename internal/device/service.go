package device

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const (

	// path
	MosquittoDir       = `C:\Program Files\Mosquitto`
	MosquittoPasswdExe = MosquittoDir + `\mosquitto_passwd.exe`
	MosquittoPwdFile   = MosquittoDir + `\pwfile`

	// Commands

	ServiceCmd      = "net"
	ServiceStopArg  = "stop"
	ServiceStartArg = "start"
	ServiceName     = "mosquitto"
)

type DeviceService struct {
	Store *DeviceStore
}

func NewDeviceService(store *DeviceStore) *DeviceService {
	return &DeviceService{Store: store}
}

func (ser *DeviceService) RegisterDeviceRequest(ctx context.Context, req RegisterRequest) error {
	// Step 1: Write propertu to database clusetr first

	if err := ser.Store.CreateDevice(ctx, req); err != nil {
		return fmt.Errorf("database transaction aborted: %w", err)
	}

	// Step 2: Trigger the decoupled flat-file and service management helper

	if err := syncMosquitto(req.MQTTDeviceUsername, req.MQTTDevicePassword); err != nil {
		return fmt.Errorf("mosquitto sync failed: %w", err)
	}
	return nil
}

// Insolated system exection helper

func syncMosquitto(userName, password string) error {

	user := strings.TrimSpace(userName)
	pass := strings.TrimSpace(password)

	// 1.Append the Credential record straing into the broker flat-file

	cmd := exec.Command(MosquittoPasswdExe, "-b", MosquittoPwdFile, user, pass)

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("mosquitto_passwd execution error: %s", string(output))
	}

	// 2. Recycle the system daemon process to clear active security memory profiles

	_ = exec.Command(ServiceCmd, ServiceStopArg, ServiceStopArg, ServiceName).Run()

	if err := exec.Command(ServiceCmd, ServiceStartArg, ServiceName).Run(); err != nil {
		log.Printf("⚠️ Critical: Device saved to DB and pwfile, but Mosquitto service failed to start: %v", err)
		return fmt.Errorf("broker service reboot timed out: %w", err)
	}
	log.Printf("🔄 Broker successfully synchronized with new user credentials: %s", user)
	return nil

}
