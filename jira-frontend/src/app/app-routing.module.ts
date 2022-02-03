import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ProjectListComponent } from './feature/project/project-list/project-list.component';

const routes: Routes = [
  {
    path: 'home',
    component: ProjectListComponent
  },
  // {
  //   path: '',
  //   loadChildren: './framework/auth/auth.module#AuthModule',
  //   pathMatch: 'prefix'
  // },
  { path: '', pathMatch: 'full', redirectTo: 'home' },
  { path: '**', redirectTo: '/home', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
