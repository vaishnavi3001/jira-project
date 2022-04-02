import { HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { ThrowStmt } from '@angular/compiler';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ApiInterfaceService } from '../api-interface.service';

@Injectable({
  providedIn: 'root'
})
export class AuthinterceptorService implements HttpInterceptor{
  token:string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2ODAzMTI1MTQsImlhdCI6MTY0ODc3NjUxNH0.kWl7lZpDsywQDUGKeEG4h7d_P9nF4HvoUnNUhzQXVsk"
  constructor(private apiService: ApiInterfaceService) { }
  
  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    if(!req.url.includes('login'))
    req = req.clone({
      setHeaders: {
        Authorization: 'Bearer '+ this.apiService.getToken()
      }
    });
    return next.handle(req);
  }
}
