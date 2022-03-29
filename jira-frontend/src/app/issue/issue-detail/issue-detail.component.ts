import { ApiInterfaceService } from 'src/app/api-interface.service';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-issue-detail',
  templateUrl: './issue-detail.component.html',
  styleUrls: ['./issue-detail.component.scss']
})
export class IssueDetailComponent implements OnInit {

  constructor(private route: ActivatedRoute, private apiService:ApiInterfaceService) { }

  ngOnInit(): void {
    this.getIssueDetails();
  }

  getIssueDetails(): void{
    const routeParams = this.route.snapshot.paramMap;
    const issueIdFromRoute = Number(routeParams.get('issueId'))
    this.apiService.getIssueDetails({'user_id':1,'issue_id':issueIdFromRoute})
    .subscribe((resp:any) => {
      console.log(resp)
    })
  }

}
