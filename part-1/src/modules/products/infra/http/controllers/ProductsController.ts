import { Request, Response } from 'express';
import { container } from 'tsyringe';

import ImportProductService from '../../../services/ImportProductService';

export default class ProductsController {
  public async create(request: Request, response: Response): Promise<Response> {
    const { objProducts } = request.body;

    const createProducts = container.resolve(ImportProductService);

    const product = await createProducts.execute(objProducts);

    return response.json(product);
  }
}
