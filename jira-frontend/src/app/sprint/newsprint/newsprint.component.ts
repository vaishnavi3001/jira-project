import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators, FormArray} from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiInterfaceService } from 'src/app/api-interface.service';
import { MAT_DATE_LOCALE } from '@angular/material/core';
import {formatDate} from '@angular/common';

export const MY_DATE_FORMATS = {
  parse: {
    dateInput: 'DD/MM/YYYY',
  },
  display: {
    dateInput: 'DD/MM/YYYY',
    monthYearLabel: 'MMMM YYYY',
    dateA11yLabel: 'LL',
    monthYearA11yLabel: 'MMMM YYYY'
  },
};

@Component({
  selector: 'app-newsprint',
  templateUrl: './newsprint.component.html',
  styleUrls: ['./newsprint.component.scss'],
  providers: [
    {provide: MAT_DATE_LOCALE, useValue: 'en-GB'},
  ]
})
export class NewsprintComponent implements OnInit {

  issueDetailsForm!: FormGroup;
  constructor(private route: ActivatedRoute, private router: Router , private apiService: ApiInterfaceService) { }

  ngOnInit(): void {
    this.issueDetailsForm = new FormGroup({
      
      sprint_name:new FormControl(),
      start_date:new FormControl(),
      end_date:new FormControl(),
      
    })
  }

  onSubmit(): void{
    

    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'));

    let body = {
      "user_id": 1,
      "sprint_name":this.issueDetailsForm.get('sprint_name')?.value,
      "start_date": this.issueDetailsForm.get('start_date')?.value,
      "end_date": this.issueDetailsForm.get('end_date')?.value,
      "project_id": projectIdFromRoute
  }
  let start_date_val = body.start_date
  start_date_val = formatDate(start_date_val,'yyyy-MM-dd','en');

  let end_date_val = body.end_date
  end_date_val = formatDate(end_date_val, 'yyyy-MM-dd','en');

  body.start_date = start_date_val
  body.end_date = end_date_val

  this.apiService.createSprint(body)
    .subscribe(res => {
      alert('Sprint Created Successfully!')
          console.log(res);
    })

  }

}
