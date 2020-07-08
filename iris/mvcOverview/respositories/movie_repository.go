package repositories

import (
	"errors"
	"mvcoverview/datamodels"
	"sync"
)

//Query 表示访问者个操作查询
type Query func(datamodels.Movie) bool

//MovieRepository 处理电影实体，模型的基本操作
//是一个可测试接口，即一个内存电影库或连接到sql数据库
type MovieRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)
	Select(query Query) (movie datamodels.Movie, found bool)
	SelectMany(query Query, limit int) (results []datamodels.Movie)
	InsertOrUpdate(movie datamodels.Movie) (updatedMovie datamodels.Movie, err error)
	Delete(query Query, limit int) (deleted bool)
}

//NewMovieRepository 返回一个新的基于电影内存的repository，
//我们示例中唯一的repository类型。
func NewMovieRepository(source map[int64]datamodels.Movie) MovieRepository {
	return &movieMemoryRepository{source: source}
}

// movieMemoryRepository 是一个MovieRepository
//使用内存数据源（map）管理电影。
type movieMemoryRepository struct {
	source map[int64]datamodels.Movie
	mu     sync.RWMutex
}

//好熟悉的操作，和login一模一样
const (
	// ReadOnlyMode 将RLock（读取）数据。
	ReadOnlyMode = iota
	// ReadWriteMode 将锁定（读/写）数据。
	ReadWriteMode
)

func (r *movieMemoryRepository) Exec(query Query, action Query, actionLimit int, mode int) (ok bool) {
	loops := 0
	if mode == ReadOnlyMode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}
	for _, movie := range r.source {
		ok = query(movie)
		if ok {
			if action(movie) {
				loops++
				if actionLimit >= loops {
					break
				}
			}
		}
	}
	return
}

//Select 接收查询函数 为内部的每个电影模型触发
func (r *movieMemoryRepository) Select(query Query) (movie datamodels.Movie, found bool) {
	found = r.Exec(query, func(m datamodels.Movie) bool {
		movie = m
		return true
	}, 1, ReadOnlyMode)
	//如果找不到，设置一个空的Movie
	if !found {
		movie = datamodels.Movie{}
	}
	return
}

//SelectMany 返回多个movie作为切片
func (r *movieMemoryRepository) SelectMany(query Query, limit int) (results []datamodels.Movie) {
	r.Exec(query, func(m datamodels.Movie) bool {
		results = append(results, m)
		return true
	}, limit, ReadOnlyMode)
	return
}

//InsertOrUpdate 将movie添加或更新到map中存储。
func (r *movieMemoryRepository) InsertOrUpdate(movie datamodels.Movie) (datamodels.Movie, error) {
	id := movie.ID
	if id == 0 {
		var lastid int64
		//找到最大的ID，防止重复
		r.mu.RLock()
		for _, item := range r.source {
			if item.ID > lastid {
				lastid = item.ID
			}
		}
		r.mu.RUnlock()
		id = lastid + 1
		movie.ID = id
		r.mu.Lock()
		r.source[id] = movie
		return movie, nil
	}
	//基于movie.ID更新动作
	current, exists := r.Select(func(m datamodels.Movie) bool {
		return m.ID == id
	})
	if !exists {
		return datamodels.Movie{}, errors.New("不能升级一个已经存在的")
	}
	if movie.Poster != "" {
		current.Poster = movie.Poster
	}
	if movie.Genre != "" {
		current.Genre = movie.Genre
	}
	//锁定数据
	r.mu.Lock()
	r.source[id] = current
	r.mu.Unlock()
	return movie, nil
}

//Delete 删除电影
func (r *movieMemoryRepository) Delete(query Query, limit int) bool {
	return r.Exec(query, func(m datamodels.Movie) bool {
		delete(r.source, m.ID)
		return true
	}, limit, ReadWriteMode)
}
