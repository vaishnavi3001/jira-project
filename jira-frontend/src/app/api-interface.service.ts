import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import * as _ from 'lodash';
import { JwtHelperService } from '@auth0/angular-jwt';
import { environment } from 'src/environments/environment';


const jwtHelper = new JwtHelperService();
@Injectable({
  providedIn: 'root'
})
export class ApiInterfaceService {
  backend_url = `${environment.backend_url}`
  api_url = `${environment.backend_url}/api`
  token = ""

  apiResponse: any = []
  constructor(private http: HttpClient) {

  }

  getProjectList(data: any): Observable<any> {
    this.apiResponse = this.http.post<any>(this.api_url + '/project/list', data)
    return this.apiResponse
  }

  updateProjectDetails(data: any): Observable<any> {
    return this.http.post<any>(this.api_url + '/update-project', data)
  }

  createProject(data: any): Observable<any> {
    return this.http.post<any>(this.api_url + '/project/create', data)
  }

  createIssue(data: any): Observable<any> {
    return this.http.post<any>(this.api_url + '/issue/create', data)
  }

  createSprint(data: any): Observable<any> {
    return this.http.post<any>(this.api_url + '/sprint/create', data)
  }
  updateIssueStatus(data:any): Observable<any> {
    return this.http.post<any>(this.api_url+'/issue/move', data)
  }

  createComment(data: any): Observable<any> {
    return this.http.post<any>(this.api_url + '/comment/add', data)
  }

  updateIssue(data: any): Observable<any> {
    return this.http.post<any>(this.api_url + '/issues/update', data)
  }

  getComments(data:any): Observable<any>{
    return this.http.post<any>(this.api_url+'/comment/view', data)
  }
  
  deleteIssue(data: any): Observable<any> {
    return this.http.post<any>(this.api_url + '/issue/delete', data)
  }


  deleteProject(data: any): Observable<any> {
    return this.http.post<any>(this.api_url + '/project/delete', data)
  }

  onChange($event: any) {
    let filteredData = _.filter(this.apiResponse, (item) => {
      return item.key.LowerCaser() == $event.value.toLowerCase();
    })
    this.apiResponse = filteredData;
    console.log(this.apiResponse)
  }

  getSprintInfo(data: any): Observable<any> {
    return this.http.post<any>(this.api_url + '/sprint/info', data)
  }

  getIssueList(data: any): Observable<any> {
    return this.apiResponse = this.http.post<any>(this.api_url + '/issue/list', data)
  }

  getSprintList(data: any): Observable<any> {
    return this.apiResponse = this.http.post<any>(this.api_url + '/sprint/list', data)
  }

  getIssueDetails(data: any): Observable<any> {
    return this.apiResponse = this.http.post<any>(this.api_url + '/issue/info', data)
  }

  getProjectDetails(data: any): Observable<any> {
    return this.apiResponse = this.http.post<any>(this.api_url + '/project/info', data)
  }


  getProjectStats(data:any): Observable<any>{
    return this.apiResponse = this.http.post<any>(this.api_url+'/project/stats', data)
  }

  login(data:any): Observable<any>{
    return this.apiResponse = this.http.post<any>(this.backend_url+'/login', data)
  }

  register(data: any): Observable<any> {
    return this.apiResponse = this.http.post<any>(this.backend_url + '/register', data)
  }


  logout(): Observable<any> {
    return this.apiResponse = this.http.post<any>(this.backend_url + '/logout', {})
  }

  getToken() {
    return localStorage.getItem("access-token");
  }

  setToken(token: string) {
    localStorage.setItem("access-token", token);
  }

  removeToken(){
    localStorage.removeItem("access-token")
  }

  isAuthenticated() {
    let currToken = this.getToken() || undefined;
    return !jwtHelper.isTokenExpired(currToken)
  }

  joinProject(invitelink: string|null) {
    return this.http.post(this.api_url + '/user/verify', { 'invite_link': invitelink })
  }

  sendInvite(data:any){
    return this.http.post(this.api_url + '/user/invite',data)
  }
  getUserProfile(): Observable<any>{
    return this.apiResponse = this.http.get<any>(this.api_url+'/user/info')
  }

  editUserProfile(data: any): Observable<any>{
    return this.apiResponse = this.http.patch<any>(this.api_url+'/user/info', data)
  }

  changePassword(data: any): Observable<any>{
    return this.apiResponse = this.http.put<any>(this.api_url+'/user/change-password', data)
  }

  getProjectMemebers(data:any): Observable<any>{
    return this.apiResponse = this.http.post<any>(this.post_url+'/project/members', data)
  }
}
