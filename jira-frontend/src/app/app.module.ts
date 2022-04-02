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
import { MatFormFieldModule } from '@angular/material/form-field';
import {HttpClient, HTTP_INTERCEPTORS} from '@angular/common/http'

import { MatInputModule } from '@angular/material/input';
import { ProjectComponent } from './project/project.component';
import { MatTableModule } from '@angular/material/table';
import { HttpClientModule } from '@angular/common/http';
import { ProjectListComponent } from './project/project-list/project-list.component';
import { ProjectSettingsComponent } from './project/project-settings/project-settings.component';
import { SprintComponent } from './sprint/sprint.component';
import { SprintListComponent } from './sprint/sprint-list/sprint-list.component';
// import { FlexLayoutModule } from "@angular/flex-layout";
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { MatListModule } from '@angular/material/list';
import { MatSelectModule } from '@angular/material/select';
import { NgbDropdown, NgbDropdownModule, NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { NewprojectComponent } from './project/newproject/newproject.component';
import { IssueListComponent } from './issue/issue-list/issue-list.component';
import { IssueCreateComponent } from './issue/issue-create/issue-create.component';
import { IssueModifyComponent } from './issue/issue-modify/issue-modify.component';
import { DragDropModule } from '@angular/cdk/drag-drop';
import { SprintModifyComponent } from './sprint/sprint-modify/sprint-modify.component';
import { NewsprintComponent } from './sprint/newsprint/newsprint.component';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import { IssueDetailComponent } from './issue/issue-detail/issue-detail.component';
import { AuthinterceptorService } from './interceptors/authinterceptor.service';
import { LoginComponentComponent } from './login-component/login-component.component';
import { RegisterComponent } from './register/register.component';
import {JwtHelperService, JwtModule, JWT_OPTIONS} from '@auth0/angular-jwt'




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
    IssueModifyComponent,
    SprintComponent,
    SprintListComponent,
    SprintModifyComponent,
    NewsprintComponent,
    IssueDetailComponent,
    LoginComponentComponent,
    RegisterComponent,
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
    DragDropModule,
    MatDatepickerModule,
    MatNativeDateModule,
    NgbModule,
    HttpClientModule,
  ],
  providers: [{
    provide: HTTP_INTERCEPTORS,
    useClass: AuthinterceptorService,
    multi: true
  },
  { provide: JWT_OPTIONS, useValue: JWT_OPTIONS },
        JwtHelperService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }