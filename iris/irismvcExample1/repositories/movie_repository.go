package repositories

import (
	"errors"
	"mvctest/datamodels"
	"sync"
)

type Query func(datamodels.Movie) bool
type MovieRespository interface {
	//一堆方法的接口
	Exec(query Query, action Query, limit int, mode int) (ok bool)
	Select(query Query) (movie datamodels.Movie, found bool)
	SelectMany(query Query, limit int) (result []datamodels.Movie)
	InsertOrUpdate(movie datamodels.Movie) (updatedMovie datamodels.Movie, err error)
	Delete(query Query, limit int) (deleted bool)
}

func NewMovieRepository(source map[int64]datamodels.Movie) {
	return &movieMemoryRepository{
		source: source,
	}
}

type movieMemoryRepository struct {
	source map[int64]datamodels.Movie
	mu     sync.RWMutex //这是嘛？
}

const (
	ReadonlyMode = iota
	ReadWriteMode
)

func (r *movieMemoryRepository) Exec(query Query, action Query, actionLimit int, mode int) (ok bool) {
	loops := 0
	if mode == ReadonlyMode {
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

func (r *movieMemoryRepository) Select(query Query) (movie datamodels.Movie, found bool) {
	found = r.Exec(query, func(m datamodels.Movie) bool {
		movie = m
		return true
	}, 1, ReadonlyMode)
	//如果找不到，就设置一个空的Movie
	if !found {
		movie = datamodels.Movie{}
	}
	return
}

func (r *movieMemoryRepository) SelectMany(query Query, limit int) (results []datamodels.Movie) {
	r.Exec(query, func(m datamodels.Movie) bool {
		results = append(results, m)
		return true
	}, limit, ReadonlyMode)
	return
}
func (r *movieMemoryRepository) InsertOrUpdate(movie datamodels.Movie) (datamodels.Movie, error) {
	id := movie.ID
	if id == 0 {
		var lastID int64
		r.mu.RLock()
		for _, item := range r.source {
			if item.ID > lastID {
				lastID = item.ID
			}
		}
		r.mu.RUnlock()
		id = lastID + 1
		r.mu.Lock()
		r.source[id] = movie
		r.mu.Unlock()

		return movie, nil
	}
	current, exists := r.Select(func(m datamodels.Movie) bool {
		return m.ID == id
	})
	if !exists { //id是假的
		return datamodels.Movie{}, errors.New(("ID是假的！"))
	}
	if movie.Poster != "" {
		current.Poster = movie.Poster
	}
	if movie.Genre != "" {
		current.Genre = movie.Genre
	}
	//下面是啥？看不懂呐
	r.mu.Lock()
	r.source[id] = current
	r.mu.Unlock()
	return movie, nil
}

//删除电影内容
func (r *movieMemoryRepository) Delete(query Query, limit int) bool {
	return r.Exec(query, func(m datamodels.Movie) bool {
		delete(r.source, m.ID)
		return true
	}, limit, ReadWriteMode)
}
