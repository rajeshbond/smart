/******************************************************************************
 *
 * MODULE      : MQTT Admin
 * FILE        : delete.go
 *
 * DESCRIPTION :
 * Delete MQTT User from Mosquitto Dynamic Security
 *
 ******************************************************************************/

package mqttadmin

import "context"

func (s *service) DeleteUser(
	ctx context.Context,
	username string,
) error {

	return s.execute(

		ctx,

		"dynsec",

		"deleteClient",

		username,
	)
}

// package mqttadmin

// import (
// 	"context"
// 	"fmt"
// 	"net/url"
// 	"os/exec"
// )

// func (s *service) DeleteUser(ctx context.Context, username string) error {
// 	//----------------------------------------------------------
// 	// Parse MQTT Broker
// 	//---------------------------------------------------------
// 	u, err := url.Parse(s.cfg.MQTTBROKER)
// 	if err != nil {
// 		return fmt.Errorf("invalid MQTT_BROKER: %w", err)
// 	}

// 	host := u.Hostname()

// 	port := u.Port()

// 	if port == "" {
// 		port = "1883"
// 	}

// 	//----------------------------------------------------------
// 	// Execute mosquitto_ctrl
// 	//----------------------------------------------------------
// 	cmd := exec.CommandContext(

// 		ctx,

// 		"mosquitto_ctrl",

// 		"-h",
// 		host,

// 		"-p",
// 		port,

// 		"-u",
// 		s.cfg.MQTTUSERNAME,

// 		"-P",
// 		s.cfg.MQTTPASSWORD,

// 		"dynsec",

// 		"deleteClient",

// 		username,
// 	)

// 	output, err := cmd.CombinedOutput()

// 	if err != nil {

// 		return fmt.Errorf(
// 			"failed to delete MQTT user: %v\n%s",
// 			err,
// 			string(output),
// 		)
// 	}

// 	return nil

// }
