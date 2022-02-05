import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators, FormArray} from '@angular/forms';
import { ApiInterfaceService } from 'src/app/api-interface.service';
import { ActivatedRoute } from '@angular/router';
import { ProjectSettings, project_data } from '../project-list/project-list.component';

@Component({
  selector: 'app-project-settings',
  templateUrl: './project-settings.component.html',
  styleUrls: ['./project-settings.component.scss']
})
export class ProjectSettingsComponent implements OnInit {
  // myForm = new FormGroup({
  //   projectName: new FormControl('', [Validators.required, Validators.minLength(3)]),
  //   projectDescription: new FormControl('', [Validators.required]),
  // });

  project_settings: ProjectSettings | undefined;
  projectDetailsForm!: FormGroup;

  constructor(private fb: FormBuilder, private apiService:ApiInterfaceService, private route: ActivatedRoute) { }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'));
    this.project_settings = project_data.find((project:any) => project.id === projectIdFromRoute)
    console.log(this.project_settings)

    // set the form values
    this.setFormValues();
  }

  // get f(){
  //   return this.myForm.controls;
  // }
  setFormValues(): void{
    this.projectDetailsForm = this.fb.group({
      id:this.project_settings?.id,
      name:this.project_settings?.name,
      key:this.project_settings?.key,
      lead:this.project_settings?.lead,
      type:this.project_settings?.type
    });
  }
  // submit(){
  //   let body = {
  //     name:this.myForm.get('projectName')?.value,
  //     description:this.myForm.get('projectDescription')?.value
  //   } 
  //   this.apiService.updateProjectDetails(body)
  //   .subscribe(res => {
  //         console.log(res);
          
  //   })
    
  // }
  onSubmit(): void{
    console.log(this.projectDetailsForm);
  }

}
