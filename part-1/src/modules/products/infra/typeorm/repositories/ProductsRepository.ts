import { getMongoRepository, MongoRepository } from 'typeorm';
import hash from 'object-hash';

import IProductsRepository from '@modules/products/repositories/IProductsRepository';

import ICreateProductDTO from '@modules/products/dtos/ICreateProductDTO';

import Product from '../schemas/Product';

class ProductsRepository implements IProductsRepository {
  private ormRepository: MongoRepository<Product>;

  constructor() {
    this.ormRepository = getMongoRepository(Product);
  }

  public async findByHash(hash: string): Promise<Product | undefined> {
    const findProducts = await this.ormRepository.findOne({
      where: { hash },
      order: { created_at: 'DESC' },
    });

    return findProducts;
  }

  public async create({ objProducts }: ICreateProductDTO): Promise<Product> {
    const dataHashed = hash(objProducts, { algorithm: 'md5' });

    const product = this.ormRepository.create({
      obj_products: objProducts,
      hash: dataHashed,
    });

    await this.ormRepository.save(product);

    return product;
  }
}

export default ProductsRepository;
