import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators, FormArray} from '@angular/forms';
import { ApiInterfaceService } from 'src/app/api-interface.service';
import { ActivatedRoute, Router } from '@angular/router';
import { ProjectSettings, project_data } from '../project-list/project-list.component';

@Component({
  selector: 'app-project-settings,',
  templateUrl: './project-settings.component.html',
  styleUrls: ['./project-settings.component.scss']
})
export class ProjectSettingsComponent implements OnInit {
  
  project_name = ""
  project_id = ""
  created_at = ""
  owner_username = ""

  constructor(private fb: FormBuilder, private apiService:ApiInterfaceService, private route: ActivatedRoute, private router: Router) { }

  ngOnInit(): void {
    this.get_project_data();
  }

  go_to_sprints(): void{
    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'));
    this.router.navigateByUrl('home/project/'+projectIdFromRoute+'/sprints')

  }

  get_project_data():void{
    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'))
    this.apiService.getProjectDetails({'project_id':projectIdFromRoute})
    .subscribe((resp:any) => {
      console.log(resp['resp']);
      this.project_name = resp['resp']['project_name'];
      this.project_id = resp['resp']['project_id'];
      this.created_at = resp['resp']['created_at'];
      this.owner_username = resp['resp']['owner_uname'];
    })
  }

  delete_project():void{
    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'));

    let body = {
      "project_id": projectIdFromRoute,
      "user_id": 1
    }
    console.log(projectIdFromRoute)

    this.apiService.deleteProject(body)
    .subscribe((resp:any) =>{
      console.log(resp['response'])
    })
  }

}
