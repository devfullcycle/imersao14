import { Test, TestingModule } from '@nestjs/testing';
import { DirectionsController } from './directions.controller';

describe('DirectionsController', () => {
  let controller: DirectionsController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [DirectionsController],
    }).compile();

    controller = module.get<DirectionsController>(DirectionsController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
