import { DebugElement } from '@angular/core';
import {
  ComponentFixture,
  //fakeAsync,
  TestBed,
  //tick,
 // waitForAsync,
} from '@angular/core/testing';
import { FormsModule } from '@angular/forms';
import { By } from '@angular/platform-browser';
import { Router } from '@angular/router';
import { throwError } from 'rxjs';
//import { Login-Component } from './login-component.component';
import { LoginComponentComponent } from './login-component.component';
//import { LoginService } from './login.service';
import { ApiInterfaceService } from 'src/app/api-interface.service';

class Page {
  get submitButton() {
    return this.fixture.nativeElement.querySelector('button');
  }
  get usernameInput() {
    return this.fixture.debugElement.nativeElement.querySelector('#username');
  }
  get passwordInput() {
    return this.fixture.debugElement.nativeElement.querySelector('#pwd');
  }

  get errorMsg() {
    return this.fixture.debugElement.query(By.css('.error')).nativeElement;
  }

  constructor(private fixture: ComponentFixture<LoginComponentComponent>) {}

  public updateValue(input: HTMLInputElement, value: string) {
    input.value = value;
    input.dispatchEvent(new Event('input'));
  }
}
describe('Login Component', () => {
  let loginComponent: LoginComponentComponent;
  let fixture: ComponentFixture<LoginComponentComponent>;
  let debugEl: DebugElement;

  let loginService: ApiInterfaceService ;
  let loginServiceSpy: { login: jasmine.Spy };
  let routerSpy: { navigateByUrl: jasmine.Spy };
  let router: Router;
  let page: Page;
  beforeEach(() => {
    loginServiceSpy = jasmine.createSpyObj(ApiInterfaceService , ['login']);
    routerSpy = jasmine.createSpyObj(Router, ['navigateByUrl']);
    TestBed.configureTestingModule({
      imports: [FormsModule],
      declarations: [LoginComponentComponent],
      providers: [
        { provide: ApiInterfaceService , useValue: loginServiceSpy },
        { provide: Router, useValue: routerSpy },
      ],
    });
    fixture = TestBed.createComponent(LoginComponentComponent);
    loginComponent = fixture.componentInstance;
    debugEl = fixture.debugElement;
    loginService = TestBed.inject(ApiInterfaceService );
    router = TestBed.inject(Router);
    page = new Page(fixture);
    fixture.detectChanges();

    

  });

  
});