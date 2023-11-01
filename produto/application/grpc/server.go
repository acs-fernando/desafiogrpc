package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/acs-fernando/produto/application/grpc/pb"
	"github.com/acs-fernando/produto/application/usecase"
	"github.com/acs-fernando/produto/infrastructure/repository"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	produtoRepository := repository.ProdutoRepositoryDb{Db: database}
	produtoUseCase := usecase.ProdutoUseCase{ProdutoKeyRepository: produtoRepository}
	produtoGrpcService := NewProdutoGrpcService(produtoUseCase)
	pb.RegisterProductServiceServer(grpcServer, produtoGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}
