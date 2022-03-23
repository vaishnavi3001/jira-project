import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators, FormArray} from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiInterfaceService } from 'src/app/api-interface.service';

@Component({
  selector: 'app-newsprint',
  templateUrl: './newsprint.component.html',
  styleUrls: ['./newsprint.component.scss']
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
      "start_date": '2021-03-07',
      "end_date": '2021-03-21',
      "project_id": 2
  }

  this.apiService.createSprint(body)
    .subscribe(res => {
      alert('Sprint Created Successfully!')
          console.log(res);
    })

  }

}
