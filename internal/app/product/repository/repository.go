package repository

import (
	"gorm.io/gorm"

	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/product/contract"
	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/product/entity"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) contract.Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repositoryImpl) FindByName(name string) ([]entity.Product, error) {
	var products []entity.Product
	if err := r.db.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repositoryImpl) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repositoryImpl) Create(product *entity.Product) (string, error) {
	if err := r.db.Create(product).Error; err != nil {
		return "", err
	}
	return product.ID, nil
}

func (r *repositoryImpl) Update(id string, updated *entity.Product) (string, error) {
	if err := r.db.Model(&entity.Product{}).Where("id = ?", id).Updates(updated).Error; err != nil {
		return "", err
	}
	return id, nil
}

//jangna lupa tambahin login pake laravel +php
