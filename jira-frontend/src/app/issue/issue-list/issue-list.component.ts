import { FormGroup } from '@angular/forms';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiInterfaceService } from 'src/app/api-interface.service';

export interface IssueDetails {
  project_id: number,
  status: string,
  type: string,
  title: string,
  created_by: string,
  sprint_id: string,
  assignee_id: string,
  assigned_to: string,
  modify: string,
}

export var issueDataSource: IssueDetails[] = [];

@Component({
  selector: 'app-issue-list',
  templateUrl: './issue-list.component.html',
  styleUrls: ['./issue-list.component.scss']
})
export class IssueListComponent implements OnInit {
  displayedColumns: string[] = ['issue_id', 'project_id', 'status', 'type', 'title', 'created_by', 'sprint_id', 'assignee_id', 'assigned_to', 'modify'];
  dataSource = issueDataSource;
  //projectIdFromRoute:any 

  constructor(private route: ActivatedRoute, private router: Router , private apiService: ApiInterfaceService) { }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'));

    this.apiService.getIssueList({'user_id':1, 'sprint_id': 1})
    .subscribe((resp:any) =>{
      this.dataSource = resp['resp']['issues']
    })
  }

  onSubmit(val=1): void{
    this.router.navigate(['home/issues/'+val+'/edit']);
  }

  onCreateIssue(): void{
    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'));
    this.router.navigate(['home/'+projectIdFromRoute+'/issues/'+'/create']);
  }

  getDetailsPage(element: any): void{
    console.log(element.issue_id)
    this.router.navigate(['home/issue/'+element.issue_id+'/details'])
  }

  onDeleteProject(): void{
    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'));

    let body = {
      "project_id": projectIdFromRoute,
      "user_id": 1
    }

    this.apiService.deleteProject({'data':body})
    .subscribe((resp:any) =>{
      this.dataSource = resp['response']
      console.log(this.dataSource)
    })

  }

}
