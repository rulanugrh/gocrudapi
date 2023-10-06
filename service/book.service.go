package service

import (
	"fmt"
	"gocrudapi/entity"
	portrepo "gocrudapi/repository/portRepo"
	portserv "gocrudapi/service/portServ"
	"log"
)

type bookstruct struct {
	repo portrepo.BookInterface
}

func NewBookService(repo portrepo.BookInterface) portserv.BookInterface {
	return &bookstruct{
		repo: repo,
	}
}

func (serv *bookstruct) Create(req entity.Book) (*entity.ResponseBook, error) {
	data, err := serv.repo.Create(req)
	if err != nil {
		return nil, err
	}

	web := entity.ResponseBook{
		Name:       data.Name,
		Year:       data.Year,
		Price:      data.Price,
		CategoryID: data.CategoryID,
	}

	return &web, nil
}

func (serv *bookstruct) Get(id uint) (*entity.ResponseBook, error) {
	data, err := serv.repo.Get(id)
	if err != nil {
		return nil, err
	}

	web := entity.ResponseBook{
		Name:       data.Name,
		Year:       data.Year,
		Price:      data.Price,
		CategoryID: data.CategoryID,
	}

	return &web, nil
}
func (serv *bookstruct) GetAll() ([]entity.ResponseBook, error) {
	var listBook []entity.ResponseBook
	data, err := serv.repo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, res := range data {
		book := entity.ResponseBook{
			Name:       res.Name,
			Year:       res.Year,
			Price:      res.Price,
			CategoryID: res.CategoryID,
		}

		listBook = append(listBook, book)
	}

	return listBook, nil

}

func (serv *bookstruct) Update(id uint, req entity.Book) (*entity.ResponseBook, error) {
	data, err := serv.repo.Update(id, req)
	if err != nil {
		log.Println(err)
		return nil, err

	}

	web := entity.ResponseBook{
		Name:       data.Name,
		Year:       data.Year,
		Price:      data.Price,
		CategoryID: data.CategoryID,
	}

	fmt.Println(web)
	return &web, nil
}
func (serv *bookstruct) Delete(id uint) error {
	err := serv.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
