package grpc

import (
	"context"

	"github.com/acs-fernando/produto/application/grpc/pb"
	"github.com/acs-fernando/produto/application/usecase"
)

type ProdutoGrpcService struct {
	ProdutoUseCase usecase.ProdutoUseCase
	pb.UnimplementedProductServiceServer
}

func (p *ProdutoGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	prod, err := p.ProdutoUseCase.CreateProduct(in.Name, in.Description, in.Price)
	if err != nil {
		return &pb.CreateProductResponse{}, err
	}

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          prod.ID,
			Name:        prod.Name,
			Description: prod.Description,
			Price:       prod.Price,
		},
	}, nil
}

func (p *ProdutoGrpcService) FindProducts(ctx context.Context, in *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	produtos, err := p.ProdutoUseCase.FindProducts()

	if err != nil {
		return &pb.FindProductsResponse{}, err
	}

	produtosResp := make([]*pb.Product, len(produtos))

	for index, prod := range produtos {
		produtosResp[index] = &pb.Product{
			Id:          prod.ID,
			Name:        prod.Name,
			Description: prod.Description,
			Price:       prod.Price,
		}
	}

	return &pb.FindProductsResponse{
		Products: produtosResp,
	}, nil
}

func NewProdutoGrpcService(usecase usecase.ProdutoUseCase) *ProdutoGrpcService {
	return &ProdutoGrpcService{
		ProdutoUseCase: usecase,
	}
}
