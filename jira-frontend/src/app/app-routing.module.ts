import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { NewprojectComponent } from './project/newproject/newproject.component';
import { ProjectListComponent } from './project/project-list/project-list.component';
import { ProjectSettingsComponent } from './project/project-settings/project-settings.component';
import { ProjectComponent } from './project/project.component';
import { IssueListComponent } from './issue/issue-list/issue-list.component';
import { IssueCreateComponent } from './issue/issue-create/issue-create.component';
import { SprintListComponent } from './sprint/sprint-list/sprint-list.component';
import { IssueModifyComponent } from './issue/issue-modify/issue-modify.component';

const routes: Routes = [
  { path: 'home', component: ProjectComponent,
    children: [
      { path: '', pathMatch: 'full', redirectTo: 'list' },
      {
        path:'list', component: ProjectListComponent
      },
      {
        path: 'newproject', component: NewprojectComponent
      },
      {
       path:'settings/:projectId', component: ProjectSettingsComponent
      },
      {
        path:'issues/create', component: IssueCreateComponent
      },
      {
        path:'issues/:projectId', component: IssueListComponent
      },
      {
        path:'sprint/:projectId', component: SprintListComponent
      },
      {
        path:'issues/:issueId/edit', component: IssueModifyComponent
      },      
    ]},

  { path: '', pathMatch: 'full', redirectTo: 'home' },
  //{ path: '**', redirectTo: '/home', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
