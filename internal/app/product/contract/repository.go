package contract

import "github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/product/entity"

type Repository interface {
	FindAll() ([]entity.Product, error)
	FindByName(name string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Create(product *entity.Product) (string, error)
	Update(id string, product *entity.Product) (string, error)
}
