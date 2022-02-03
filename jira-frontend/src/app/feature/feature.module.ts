import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProjectComponent } from './project/project.component';
import { ProjectListComponent } from './project/project-list/project-list.component';
import { UserComponent } from './user/user.component';
import { AuthComponent } from '../framework/auth/auth.component';
import { FrameworkModule } from '../framework/framework.module';
import { ListComponent } from '../framework/widgets/list/list.component';



@NgModule({
  declarations: [
    ProjectComponent,
    ProjectListComponent,
    UserComponent,
    AuthComponent,
    ListComponent
  ],
  imports: [
    CommonModule,
    FrameworkModule
  ]
})
export class FeatureModule { }
