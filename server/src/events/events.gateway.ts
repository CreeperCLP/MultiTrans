import {
  WebSocketGateway,
  WebSocketServer,
  SubscribeMessage,
  MessageBody,
  ConnectedSocket
} from '@nestjs/websockets';
import { Server, Socket } from 'socket.io';

@WebSocketGateway({ cors: true })
export class EventsGateway {
  @WebSocketServer()
  server: Server;

  private connectedClients = new Map<string, string>(); // clientId -> deviceInfo

  handleConnection(client: Socket) {
    console.log(`Client connected: ${client.id}`);
  }

  handleDisconnect(client: Socket) {
    this.connectedClients.delete(client.id);
    console.log(`Client disconnected: ${client.id}`);
    this.broadcastDevices();
  }

  @SubscribeMessage('registerDevice')
  handleRegister(@MessageBody() data: any, @ConnectedSocket() client: Socket) {
    this.connectedClients.set(client.id, data.deviceInfo);
    this.broadcastDevices();
  }

  @SubscribeMessage('clipboardSync')
  handleClipboard(@MessageBody() data: any, @ConnectedSocket() client: Socket) {
    // 广播接收到的剪贴板数据给除了发送者之外的所有设备
    client.broadcast.emit('onClipboardReceive', data);
  }

  private broadcastDevices() {
    this.server.emit('devicesUpdated', Array.from(this.connectedClients.values()));
  }
}