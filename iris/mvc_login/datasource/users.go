package datasource

import (
	"errors"
	"mvc_login/datamodels"
)

type Engine uint32

const (
	Memory Engine = iota
	Bolt
	Mysql
)

//LoadUser 从内存中返回所有的用户
func LoadUser(engine Engine) (map[int64]datamodels.User, error) {
	if engine != Memory {
		return nil, errors.New("for the shake of simplicity we're using a simple map as the data source")
	}
	return make(map[int64]datamodels.User), nil
}
