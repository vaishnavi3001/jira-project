import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { JwtModule } from '@auth0/angular-jwt';

import { NewsprintComponent } from './newsprint.component';

describe('NewsprintComponent', () => {
  let component: NewsprintComponent;
  let fixture: ComponentFixture<NewsprintComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NewsprintComponent ],
      imports:[RouterTestingModule,HttpClientModule,JwtModule.forRoot({
        config: {
          tokenGetter: () => {
            return '';
          }
        }
      })]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NewsprintComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should have',()=>{
    const  body = fixture.debugElement.nativeElement.querySelector('.card-body');
    expect(body.innerHTML).not.toBeNull();
    const  f= fixture.debugElement.nativeElement.querySelector('.form-group');
    expect(f.innerHTML).not.toBeNull();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('h1').textContent).toContain('Create Sprint');
  });

  
});
