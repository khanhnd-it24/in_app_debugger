export interface INetwork {
  _id: string;
  deviceId: string;
  method: string;
  path: string;
  statusCode: number;
  request: string;
  response: string;
  createdAt: Date;
}
