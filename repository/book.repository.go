package repository

import (
	"context"
	"fmt"
	"gocrudapi/entity"
	portrepo "gocrudapi/repository/portRepo"
	"log"

	"gorm.io/gorm"
)

type bookrepo struct {
	DB *gorm.DB
}

func NewBookRepository(DB *gorm.DB) portrepo.BookInterface {
	return &bookrepo{
		DB: DB,
	}
}

func (repo *bookrepo) Create(req entity.Book) (*entity.Book, error) {
	err := repo.DB.Create(&req).Error
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (repo *bookrepo) Get(id uint) (*entity.Book, error) {
	var book entity.Book
	err := repo.DB.Where("id = ?", id).Find(&book).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (repo *bookrepo) GetAll() ([]entity.Book, error) {
	var getBook []entity.Book
	err := repo.DB.Find(&getBook).Error
	if err != nil {
		return nil, err
	}

	return getBook, nil
}

func (repo *bookrepo) Update(id uint, req entity.Book) (*entity.Book, error) {
	var book entity.Book

	err := repo.DB.WithContext(context.Background()).Model(&req).Where("id = ?", id).Updates(&book).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Println(book)
	return &book, nil
}

func (repo *bookrepo) Delete(id uint) error {
	var book entity.Book
	err := repo.DB.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}
