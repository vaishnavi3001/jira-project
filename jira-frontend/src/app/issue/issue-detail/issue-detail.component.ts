import { ApiInterfaceService } from 'src/app/api-interface.service';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-issue-detail',
  templateUrl: './issue-detail.component.html',
  styleUrls: ['./issue-detail.component.scss']
})
export class IssueDetailComponent implements OnInit {

  title = ""
  issue_creator = ""
  assignee_name = ""
  description = ""
  issue_created_at = ""

  constructor(private route: ActivatedRoute, private apiService:ApiInterfaceService) { }

  ngOnInit(): void {
    this.getIssueDetails();
  }

  sendcomment(comment:string): void {
    const routeParams = this.route.snapshot.paramMap;
    const issueIdFromRoute = Number(routeParams.get('issueId'));
    let body = {"issue_id":issueIdFromRoute,"comment":comment}
    console.log(body)
    this.apiService.createComment(body)
    .subscribe(res => {
          console.log(res);
    })
  }

  getIssueDetails(): void{
    const routeParams = this.route.snapshot.paramMap;
    const issueIdFromRoute = Number(routeParams.get('issueId'))
    this.apiService.getIssueDetails({'user_id':1,'issue_id':issueIdFromRoute})
    .subscribe((resp:any) => {
      console.log(resp)
      this.title = resp['resp']['title']
      this.issue_creator = resp['resp']['creator_name']
      this.assignee_name = resp['resp']['assignee_name']
      this.description = resp['resp']['description']
      this.issue_created_at = resp['resp']['created_at']

    })
  }

}
