//import { startOfHour, isBefore, getHours, format } from 'date-fns';
import { injectable, inject } from 'tsyringe';
import hash from 'object-hash';
import { isAfter, addMinutes } from 'date-fns';

import AppError from '../../../shared/errors/AppError';

import Product from '../infra/typeorm/schemas/Product';

import IProductsRepository from '../repositories/IProductsRepository';

type IRequest = Array<{
  id: string;
  name: string;
}>;

@injectable()
class ImportProductService {
  constructor(
    @inject('ProductsRepository')
    private productRepository: IProductsRepository,
  ) {}

  public async execute(objProducts: IRequest): Promise<Product> {
    const dataHashed = hash(objProducts, { algorithm: 'md5' });
    const findProductsHashed = await this.productRepository.findByHash(
      dataHashed,
    );

    if (
      findProductsHashed &&
      !isAfter(Date.now(), addMinutes(findProductsHashed.created_at, 10))
    ) {
      throw new AppError(
        'This products was already imported in last 10 minutes.',
        403,
      );
    }

    const product = await this.productRepository.create({
      objProducts,
      dataHashed,
    });

    return product;
  }
}

export default ImportProductService;
