import { Test, TestingModule } from '@nestjs/testing';
import { DirectionsService } from './directions.service';

describe('DirectionsService', () => {
  let service: DirectionsService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [DirectionsService],
    }).compile();

    service = module.get<DirectionsService>(DirectionsService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
