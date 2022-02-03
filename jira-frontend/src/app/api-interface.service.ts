import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class ApiInterfaceService {
  url = 'http://10.20.81.30:5000'

  constructor(private http: HttpClient) { }
  
  getProjectList(data:any): Observable<any> {
    return this.http.get<any>(this.url+'/project')
  }

  updateProjectDetails(data:any): Observable<any> {
    return this.http.post<any>(this.url+'/update-project', data)
  }



  
}
