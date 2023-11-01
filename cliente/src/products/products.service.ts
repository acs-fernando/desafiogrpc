import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { CreateProductDto } from './dto/create-product.dto';
import { ClientGrpc } from '@nestjs/microservices';
import {
  FindProductRpcResponse,
  ProductClientGrpc,
  RegisterProductRpcResponse,
} from './product.grpc';
import { lastValueFrom } from 'rxjs';

@Injectable()
export class ProductsService implements OnModuleInit {
  private productGrpcService: ProductClientGrpc;

  constructor(
    @Inject('PRODUCTS_PACKAGE')
    private productGrpcPackage: ClientGrpc,
  ) {}

  onModuleInit() {
    this.productGrpcService =
      this.productGrpcPackage.getService('ProductService');
  }

  create(createProductDto: CreateProductDto) {
    return this.registerProduct(createProductDto);
  }

  findAll() {
    return this.findProducts('');
  }

  private async registerProduct(
    product: CreateProductDto,
  ): Promise<RegisterProductRpcResponse | null> {
    try {
      return await lastValueFrom(
        this.productGrpcService.CreateProduct(product),
      );
    } catch (e) {
      console.error(e);
      if (e.details == 'no product was found') {
        return null;
      }
      throw new ProductGrpcUnknownError('Grpc Internal Error');
    }
  }

  private async findProducts(
    data: unknown,
  ): Promise<FindProductRpcResponse | null> {
    try {
      return await lastValueFrom(this.productGrpcService.FindProducts(data));
    } catch (e) {
      console.error(e);
      if (e.details == 'no product was found') {
        return null;
      }
      throw new ProductGrpcUnknownError('Grpc Internal Error');
    }
  }
}

export class ProductGrpcUnknownError extends Error {}
