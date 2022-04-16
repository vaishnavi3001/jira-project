import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IssueModifyComponent } from './issue-modify.component';

describe('IssueModifyComponent', () => {
  let component: IssueModifyComponent;
  let fixture: ComponentFixture<IssueModifyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IssueModifyComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(IssueModifyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  
  it('should have',()=>{
    const header= fixture.debugElement.nativeElement.querySelector('.card-header');
    expect(header.innerHTML).not.toBeNull();
    const random_text = fixture.debugElement.nativeElement.querySelector('.row');
    expect(random_text.innerHTML).not.toBeNull();
    const body = fixture.debugElement.nativeElement.querySelector('.card-body');
    expect(body.innerHTML).not.toBeNull();
    
    
  })
  
});
