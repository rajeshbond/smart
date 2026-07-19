package mqttadmin

import "context"

type Service interface {
	RegisterUser(ctx context.Context, username string, password string) error

	DeleteUser(ctx context.Context, username string) error
}
