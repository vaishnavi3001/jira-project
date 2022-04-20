import { SprintDetailComponent } from './sprint/sprint-detail/sprint-detail.component';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { NewprojectComponent } from './project/newproject/newproject.component';
import { ProjectListComponent } from './project/project-list/project-list.component';
import { ProjectSettingsComponent } from './project/project-settings/project-settings.component';
import { ProjectComponent } from './project/project.component';
import { IssueListComponent, IssueDetails } from './issue/issue-list/issue-list.component';
import { IssueCreateComponent } from './issue/issue-create/issue-create.component';
import { SprintListComponent } from './sprint/sprint-list/sprint-list.component';
import { IssueModifyComponent } from './issue/issue-modify/issue-modify.component';
import { IssueDetailComponent } from './issue/issue-detail/issue-detail.component';
import { NewsprintComponent } from './sprint/newsprint/newsprint.component';
import { SprintModifyComponent } from './sprint/sprint-modify/sprint-modify.component';
import { RegisterComponent } from './register/register.component';
import { LoginComponentComponent } from './login-component/login-component.component';
import { AuthguardService } from './guards/authguard.service';
import { UserProfileComponent } from './user/user-profile/user-profile.component';


const routes: Routes = [
  { path: 'register', component: RegisterComponent },
  { path: 'login', component: LoginComponentComponent },

  { path: 'home', component: ProjectComponent, canActivate: [AuthguardService], 

    children: [
      { path: '', pathMatch: 'full', redirectTo: 'list' },
      {
        path:'projects', component: ProjectListComponent
      },
      {
        path: 'project/newproject', component: NewprojectComponent
      },
      {
       path:'project/:projectId/settings', component: ProjectSettingsComponent
      },
      {
        path:':projectId/issues/create', component: IssueCreateComponent
      },
      {
        path:'project/:projectId/issues', component: IssueListComponent
      },
      {
        path:'issue/:issueId/edit', component: IssueModifyComponent
      },
      {
        path:'issue/:issueId/details', component: IssueDetailComponent
      },  
      {
        path:':projectId/sprint/create', component: NewsprintComponent
      },  
      {
        path:'sprint/:sprintId/edit', component: SprintModifyComponent
      },    
      {
        path:'project/:projectId/sprints', component: SprintListComponent
      },
      {
        path:'sprint/:sprintId/details', component: SprintDetailComponent
      },
      {
        path:'sampleprofile', component: UserProfileComponent
      }
    ]},

  { path: '**', pathMatch: 'full', redirectTo: 'home/projects' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
