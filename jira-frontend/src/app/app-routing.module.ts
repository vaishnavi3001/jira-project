import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ProjectListComponent } from './project/project-list/project-list.component';
import { ProjectSettingsComponent } from './project/project-settings/project-settings.component';
import { ProjectComponent } from './project/project.component';

const routes: Routes = [
  { path: 'home', component: ProjectComponent,
    children: [
      { path: '', pathMatch: 'full', redirectTo: 'list' },
      {
        path:'list', component: ProjectListComponent
      },
      {
        path:'settings', component: ProjectSettingsComponent
      }
    ]},

  { path: '', pathMatch: 'full', redirectTo: 'home' },
  { path: '**', redirectTo: '/home', pathMatch: 'full' }];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
