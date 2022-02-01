import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ListComponent } from './widgets/list/list.component';
import { PaginationComponent } from './widgets/pagination/pagination.component';



@NgModule({
  declarations: [
    ListComponent,
    PaginationComponent
  ],
  imports: [
    CommonModule
  ]
})
export class FrameworkModule { }
