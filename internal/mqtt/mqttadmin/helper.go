/******************************************************************************
 *
 * MODULE      : MQTT Admin
 * FILE        : helper.go
 *
 ******************************************************************************/

package mqttadmin

import (
	"context"
	"fmt"
	"net/url"
	"os/exec"
)

func (s *service) execute(ctx context.Context, args ...string) error {
	//----------------------------------------------------------
	// Parse Broker
	//----------------------------------------------------------

	u, err := url.Parse(s.cfg.MQTTBROKER)

	if err != nil {
		return fmt.Errorf("invalid MQTT_BROKER : %w", err)
	}

	host := u.Hostname()

	port := u.Port()

	if port == "" {
		port = "1883"
	}

	//----------------------------------------------------------
	// Common Arguments
	//----------------------------------------------------------

	command := []string{

		"-h",
		host,

		"-p",
		port,

		"-u",
		s.cfg.MQTTUSERNAME,

		"-P",
		s.cfg.MQTTPASSWORD,
	}

	command = append(command, args...)

	cmd := exec.CommandContext(

		ctx,

		"mosquitto_ctrl",

		command...,
	)

	output, err := cmd.CombinedOutput()

	if err != nil {

		return fmt.Errorf(

			"%v\n%s",

			err,

			string(output),
		)
	}

	return nil

}
