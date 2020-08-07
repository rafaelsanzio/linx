type IProduct = Array<{
  id: string;
  name: string;
}>;

export default interface ICreateProductDTO {
  objProducts: IProduct;
  dataHashed: string;
}
