package repository

import (
	"fmt"

	"github.com/acs-fernando/produto/domain/model"
	"github.com/jinzhu/gorm"
)

type ProdutoRepositoryDb struct {
	Db *gorm.DB
}

func (t ProdutoRepositoryDb) CreateProduct(produto *model.Produtos) (*model.Produtos, error) {
	err := t.Db.Create(produto).Error
	if err != nil {
		return nil, err
	}
	return produto, nil
}

func (t ProdutoRepositoryDb) FindProducts() ([]model.Produtos, error) {
	var produtos []model.Produtos
	t.Db.Find(&produtos)

	if produtos == nil {
		return nil, fmt.Errorf("no product was found")
	}
	return produtos, nil
}
