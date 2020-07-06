package repositories

import (
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
