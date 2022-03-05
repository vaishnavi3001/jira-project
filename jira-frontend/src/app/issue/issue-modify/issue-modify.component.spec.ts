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
});
