import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ApiInterfaceService } from 'src/app/api-interface.service';

@Component({
  selector: 'app-newproject',
  templateUrl: './newproject.component.html',
  styleUrls: ['./newproject.component.scss']
})
export class NewprojectComponent implements OnInit {
  newProjectForm = new FormGroup({
    projectName: new FormControl('', [Validators.required, Validators.minLength(3)]),
    projectKey: new FormControl('', [Validators.required]),
    projectDesc: new FormControl(''),
  });

  constructor(private apiService: ApiInterfaceService) { }

  ngOnInit(): void {
  }

  submit(){
    let body = {
      name:this.newProjectForm.get('projectName')?.value,
      description: this.newProjectForm.get('projectDesc')?.value,
      key: this.newProjectForm.get('projectKey')?.value,
      user_id: 1,
    } 

    this.apiService.createProject(body)
    .subscribe(res => {
          console.log(res);
    })
    
  }

}
