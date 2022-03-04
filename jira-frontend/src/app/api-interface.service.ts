import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import * as _ from 'lodash';
import { MatTableDataSource } from '@angular/material/table';

@Injectable({
  providedIn: 'root'
})
export class ApiInterfaceService {
  url = 'http://10.20.81.53:5000'
  apiResponse : any = []
  constructor(private http: HttpClient) { }
  
  getProjectList(data:any): Observable<any> {
    this.apiResponse = this.http.get<any>(this.url+'/project')
    return this.apiResponse
  }

  updateProjectDetails(data:any): Observable<any> {
    return this.http.post<any>(this.url+'/update-project', data)
  }

  createProject(data:any): Observable<any> {
    return this.http.post<any>(this.url+'/create-project', data)
  }

  deleteProject(data:any): Observable<any> {
    return this.http.post<any>(this.url+'/delete-project', data)
  }

  onChange($event:any){
    let filteredData = _.filter(this.apiResponse,(item) => {
      return item.key.LowerCaser() == $event.value.toLowerCase();
    })
    this.apiResponse = filteredData;
    console.log(this.apiResponse)
  }

  getIssueList(data:any): Observable<any>{
    return this.apiResponse = this.http.get<any>(this.url+'/issues/'+data.data)
  }
}
