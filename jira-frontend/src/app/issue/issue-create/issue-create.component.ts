import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators, FormArray} from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-issue-create',
  templateUrl: './issue-create.component.html',
  styleUrls: ['./issue-create.component.scss']
})
export class IssueCreateComponent implements OnInit {

  issueDetailsForm!: FormGroup;
  constructor() { }

  ngOnInit(): void {
    this.issueDetailsForm = new FormGroup({
      issue_title:new FormControl(),
      project_name:new FormControl(),
      issue_type:new FormControl(),
      sprint_number:new FormControl(),
      upload_attachments:new FormControl(),
      issue_description:new FormControl(),
      issue_assignee:new FormControl(),
      issue_reporter:new FormControl(),
    })
  }

}
