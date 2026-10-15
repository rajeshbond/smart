package deptservice

import deptstore "github.com/rajeshbond/smart/internal/http/department_master/dept_store"

type DeptService struct {
	DeptStore *deptstore.DeptStore
}

func NewDeptService(store *deptstore.DeptStore) *DeptService {
	return &DeptService{
		DeptStore: store,
	}
}
