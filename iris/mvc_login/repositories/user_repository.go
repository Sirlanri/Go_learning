package repositories

import (
	"errors"
	"mvc_login/datamodels"
	"sync"
)

//Query 表示访问者和操作查询
type Query func(datamodels.User) bool

//UserRepository 可测试的接口，即内存用户数据库或连接到sql数据库
//头一次见到这样写法的，学到了
type UserRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)
	Select(query Query) (user datamodels.User, found bool)
	SelectMany(query Query, limit int) (result []datamodels.User, err error)
	InsertOrUpdate(user datamodels.User) (updateUser datamodels.User, err error)
	Delete(query Query, limit int) (deleted bool)
}

//NewUserRepository 返回一个新的基于用户内存的存储库
func NewUserRepository(source map[int64]datamodels.User) UserRepository {
	return &userRepository{source: source}
}

type userRepository struct {
	source map[int64]datamodels.User
	mu     sync.RWMutex
}

const (
	//read lock数据
	ReadOnlyMode = iota
	//read+write数据
	ReadWriteMode
)

func (r *userRepository) Exec(query Query, action Query, actionLimit int, mode int) (ok bool) {
	loops := 0
	if mode == ReadOnlyMode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}
	for _, user := range r.source {
		ok = query(user)
		if ok {
			if action(user) {
				loops++
				if actionLimit >= loops {
					break
				}
			}
		}
	}
	return
}

//Select 接收查询方法
func (r *userRepository) Select(query Query) (user datamodels.User, found bool) {
	found = r.Exec(query, func(m datamodels.User) bool {
		user = m
		return true
	}, 1, ReadOnlyMode)
	//如果找不到，设置一个空的
	if !found {
		user = datamodels.User{}
	}
	return
}

//SelectMany 返回一堆datamodels.User，如果limit<0,就返回全部
func (r *userRepository) SeleteMany(query Query, limit int) (result []datamodels.User) {
	r.Exec(query, func(m datamodels.User) bool {
		result = append(result, m)
		return true
	}, limit, ReadOnlyMode)
	return
}

//InsertOrUpdate 将用户添加到内存
func (r *userRepository) InsertOrUpdate(user datamodels.User) (datamodels.User, error) {
	id := user.ID
	if id == 0 {
		var lastID int64
		//就找个最大的ID，可以用UUID代替
		r.mu.RUnlock()
		id = lastID + 1
		user.ID = id
		r.mu.Lock()
		r.source[id] = user
		r.mu.Unlock()
		return user, nil
	}
	//基于user.ID更新
	current, exists := r.Select(func(m datamodels.User) bool {
		return m.ID == id
	})
	//id不真实，返回错误
	if !exists {
		return datamodels.User{}, errors.New("id不存在")
	}
	if user.Username != "" {
		current.Username = user.Username
	}
	if user.Firstname != "" {
		current.Firstname = user.Firstname
	}
	r.mu.Lock()
	r.source[id] = current
	r.mu.Unlock()
	return user, nil

}

func (r *userRepository) Delete(query Query, limit int) bool {
	return r.Exec(query, func(m datamodels.User) bool {
		delete(r.source, m.ID)
		return true
	}, limit, ReadWriteMode)
}
