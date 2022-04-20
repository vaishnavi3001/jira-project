import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators, FormArray } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiInterfaceService } from 'src/app/api-interface.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-issue-create',
  templateUrl: './issue-create.component.html',
  styleUrls: ['./issue-create.component.scss']
})
export class IssueCreateComponent implements OnInit {
  projects: any[] = [];
  sprints: any[] = [];
  members: any[] = [];
  issueDetailsForm!: FormGroup;
  constructor(private route: ActivatedRoute, private router: Router, private apiService: ApiInterfaceService,
    private _snackBar: MatSnackBar) {

  }


  ngOnInit(): void {
    this.issueDetailsForm = new FormGroup({
      issue_title: new FormControl(),
      project_id: new FormControl(),
      issue_type: new FormControl(),
      sprint_number: new FormControl(),
      upload_attachments: new FormControl(),
      issue_description: new FormControl(),
      issue_assignee: new FormControl(),
      issue_reporter: new FormControl(),
    })

    this.getProjects()
   
    
  }

  getProjects(){
    this.apiService.getProjectList({})
    .subscribe((result: any) => {
      this.projects = result['resp']['projects']
    })


  }

  getDetails(){
    this.getProjectSprints()
    this.getProjectMemebers()
  }

  getProjectSprints(){
    var body = {"project_id":Number(this.issueDetailsForm.get('project_id')?.value) }
    this.apiService.getSprintList(body)
    .subscribe((result:any) => {
      this.sprints = result['resp']['sprints']
    })
    console.log( Number(this.issueDetailsForm.get('project_id')?.value))
  }

  getProjectMemebers(){
    var body = {"project_id":Number(this.issueDetailsForm.get('project_id')?.value) }
    this.apiService.getProjectMemebers(body)
    .subscribe((result:any) => {
      this.members = result['resp']['members']
    })
  
  }

  onSubmit(): void {

    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'));

    let issueTypeMap = new Map<string, number>([["Epic", 1],["Task", 1], ["Sub Task", 2], ["Bug", 3]]);

    let _issue_type: string = this.issueDetailsForm.get('issue_type')?.value

    let body = {
      "user_id":1,
      "issue_title": this.issueDetailsForm.get('issue_title')?.value,
      "issue_description": this.issueDetailsForm.get('issue_description')?.value,
      "issue_type": issueTypeMap.get(_issue_type),
      "creator": Number(this.issueDetailsForm.get('issue_reporter')?.value),
      "assignee": Number(this.issueDetailsForm.get('issue_assignee')?.value),
      "sprint_id": Number(this.issueDetailsForm.get('sprint_number')?.value),
      "project_id": Number(this.issueDetailsForm.get('project_id')?.value)
    }

    this.apiService.createIssue(body)
      .subscribe(res => {
        
        if (res['status'] == true){
          this.createAlert("Issue created successfully!")
        }
        else {
          this.createAlert("Something went wrong.")
        }
      })

  }

  createAlert(message:string): void{
    this._snackBar.open(message, "Done");
  }

}
