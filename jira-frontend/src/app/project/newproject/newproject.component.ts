import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ApiInterfaceService } from 'src/app/api-interface.service';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';

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

  constructor(private apiService: ApiInterfaceService, private _snackBar: MatSnackBar, private router: Router) { }

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
          if (res['status'] == true){
            this.createAlert("Project created successfully!")
            this.router.navigateByUrl('/home/projects');
          }
          else {
            this.createAlert("Something went wrong.")
          }
    })
    
  }

  createAlert(message:string): void{
    this._snackBar.open(message, "Done");
  }

}
