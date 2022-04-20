import { HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ApiInterfaceService } from '../api-interface.service';

@Injectable({
  providedIn: 'root'
})
export class AuthinterceptorService implements HttpInterceptor{

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
