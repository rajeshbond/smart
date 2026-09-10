package depthandler

import (
	"github.com/go-chi/jwtauth/v5"
	deptservice "github.com/rajeshbond/smart/internal/http/department_master/dept_service"
)

type DeptHandler struct {
	DeptService *deptservice.DeptService
	tokenAuth   *jwtauth.JWTAuth
}

func NewDeptHandler(service *deptservice.DeptService, tokenAuth *jwtauth.JWTAuth) *DeptHandler {
	return &DeptHandler{
		tokenAuth:   tokenAuth,
		DeptService: service,
	}
}
