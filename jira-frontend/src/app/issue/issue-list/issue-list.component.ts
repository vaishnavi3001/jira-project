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
  displayedColumns: string[] = ['project_id', 'status', 'type', 'title', 'created_by', 'sprint_id', 'assignee_id', 'assigned_to', 'modify'];
  dataSource = issueDataSource;

  constructor(private route: ActivatedRoute, private router: Router , private apiService: ApiInterfaceService) { }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'));

    this.apiService.getIssueList({'data':projectIdFromRoute})
    .subscribe((resp:any) =>{
      this.dataSource = resp['response']
    })
  }

  onSubmit(val=1): void{
    this.router.navigate(['home/issues/'+val+'/edit']);
  }

}
