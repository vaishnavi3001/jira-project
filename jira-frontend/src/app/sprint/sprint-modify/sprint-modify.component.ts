import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators, FormArray} from '@angular/forms';

@Component({
  selector: 'app-sprint-modify',
  templateUrl: './sprint-modify.component.html',
  styleUrls: ['./sprint-modify.component.scss']
})
export class SprintModifyComponent implements OnInit {
  sprintEditForm!: FormGroup;
  constructor() { }

  ngOnInit(): void {
    this.sprintEditForm = new FormGroup({
      sprint_name:new FormControl(),
      start_date:new FormControl(),
      end_date:new FormControl(),
    })
  }

  onSubmit(): void{
    console.log(this.sprintEditForm.value);
  }

}
