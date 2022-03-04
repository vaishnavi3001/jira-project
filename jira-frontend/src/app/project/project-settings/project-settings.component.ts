import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators, FormArray} from '@angular/forms';
import { ApiInterfaceService } from 'src/app/api-interface.service';
import { ActivatedRoute } from '@angular/router';
import { ProjectSettings, project_data } from '../project-list/project-list.component';

@Component({
  selector: 'app-project-settings,',
  templateUrl: './project-settings.component.html',
  styleUrls: ['./project-settings.component.scss']
})
export class ProjectSettingsComponent implements OnInit {
  
  project_settings: ProjectSettings | undefined;
  projectDetailsForm!: FormGroup;

  constructor(private fb: FormBuilder, private apiService:ApiInterfaceService, private route: ActivatedRoute) { }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'));
    this.project_settings = project_data.find((project:any) => project.id === projectIdFromRoute)

    // set the form values
    this.setFormValues();
  }

  setFormValues(): void{
    this.projectDetailsForm = this.fb.group({
      id:this.project_settings?.id,
      name:this.project_settings?.name,
      key:this.project_settings?.key,
      lead:this.project_settings?.lead,
      type:this.project_settings?.type
    });
  }

  onSubmit(): void{
    console.log(this.projectDetailsForm);
  }

}
