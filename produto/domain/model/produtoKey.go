package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type ProdutoKeyRepositoryInterface interface {
	CreateProduct(produtoKey *Produtos) (*Produtos, error)
	FindProducts() ([]Produtos, error)
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Produtos struct {
	Base        `valid:"required"`
	Name        string  `json:"name" gorm:"type:varchar(30)" valid:"notnull"`
	Description string  `json:"description" gorm:"type:varchar(255)" valid:"notnull"`
	Price       float32 `json:"price" gorm:"type:float" valid:"notnull"`
}

func (p *Produtos) isValid() error {
	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return err
	}
	return nil
}

func NewProdutoKey(name string, description string, price float32) (*Produtos, error) {
	produtoKey := Produtos{
		Name:        name,
		Description: description,
		Price:       price,
	}
	produtoKey.CreatedAt = time.Now()
	err := produtoKey.isValid()
	if err != nil {
		return &Produtos{}, err
	}
	return &produtoKey, nil
}
