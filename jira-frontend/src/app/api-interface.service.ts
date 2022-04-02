import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import * as _ from 'lodash';
import { MatTableDataSource } from '@angular/material/table';

@Injectable({
  providedIn: 'root'
})
export class ApiInterfaceService {
  url = 'http://0.0.0.0:6000'
  post_url = "http://api.jira-clone.com/api"
  auth_url = "http://api.jira-clone.com"
  token  = ""

  apiResponse : any = []
  constructor(private http: HttpClient) { }
  
  getProjectList(data:any): Observable<any> {
    this.apiResponse = this.http.post<any>(this.post_url+'/project/list', data)
    return this.apiResponse
  }

  updateProjectDetails(data:any): Observable<any> {
    return this.http.post<any>(this.url+'/update-project', data)
  }

  createProject(data:any): Observable<any> {
    return this.http.post<any>(this.post_url+'/project/create', data)
  }

  createIssue(data:any): Observable<any> {
    return this.http.post<any>(this.post_url+'/issue/create', data)
  }

  createSprint(data:any): Observable<any> { 
    return this.http.post<any>(this.post_url+'/sprint/create', data)
  }

  updateIssue(data:any): Observable<any> {
    return this.http.post<any>(this.post_url+'/issues/update', data)
  }

  deleteIssue(data:any): Observable<any> {
    return this.http.post<any>(this.post_url+'/issue/delete', data)
  }

  deleteProject(data:any): Observable<any> {
    return this.http.post<any>(this.post_url+'/project/delete', data)
  }

  onChange($event:any){
    let filteredData = _.filter(this.apiResponse,(item) => {
      return item.key.LowerCaser() == $event.value.toLowerCase();
    })
    this.apiResponse = filteredData;
    console.log(this.apiResponse)
  }

  getSprintInfo(data:any): Observable<any>{
    return this.http.post<any>(this.post_url+'/sprint/info', data)
  }

  getIssueList(data:any): Observable<any>{
    return this.apiResponse = this.http.post<any>(this.post_url+'/issue/list',data)
  }

  getSprintList(data:any): Observable<any>{
     return this.apiResponse = this.http.post<any>(this.post_url+'/sprint/list', data)
  }

  getIssueDetails(data:any): Observable<any>{
    return this.apiResponse = this.http.post<any>(this.post_url+'/issue/info', data)
  }

  login(data:any): Observable<any>{
    return this.apiResponse = this.http.post<any>(this.auth_url+'/login', data)
  }

  register(data:any): Observable<any>{
    return this.apiResponse = this.http.post<any>(this.auth_url+'/register', data)
  }

  logout(data:any): Observable<any>{
    return this.apiResponse = this.http.post<any>(this.auth_url+'/logout', data)
  }

  getToken() {
    return localStorage.getItem("access-token");
  }

  setToken(token:string) {
    localStorage.setItem("access-token",token);
  }

}
