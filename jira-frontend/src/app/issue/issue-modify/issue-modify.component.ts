import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators, FormArray} from '@angular/forms';

@Component({
  selector: 'app-issue-modify',
  templateUrl: './issue-modify.component.html',
  styleUrls: ['./issue-modify.component.scss']
})
export class IssueModifyComponent implements OnInit {
  issueEditForm!: FormGroup;
  constructor() { }

  ngOnInit(): void {
    this.issueEditForm = new FormGroup({
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

  onSubmit(): void{
    console.log(this.issueEditForm.value);
  }

}
