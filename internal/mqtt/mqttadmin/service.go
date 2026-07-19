package mqttadmin

import (
	"github.com/rajeshbond/smart/config"
)

type service struct {
	cfg *config.Config
}

func NewService(
	cfg *config.Config,
) Service {

	return &service{
		cfg: cfg,
	}
}
