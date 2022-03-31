import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators, FormArray } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiInterfaceService } from 'src/app/api-interface.service';

@Component({
  selector: 'app-issue-create',
  templateUrl: './issue-create.component.html',
  styleUrls: ['./issue-create.component.scss']
})
export class IssueCreateComponent implements OnInit {

  issueDetailsForm!: FormGroup;
  constructor(private route: ActivatedRoute, private router: Router, private apiService: ApiInterfaceService) {

  }

  ngOnInit(): void {
    this.issueDetailsForm = new FormGroup({
      issue_title: new FormControl(),
      project_name: new FormControl(),
      issue_type: new FormControl(),
      sprint_number: new FormControl(),
      upload_attachments: new FormControl(),
      issue_description: new FormControl(),
      issue_assignee: new FormControl(),
      issue_reporter: new FormControl(),
    })
  }

  onSubmit(): void {

    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'));

    let issueTypeMap = new Map<string, number>([["Task", 1], ["Sub Task", 2], ["Bug", 3]]);

    let _issue_type: string = this.issueDetailsForm.get('issue_type')?.value

    let body = {
      "user_id": 1,
      "issue_title": this.issueDetailsForm.get('issue_title')?.value,
      "issue_description": this.issueDetailsForm.get('issue_description')?.value,
      "issue_type": issueTypeMap.get(_issue_type),
      "creator": 1,
      "assignee": 1,
      "sprint_id": 1,
      "project_id": projectIdFromRoute
    }

    this.apiService.createIssue(body)
      .subscribe(res => {
        console.log(body)
        console.log(res);
      })

  }

}
