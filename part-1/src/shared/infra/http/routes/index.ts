import { Router } from 'express';

import productsRouter from '../../../../modules/products/infra/http/routes/products.routes';

const routes = Router();

routes.use('/v1/products', productsRouter);

export default routes;
