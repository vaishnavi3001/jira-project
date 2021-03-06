import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';

import { IssueDetailComponent } from './issue-detail.component';

describe('IssueDetailComponent', () => {
  let component: IssueDetailComponent;
  let fixture: ComponentFixture<IssueDetailComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IssueDetailComponent ],
      imports:[RouterTestingModule,HttpClientModule]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(IssueDetailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  
  
  it('should have',()=>{
    const title = fixture.debugElement.nativeElement.querySelector('.m-b-md');
    expect(title.innerHTML).not.toBeNull();
    const b = fixture.debugElement.nativeElement.querySelector('.panel-body');
    expect(b.innerHTML).not.toBeNull();
    const c = fixture.debugElement.nativeElement.querySelector('.container');
    expect(c.innerHTML).not.toBeNull();
    
  })

  
});
