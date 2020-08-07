import {
  ObjectID,
  Entity,
  Column,
  CreateDateColumn,
  ObjectIdColumn,
} from 'typeorm';

type IProduct = Array<{
  id: string;
  name: string;
}>;

@Entity('product')
class Product {
  @ObjectIdColumn()
  id: ObjectID;

  @Column()
  obj_products: IProduct;

  @Column()
  hash: String;

  @CreateDateColumn()
  created_at: Date;
}

export default Product;
