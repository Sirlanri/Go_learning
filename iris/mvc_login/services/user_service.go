package services

import (
	"errors"
	"mvc_login/datamodels"
	"mvc_login/repositories"
)

//UserService 处理用户数据模型的CRUID操作，
type UserService interface {
	GetAll() []datamodels.User
	GetByID(id int64) (datamodels.User, bool)
	GetByUserNameAndPassword(username, userpassword string) (datamodels.User, bool)
	DeleteByID(id int64) bool
	Update(id int64, user datamodels.User) (datamodels.User, error)
	UpdatePassword(id int64, newPassword string) (datamodels.User, error)
	UpdateUserName(id int64, user datamodels.User) (datamodels.User, error)
	Create(userpassword string, user datamodels.User) (datamodels.User, error)
}

//NewUserService 返回默认用户服务
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

type userService struct {
	repo repositories.UserRepository
}

//GetAll 返回全部用户
func (s *userService) GetAll() []datamodels.User {
	return s.repo.SelectMany(func(_ datamodels.User) bool {
		return true
	}, -1)
}

//GetByID 根据ID返回用户
func (s *userService) GetByID(ID int64) (datamodels.User, bool) {
	return s.repo.Select(func(m datamodels.User) bool {
		return m.ID == ID
	})
}

//GetByUsernameAndPassword 根据用户名和密码返回用户，用于身份验证。
func (s *userService) GetByUserNameAndPassword(username, userpassword string) (datamodels.User, bool) {
	if username == "" || userpassword == "" {
		return datamodels.User{}, false
	}
	return s.repo.Select(func(m datamodels.User) bool {
		if m.Username == username {
			hashed := m.HashedPassword
			if ok, _ := datamodels.ValidatePassword(userpassword, hashed); ok {
				return true
			}
		}
		return false
	})
}

//Update 更新现有用户的每个字段？？？？？
func (s *userService) Update(id int64, user datamodels.User) (datamodels.User, error) {
	user.ID = id
	return s.repo.InsertOrUpdate(user)
}

//UpdatePassword 更新密码
func (s *userService) UpdatePassword(id int64, newPassword string) (datamodels.User, error) {
	//更新用户并返回
	hashed, err := datamodels.GeneratePassword(newPassword)
	if err != nil {
		return datamodels.User{}, err
	}
	return s.Update(id, datamodels.User{
		HashedPassword: hashed,
	})
}

//UpdateUsername 更新用户名
func (s *userService) UpdateUsername(id int64, newname string) (datamodels.User, error) {
	return s.Update(id, datamodels.User{
		Username: newname,
	})
}

//Create 创建插入新用户
func (s *userService) Create(userpassword string, user datamodels.User) (datamodels.User, error) {
	if user.ID > 0 || userpassword == "" || user.Firstname == "" || user.Username == "" {
		return datamodels.User{}, errors.New("无法创建此用户")
	}
	hashed, err := datamodels.GeneratePassword(userpassword)
	if err != nil {
		return datamodels.User{}, err
	}
	user.HashedPassword = hashed
	return s.repo.InsertOrUpdate(user)
}

//DeleteByID 按照ID删除用户。成功删除返回true
func (s *userService) DeleteByID(id int64) bool {
	return s.repo.Delete(func(m datamodels.User) bool {
		return m.ID == id
	}, 1)
}
