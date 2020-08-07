import 'reflect-metadata';

import { Router } from 'express';

import ProductsController from '../controllers/ProductsController';

const productsRouter = Router();

const appointmentsController = new ProductsController();

productsRouter.post('/', appointmentsController.create);

export default productsRouter;
