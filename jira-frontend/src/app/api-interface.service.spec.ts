import { HttpClientModule } from '@angular/common/http';
import { TestBed } from '@angular/core/testing';
import { JwtHelperService, JwtModule } from '@auth0/angular-jwt';

import { ApiInterfaceService } from './api-interface.service';

describe('ApiInterfaceService', () => {
  let service: ApiInterfaceService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports:[HttpClientModule,JwtModule.forRoot({
        config: {
          tokenGetter: () => {
            return '';
          }
        }
      })]
    });
    service = TestBed.inject(ApiInterfaceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
