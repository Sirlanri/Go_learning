package services

import (
	"mvctest/datamodels"
	"mvctest/repositories"
)

type MovieService interface {
	//定义方法的接口（虽然我现在还不知道这玩意儿有什么用
	Getall() []datamodels.Movie
	GetByID(id int64) (datamodels.Movie, bool)
	DeleteByID(id int64) bool
	UpdataPosterAndGenreByID(id int64, poster string, genre string) (datamodels.Movie, error)
}

func NewMovieService(repo repositories.MovieRespository) MovieService {
	return &MovieService{
		repo: repo,
	}
}

type movieService struct {
	//是个和接口同名的结构体，有啥用？我也不造_(:з」∠)_
	//好像下面那些方法，都属于结构体
	repo repositories.MovieRespository
}

//GetAll 获取所有的Movie
func (s *movieService) GetAll() []datamodels.Movie {
	return s.repo.SelectMany(func(_ datamodels.Movie) bool {
		return true
	}, -1)
}

//GetByID 根据ID返回最后一行
func (s *movieService) UpdataPosterAndGenreByID(id int64, poster string, genre string) (updatedMovie datamodels.Movie, err error) {
	//升级并返回Movie信息
	return s.repo.InsertOrUpdate(datamodels.Movie{
		ID:     id,
		Poster: poster,
		Genre:  genre,
	})
}

//删除电影
func (s *movieService) DeleteByID(id int64) bool {
	return s.repo.Delete(func(m datamodels.Movie) bool {
		return m.ID == id
	}, 1)
}
