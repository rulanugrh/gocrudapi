package portrepo

import "gocrudapi/entity"

type BookInterface interface {
	Create(req entity.Book) (*entity.Book, error)
	Get(id uint) (*entity.Book, error)
	GetAll() ([]entity.Book, error)
	Update(id uint, req entity.Book) (*entity.Book, error)
	Delete(id uint) error
}
