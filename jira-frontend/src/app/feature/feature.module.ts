import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProjectComponent } from './project/project.component';
import { ProjectListComponent } from './project/project-list/project-list.component';
import { UserComponent } from './user/user.component';
import { AuthComponent } from './auth/auth.component';



@NgModule({
  declarations: [
    ProjectComponent,
    ProjectListComponent,
    UserComponent,
    AuthComponent
  ],
  imports: [
    CommonModule
  ]
})
export class FeatureModule { }
