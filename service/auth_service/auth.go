package auth_service

import (
	"template/model"
	"template/pkg/util"
)

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	password := util.EncodeMD5(a.Password)
	return model.CheckUser(a.Username, password)
}
func (a *Auth) Register() (bool, error) {
	password := util.EncodeMD5(a.Password)
	return model.Register(a.Username, password)
}
func (a *Auth) IsExist() bool {
	return model.IsExist(a.Username)
}
