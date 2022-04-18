import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import * as _ from 'lodash';
import {JwtHelperService} from '@auth0/angular-jwt';
import { NgbTypeaheadWindow } from '@ng-bootstrap/ng-bootstrap/typeahead/typeahead-window';
import { environment } from 'src/environments/environment';

const jwtHelper = new JwtHelperService();
@Injectable({
  providedIn: 'root'
})
export class ApiInterfaceService {
  url = `${environment.url}`
  post_url = `${environment.post_url}`
  auth_url = `${environment.auth_url}`
  token  = ""

  
  apiResponse : any = []
  constructor(private http: HttpClient) { 

  }
  
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

  createComment(data:any): Observable<any> {
    return this.http.post<any>(this.post_url+'/comment/add', data)
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

  getProjectDetails(data:any): Observable<any>{
    return this.apiResponse = this.http.post<any>(this.post_url+'/project/info', data)
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

  isAuthenticated(){
    let currToken = this.getToken()||undefined;
    return !jwtHelper.isTokenExpired(currToken)
  }

}
