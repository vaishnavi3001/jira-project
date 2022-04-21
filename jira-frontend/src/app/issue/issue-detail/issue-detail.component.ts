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
  status = ""
  issue_id:any;
  comments = 0;
  memCount = 0;
  

  constructor(private route: ActivatedRoute, private apiService:ApiInterfaceService) { }

  ngOnInit(): void {
    this.getIssueDetails();

  }

  getIssueDetails(): void{
    const routeParams = this.route.snapshot.paramMap;
    const issueIdFromRoute = Number(routeParams.get('issueId'))
    this.issue_id = issueIdFromRoute
    
    this.apiService.getIssueDetails({'user_id':1,'issue_id':issueIdFromRoute})
    .subscribe((resp:any) => {
      console.log(resp)
      let issueTypeMap = new Map<number, string>([[0,"Epic"],[1, "Task"], [2,"Sub Task"], [3,"Bug"]]);
      let _issue_type: number = Number(resp['resp']['issue_status'])
      this.title = resp['resp']['title']
      this.issue_creator = resp['resp']['creator_name']
      this.assignee_name = resp['resp']['assignee_name']
      this.description = resp['resp']['description']
      this.issue_created_at = resp['resp']['created_at']
      this.status = String(issueTypeMap.get(_issue_type))
      this.apiService.getProjectStats({"project_id":resp['resp']['project_id']})
      .subscribe((resp)=> {
        this.comments = resp['resp']['comment_count']
        this.memCount = resp['resp']['member_count']
      })

    })
  }

  deleteIssue() {
    this.apiService.deleteIssue({"issue_id":this.issue_id})
    .subscribe((resp:any)=> {
      console.log("Issue Deleted successfully!")
    })

  }

}
