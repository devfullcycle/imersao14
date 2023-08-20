import { SubscribeMessage, WebSocketGateway } from '@nestjs/websockets';
import { Socket } from 'socket.io';
//import { RoutesDriverService } from '../routes-driver/routes-driver.service';
import { Queue } from 'bull';
import { InjectQueue } from '@nestjs/bull';

@WebSocketGateway({
  cors: {
    origin: '*',
  },
})
export class RoutesGateway {
  //constructor(private routesDriverService: RoutesDriverService) {}

  constructor(@InjectQueue('new-points') private newPointsQueue: Queue) {}

  @SubscribeMessage('new-points')
  async handleMessage(
    client: Socket,
    payload: { route_id: string; lat: number; lng: number },
  ) {
    client.broadcast.emit('admin-new-points', payload);
    client.broadcast.emit(`new-points/${payload.route_id}`, payload);
    await this.newPointsQueue.add(payload);
    //await this.routesDriverService.createOrUpdate(payload);
  }
}
