package moldhandler

import (
	"github.com/go-chi/jwtauth/v5"
	moldservice "github.com/rajeshbond/smart/internal/http/mold/mold_service"
)

type MoldHanlder struct {
	MoldService *moldservice.MoldSerice
	tokenAuth   *jwtauth.JWTAuth
}

func NewMoldHandler(moldService *moldservice.MoldSerice, tokenAuth *jwtauth.JWTAuth) *MoldHanlder {
	return &MoldHanlder{
		tokenAuth:   tokenAuth,
		MoldService: moldService,
	}
}
