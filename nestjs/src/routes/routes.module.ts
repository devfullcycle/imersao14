import { Module } from '@nestjs/common';
import { RoutesService } from './routes.service';
import { RoutesController } from './routes.controller';
import { MapsModule } from '../maps/maps.module';
import { RoutesDriverService } from './routes-driver/routes-driver.service';
import { RoutesGateway } from './routes/routes.gateway';
import { BullModule } from '@nestjs/bull';
import { NewPointsJob } from './new-points.job';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { RouteKafkaProducerJob } from './route-kafka-producer.job';
import { makeCounterProvider } from '@willsoto/nestjs-prometheus';

@Module({
  imports: [
    MapsModule,
    BullModule.registerQueue(
      { name: 'new-points' },
      { name: 'kafka-producer' },
    ),
    // ClientsModule.register([
    //   {
    //     name: 'KAFKA_SERVICE',
    //     transport: Transport.KAFKA,
    //     options: {
    //       client: {
    //         clientId: 'nest',
    //         //kafka:9092
    //         brokers: [process.env.KAFKA_BROKER],
    //       },
    //     },
    //   },
    // ]),
    ClientsModule.registerAsync([
      {
        name: 'KAFKA_SERVICE',
        useFactory: () => ({
          transport: Transport.KAFKA,
          options: {
            client: {
              clientId: 'nest',
              brokers: [process.env.KAFKA_BROKER],
            },
          },
        }),
      },
    ]),
  ],
  controllers: [RoutesController],
  providers: [
    RoutesService,
    RoutesDriverService,
    RoutesGateway,
    NewPointsJob,
    RouteKafkaProducerJob,
    makeCounterProvider({
      name: 'route_started_counter',
      help: 'Number of routes started',
    }),
    makeCounterProvider({
      name: 'route_finished_counter',
      help: 'Number of routes finished',
    }),
  ],
})
export class RoutesModule {}
