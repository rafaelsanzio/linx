import { hash } from 'bcryptjs';
import { addMinutes } from 'date-fns';

import AppError from '../../../shared/errors/AppError';

import FakeProductsRepository from '../repositories/fakes/FakeProductsRepository';
import ImportProductService from './ImportProductService';

let fakeProductsRepository: FakeProductsRepository;
let createProduct: ImportProductService;

type IProduct = Array<{
  id: string;
  name: string;
}>;

describe('Import Products', () => {
  beforeEach(() => {
    fakeProductsRepository = new FakeProductsRepository();
    createProduct = new ImportProductService(fakeProductsRepository);
  });

  it('should be able to import a new products', async () => {
    const objProducts: IProduct = [];

    const object1 = { id: '123', name: 'mesa' };
    const object2 = { id: '321', name: 'cadeira' };

    objProducts.push(object1, object2);

    const product = await createProduct.execute(objProducts);

    expect(product).toHaveProperty('id');
    expect(product.obj_products).toEqual([object1, object2]);
  });

  it('should not be able to import same products as before 10 minutes', async () => {
    const objProducts: IProduct = [];

    const object1 = { id: '123', name: 'mesa' };
    const object2 = { id: '321', name: 'cadeira' };

    objProducts.push(object1, object2);

    await createProduct.execute(objProducts);

    await expect(createProduct.execute(objProducts)).rejects.toBeInstanceOf(
      AppError,
    );
  });

  it('should be able to import same products as before after 10 minutes', async () => {
    const objProducts: IProduct = [];

    const object1 = { id: '123', name: 'mesa' };
    const object2 = { id: '321', name: 'cadeira' };

    objProducts.push(object1, object2);

    await createProduct.execute(objProducts);

    const dateAfterTenMinutes = addMinutes(Date.now(), 15);

    jest.spyOn(Date, 'now').mockImplementationOnce(() => {
      return new Date(dateAfterTenMinutes).getTime();
    });

    await expect(createProduct.execute(objProducts)).rejects.toBeInstanceOf(
      AppError,
    );
  });
});
