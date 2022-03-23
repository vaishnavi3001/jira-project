import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators, FormArray} from '@angular/forms';
import { ApiInterfaceService } from 'src/app/api-interface.service';

@Component({
  selector: 'app-sprint-modify',
  templateUrl: './sprint-modify.component.html',
  styleUrls: ['./sprint-modify.component.scss']
})
export class SprintModifyComponent implements OnInit {
  sprintEditForm!: FormGroup;
  
  constructor(private apiService:ApiInterfaceService) { }

  ngOnInit(): void {
    this.sprintEditForm = new FormGroup({
      sprint_name:new FormControl(),
      start_date:new FormControl(),
      end_date:new FormControl(),
    })
    this.apiService
    .getSprintInfo({"sprint_id":2, "user_id":1 })
    .subscribe((resp:any) => {
      console.log(resp['resp']['sprint_name'])
      this.sprintEditForm.setValue({'sprint_name':resp['resp']['sprint_name'], 'start_date': resp['resp']['start_date'],'end_date': resp['resp']['end_date']})
      // this.sprintEditForm.setValue({'sprint_name':resp['resp']['sprint_name']})/
    })
    

    
  }

  onSubmit(): void{
    console.log(this.sprintEditForm.value);
    alert('Response saved successfully!')
  }

}
