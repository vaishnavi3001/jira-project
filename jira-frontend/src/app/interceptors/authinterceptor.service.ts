import { HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { ThrowStmt } from '@angular/compiler';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthinterceptorService implements HttpInterceptor{
  token:string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2ODAzMTI1MTQsImlhdCI6MTY0ODc3NjUxNH0.kWl7lZpDsywQDUGKeEG4h7d_P9nF4HvoUnNUhzQXVsk"
  constructor() { }
  
  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    req = req.clone({
      setHeaders: {
        Authorization: 'Bearer '+ this.token
      },
      withCredentials:true
    });
    return next.handle(req);
  }
}
