package mqtt

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/url"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"

	"github.com/rajeshbond/smart/config"
	"github.com/rajeshbond/smart/database"
)

func NewClient(
	cfg *config.Config,
	db *database.DB,
) paho.Client {

	opts := paho.NewClientOptions()
	opts.AddBroker(cfg.MQTTBROKER)

	// Unique Client ID to prevent kicking off VPS / existing instances
	uniqueID := fmt.Sprintf("%s-dev-%d", cfg.MQTTCLIENTID, time.Now().UnixNano())
	opts.SetClientID(uniqueID)

	opts.SetUsername(cfg.MQTTUSERNAME)
	opts.SetPassword(cfg.MQTTPASSWORD)

	opts.SetProtocolVersion(4) // MQTT 3.1.1
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(15 * time.Second)
	opts.SetKeepAlive(30 * time.Second)
	opts.SetPingTimeout(10 * time.Second)

	// ------------------------------------------------------------------
	// CRITICAL FOR SSL (Port 8883): Set explicit connect timeout for TLS
	// ------------------------------------------------------------------
	opts.SetConnectTimeout(15 * time.Second)
	opts.SetCleanSession(true)

	// Configure TLS parameters when using SSL / TLS port
	parsedURL, err := url.Parse(cfg.MQTTBROKER)
	if err == nil {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         parsedURL.Hostname(), // "mqtt.xoomgrid.com"
		}
		opts.SetTLSConfig(tlsConfig)
	} else {
		log.Printf("⚠️ Warning parsing MQTT_BROKER URL: %v", err)
	}

	opts.OnConnect = func(client paho.Client) {
		log.Println("✅ MQTT Connected")
		RegisterRoutes(client, db)
		log.Println("✅ MQTT Routes Registered")
	}

	opts.OnConnectionLost = func(client paho.Client, err error) {
		log.Printf("❌ MQTT Connection Lost : %v", err)
	}

	return paho.NewClient(opts)
}

// package mqtt

// import (
// 	"crypto/tls"
// 	"fmt"
// 	"log"
// 	"net/url"
// 	"time"

// 	paho "github.com/eclipse/paho.mqtt.golang"

// 	"github.com/rajeshbond/smart/config"
// 	"github.com/rajeshbond/smart/database"
// )

// func NewClient(
// 	cfg *config.Config,
// 	db *database.DB,
// ) paho.Client {

// 	opts := paho.NewClientOptions()
// 	opts.AddBroker(cfg.MQTTBROKER)

// 	// Generate dynamic Client ID to prevent collision between local dev and VPS
// 	uniqueID := fmt.Sprintf("%s-dev-%d", cfg.MQTTCLIENTID, time.Now().UnixNano())
// 	opts.SetClientID(uniqueID)

// 	opts.SetUsername(cfg.MQTTUSERNAME)
// 	opts.SetPassword(cfg.MQTTPASSWORD)

// 	opts.SetProtocolVersion(4) // MQTT 3.1.1
// 	opts.SetAutoReconnect(true)
// 	opts.SetMaxReconnectInterval(15 * time.Second)
// 	opts.SetKeepAlive(30 * time.Second)
// 	opts.SetPingTimeout(10 * time.Second)
// 	opts.SetCleanSession(true)

// 	// Dynamically parse ServerName from MQTTBROKER URL
// 	if parsedURL, err := url.Parse(cfg.MQTTBROKER); err == nil {
// 		tlsConfig := &tls.Config{
// 			InsecureSkipVerify: true,
// 			ServerName:         parsedURL.Hostname(), // Automatically extracts "mqtt.xoomgrid.com" or IP
// 		}
// 		opts.SetTLSConfig(tlsConfig)
// 	} else {
// 		log.Printf("⚠️ Warning: Could not parse MQTT_BROKER URL for TLS config: %v", err)
// 	}

// 	opts.OnConnect = func(client paho.Client) {
// 		log.Println("✅ MQTT Connected")
// 		RegisterRoutes(client, db)
// 		log.Println("✅ MQTT Routes Registered")
// 	}

// 	opts.OnConnectionLost = func(client paho.Client, err error) {
// 		log.Printf("❌ MQTT Connection Lost : %v", err)
// 	}

// 	return paho.NewClient(opts)
// }

// package mqtt

// import (
// 	"crypto/tls"
// 	"fmt"
// 	"log"
// 	"time"

// 	paho "github.com/eclipse/paho.mqtt.golang"

// 	"github.com/rajeshbond/smart/config"
// 	"github.com/rajeshbond/smart/database"
// )

// func NewClient(
// 	cfg *config.Config,
// 	db *database.DB,
// ) paho.Client {

// 	opts := paho.NewClientOptions()

// 	opts.AddBroker(cfg.MQTTBROKER)

// 	// ------------------------------------------------------------------
// 	// CRITICAL FIX: Append Nano Timestamp to ensure unique Client ID.
// 	// This prevents kicking off (or being kicked off by) the VPS process!
// 	// ------------------------------------------------------------------
// 	uniqueID := fmt.Sprintf("%s-dev-%d", cfg.MQTTCLIENTID, time.Now().UnixNano())
// 	opts.SetClientID(uniqueID)

// 	opts.SetUsername(cfg.MQTTUSERNAME)
// 	opts.SetPassword(cfg.MQTTPASSWORD)

// 	opts.SetProtocolVersion(4) // MQTT 3.1.1

// 	opts.SetAutoReconnect(true)
// 	opts.SetMaxReconnectInterval(15 * time.Second)
// 	opts.SetKeepAlive(30 * time.Second)
// 	opts.SetPingTimeout(10 * time.Second)
// 	opts.SetCleanSession(true)

// 	tlsConfig := &tls.Config{
// 		InsecureSkipVerify: true,
// 		ServerName:         "mqtt.xoomgrid.com",
// 	}
// 	opts.SetTLSConfig(tlsConfig)

// 	opts.OnConnect = func(client paho.Client) {
// 		log.Println("✅ MQTT Connected")
// 		RegisterRoutes(client, db)
// 		log.Println("✅ MQTT Routes Registered")
// 	}

// 	opts.OnConnectionLost = func(client paho.Client, err error) {
// 		log.Printf("❌ MQTT Connection Lost : %v", err)
// 	}

// 	return paho.NewClient(opts)
// }

// package mqtt

// import (
// 	"log"
// 	"time"

// 	paho "github.com/eclipse/paho.mqtt.golang"

// 	"github.com/rajeshbond/smart/config"
// 	"github.com/rajeshbond/smart/database"
// )

// func NewClient(
// 	cfg *config.Config,
// 	db *database.DB,
// ) paho.Client {

// 	opts := paho.NewClientOptions()

// 	opts.AddBroker(cfg.MQTTBROKER)
// 	opts.SetClientID(cfg.MQTTCLIENTID)
// 	opts.SetUsername(cfg.MQTTUSERNAME)
// 	opts.SetPassword(cfg.MQTTPASSWORD)

// 	opts.SetAutoReconnect(true)
// 	opts.SetMaxReconnectInterval(15 * time.Second)
// 	opts.SetKeepAlive(30 * time.Second)

// 	// Optional but recommended
// 	opts.SetCleanSession(false)

// 	// Called after FIRST connect and EVERY reconnect
// 	opts.OnConnect = func(client paho.Client) {

// 		log.Println("✅ MQTT Connected")

// 		RegisterRoutes(client, db)

// 		log.Println("✅ MQTT Routes Registered")
// 	}
// 	// -------------------------------------------------------------
// 	// CRITICAL FIX: Enable TLS/SSL Configuration for ssl:// brokers
// 	// -------------------------------------------------------------
// 	// tlsConfig := &tls.Config{
// 	// 	// Set to true to allow connecting to domain/ip during dev.
// 	// 	// Set to false in strict production if your SSL cert chain is fully trusted.
// 	// 	InsecureSkipVerify: true,
// 	// }
// 	// opts.SetTLSConfig(tlsConfig)
// 	opts.OnConnectionLost = func(client paho.Client, err error) {

// 		log.Printf("❌ MQTT Connection Lost : %v", err)

// 	}

// 	return paho.NewClient(opts)

// }
