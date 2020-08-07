import Product from '../infra/typeorm/schemas/Product';

import ICreateProductDTO from '@modules/products/dtos/ICreateProductDTO';

export default interface IProductsRepository {
  create(objProducts: ICreateProductDTO): Promise<Product>;
  findByHash(hash: string): Promise<Product | undefined>;
}
