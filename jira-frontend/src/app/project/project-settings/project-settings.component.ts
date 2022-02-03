import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators} from '@angular/forms';
import { ApiInterfaceService } from 'src/app/api-interface.service';

@Component({
  selector: 'app-project-settings',
  templateUrl: './project-settings.component.html',
  styleUrls: ['./project-settings.component.scss']
})
export class ProjectSettingsComponent implements OnInit {
  myForm = new FormGroup({
    projectName: new FormControl('', [Validators.required, Validators.minLength(3)]),
    projectDescription: new FormControl('', [Validators.required]),
  });

  constructor(private apiService:ApiInterfaceService) { }

  ngOnInit(): void {
  }

  get f(){
    return this.myForm.controls;
  }

  submit(){
    let body = {
      name:this.myForm.get('projectName')?.value,
      description:this.myForm.get('projectDescription')?.value
    } 
    this.apiService.updateProjectDetails(body)
    .subscribe(res => {
          console.log(res);
          
    })
    
  }

}
