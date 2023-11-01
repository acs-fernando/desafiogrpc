import { Observable } from 'rxjs';

interface Product {
  id: number;
  name: string;
  description: string;
  price: number;
}

export interface RegisterProductRpcResponse {
  product: Product;
}

export interface FindProductRpcResponse {
  products: Product[];
}

export interface ProductClientGrpc {
  CreateProduct: (data: {
    name: string;
    description: string;
    price: number;
  }) => Observable<{ product: Product }>;
  FindProducts: (data: unknown) => Observable<FindProductRpcResponse>;
}
