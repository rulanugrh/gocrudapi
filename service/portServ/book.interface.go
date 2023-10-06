package portserv

import "gocrudapi/entity"

type BookInterface interface {
	Create(req entity.Book) (*entity.ResponseBook, error)
	Get(id uint) (*entity.ResponseBook, error)
	GetAll() ([]entity.ResponseBook, error)
	Update(id uint, req entity.Book) (*entity.ResponseBook, error)
	Delete(id uint) error
}
