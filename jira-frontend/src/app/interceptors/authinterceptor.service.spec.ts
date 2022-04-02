import { HttpClientModule } from '@angular/common/http';
import { TestBed } from '@angular/core/testing';

import { AuthinterceptorService } from './authinterceptor.service';

describe('AuthinterceptorService', () => {
  let service: AuthinterceptorService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports:[HttpClientModule]
    });
    service = TestBed.inject(AuthinterceptorService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
