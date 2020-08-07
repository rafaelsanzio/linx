import { ObjectID } from 'mongodb';
import hash from 'object-hash';

import IProductsRepository from '@modules/products/repositories/IProductsRepository';

import Product from '../../infra/typeorm/schemas/Product';
import ICreateProductDTO from '@modules/products/dtos/ICreateProductDTO';

class ProductsRepository implements IProductsRepository {
  private products: Product[] = [];

  public async findByHash(hash: string): Promise<Product | undefined> {
    const findProductsHashed = this.products.find(
      product => product.hash === hash,
    );

    return findProductsHashed;
  }

  public async create({ objProducts }: ICreateProductDTO): Promise<Product> {
    const product = new Product();

    const dataHashed = hash(objProducts, { algorithm: 'md5' });

    Object.assign(product, {
      id: new ObjectID(),
      obj_products: objProducts,
      hash: dataHashed,
      create_at: Date.now(),
    });

    this.products.push(product);

    return product;
  }
}

export default ProductsRepository;
