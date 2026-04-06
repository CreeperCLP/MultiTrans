import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { EventsGateway } from './events/events.gateway';
import { FileController } from './file/file.controller';

@Module({
  imports: [],
  controllers: [AppController, FileController],
  providers: [EventsGateway],
})
export class AppModule {}