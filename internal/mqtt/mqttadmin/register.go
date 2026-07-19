package mqttadmin

import "context"

func (s *service) RegisterUser(
	ctx context.Context,
	username string,
	password string,
) error {

	return s.execute(

		ctx,

		"dynsec",

		"createClient",

		username,

		"-p",

		password,
	)
}

// func (s *service) RegisterUser(ctx context.Context, username, password string) error {

// 	//----------------------------------------------------------
// 	// Parse MQTT Broker
// 	//----------------------------------------------------------

// 	u, err := url.Parse(s.cfg.MQTTBROKER)

// 	if err != nil {
// 		return err
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
// 		"-h", host,
// 		"-p", port,
// 		"-u", s.cfg.MQTTUSERNAME,
// 		"-P", s.cfg.MQTTPASSWORD,
// 		"dynsec",
// 		"createClient",
// 		username,
// 		"-p",
// 		password,
// 	)

// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return fmt.Errorf(
// 			"failed to register MQTT user: %v\n%s",
// 			err,
// 			string(output),
// 		)
// 	}
// 	return nil
// }
