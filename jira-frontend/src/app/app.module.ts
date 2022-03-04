import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
//import { NavbarComponent } from './navbar/navbar.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatButtonModule } from '@angular/material/button';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { NavbarComponent } from './navbar/navbar.component';
import {MatFormFieldModule} from '@angular/material/form-field';

import { MatInputModule } from '@angular/material/input';
import { ProjectComponent } from './project/project.component';
import {MatTableModule} from '@angular/material/table';
import { HttpClientModule } from '@angular/common/http';
import { ProjectListComponent } from './project/project-list/project-list.component';
import { ProjectSettingsComponent } from './project/project-settings/project-settings.component';
// import { FlexLayoutModule } from "@angular/flex-layout";
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import {MatListModule} from '@angular/material/list';
import {MatSelectModule} from '@angular/material/select';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { NewprojectComponent } from './project/newproject/newproject.component';
import { IssueListComponent } from './issue/issue-list/issue-list.component';
import { IssueCreateComponent } from './issue/issue-create/issue-create.component';




//import { NavbarComponent } from './navbar/navbar.component';






@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    ProjectComponent,
    ProjectListComponent,
    ProjectSettingsComponent,
    NewprojectComponent,
    IssueListComponent,
    IssueCreateComponent,
    //NavbarComponent,
    //NavbarComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    //NavbarComponent,
    MatButtonModule,
    MatToolbarModule,
    MatIconModule,
    MatFormFieldModule,
    MatTableModule,
    HttpClientModule,
    ReactiveFormsModule,
    FormsModule,
    MatIconModule,
    MatInputModule,
    MatListModule,
    MatSelectModule,
    NgbModule
    
    
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
