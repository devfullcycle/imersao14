import { Test, TestingModule } from '@nestjs/testing';
import { RoutesDriverService } from './routes-driver.service';

describe('RoutesDriverService', () => {
  let service: RoutesDriverService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [RoutesDriverService],
    }).compile();

    service = module.get<RoutesDriverService>(RoutesDriverService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
