package usecase

import (
	"github.com/acs-fernando/produto/domain/model"
)

type ProdutoUseCase struct {
	ProdutoKeyRepository model.ProdutoKeyRepositoryInterface
}

func (p *ProdutoUseCase) CreateProduct(name string, description string, price float32) (*model.Produtos, error) {
	produtoKey, err := model.NewProdutoKey(name, description, price)
	if err != nil {
		return &model.Produtos{}, err
	}

	p.ProdutoKeyRepository.CreateProduct(produtoKey)
	if produtoKey.ID == 0 {
		return &model.Produtos{}, err
	}

	return produtoKey, nil
}

func (p *ProdutoUseCase) FindProducts() ([]model.Produtos, error) {
	produtoKey, err := p.ProdutoKeyRepository.FindProducts()
	if err != nil {
		return nil, err
	}
	return produtoKey, nil
}
