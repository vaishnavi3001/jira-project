import { HttpClient, HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';

import { LoginComponentComponent } from './login-component.component';

describe('LoginComponentComponent', () => {
  let component: LoginComponentComponent;
  let fixture: ComponentFixture<LoginComponentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LoginComponentComponent ],
      imports:[HttpClientModule,RouterTestingModule]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LoginComponentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should have',()=>{
    const title = fixture.debugElement.nativeElement.querySelector('.modal-title');
    expect(title.innerHTML).not.toBeNull();
    const body= fixture.debugElement.nativeElement.querySelector('.modal-body');
    expect(body.innerHTML).not.toBeNull();
    const footer = fixture.debugElement.nativeElement.querySelector('.modal-footer');
    expect(body.innerHTML).not.toBeNull();
    
    
  })


});
