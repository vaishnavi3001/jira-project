import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';

import { IssueCreateComponent } from './issue-create.component';

describe('IssueCreateComponent', () => {
  let component: IssueCreateComponent;
  let fixture: ComponentFixture<IssueCreateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IssueCreateComponent],
      imports:[RouterTestingModule,HttpClientModule]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(IssueCreateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
    
  });

  it('should have',()=>{
    const title = fixture.debugElement.nativeElement.querySelector('.fa-plus');
    expect(title.innerHTML).not.toBeNull();
    const random_text = fixture.debugElement.nativeElement.querySelector('.text-muted');
    expect(random_text.innerHTML).toBe('Max Size 3mb');
  })
});
