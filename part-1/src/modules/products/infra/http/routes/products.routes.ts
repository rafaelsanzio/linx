import 'reflect-metadata';

import { Router } from 'express';

import ProductsController from '../controllers/ProductsController';

const productsRouter = Router();

const Controller = new ProductsController();

productsRouter.post('/', Controller.create);

export default productsRouter;
