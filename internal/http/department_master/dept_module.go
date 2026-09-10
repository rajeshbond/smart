package departmentmaster

import (
	"database/sql"

	"github.com/go-chi/jwtauth/v5"
	depthandler "github.com/rajeshbond/smart/internal/http/department_master/dept_handler"
	deptservice "github.com/rajeshbond/smart/internal/http/department_master/dept_service"
	deptstore "github.com/rajeshbond/smart/internal/http/department_master/dept_store"
)

type DeptModeule struct {
	DeptStore   *deptstore.DeptStore
	DeptService *deptservice.DeptService
	DeptHandler *depthandler.DeptHandler
	tokenAuth   *jwtauth.JWTAuth
}

func NewDeptModule(db *sql.DB, tokenAuth *jwtauth.JWTAuth) *DeptModeule {
	deptStore := deptstore.NewDeptStore(db)
	deptService := deptservice.NewDeptService(deptStore)
	depthandler := depthandler.NewDeptHandler(deptService, tokenAuth)

	return &DeptModeule{
		tokenAuth:   tokenAuth,
		DeptStore:   deptStore,
		DeptService: deptService,
		DeptHandler: depthandler,
	}
}
